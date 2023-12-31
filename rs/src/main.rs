use axum::{
    extract::State,
    http::{HeaderMap, StatusCode},
    response::IntoResponse,
    routing::post,
    Router,
};
use axum_macros::debug_handler;
use chrono;
use redis::Commands;

#[derive(Clone)]
pub struct AppState {
    redis_client: redis::Client,
}

impl AppState {
    pub fn new(redis_client: redis::Client) -> Self {
        AppState { redis_client }
    }

    pub fn get_redis_connection(&self) -> redis::Connection {
        self.redis_client
            .get_connection()
            .expect("Failed to get Redis connection")
    }
}

#[tokio::main]
async fn main() {
    println!("Server started...");
    // TODO: Logger

    // redis
    // let mut client = redis::Client::open("redis://127.0.0.1/")
    //     .expect("Invalid connection URL")
    //     .get_connection()
    //     .expect("Couldn't connect to redis");

    let client = redis::Client::open("redis://127.0.0.1/").expect("Invalid connection URL");

    let state = AppState {
        redis_client: client,
    };
    let app = Router::new()
        .route("/api/v1/ratelimit", post(post_rate_limit))
        .fallback(handler_404)
        .with_state(state.clone());

    // run our app with hyper, listening on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

#[debug_handler]
// TODO: Fix the return type
async fn post_rate_limit(headers: HeaderMap, State(state): State<AppState>) -> impl IntoResponse {
    // for (key, value) in headers.iter() {
    //     println!("Header: {:?} - {:?}", key, value);
    // }

    let api_key = match headers.get("x-api-key") {
        // TODO: Change the unwrap
        Some(val) => val.to_str().map(|x| x.to_string()).unwrap(),
        None => return (StatusCode::BAD_REQUEST, "Invalid api key."),
    };

    // Rate Limit config
    let rc = RateLimitConfig {
        id: api_key,
        limit: 5,
        window_size: 1,
    };

    let conn = state.get_redis_connection();

    match rate_limit(conn, &rc) {
        Ok(should_rate_limit) => {
            if should_rate_limit {
                (StatusCode::TOO_MANY_REQUESTS, "Rate limited")
            } else {
                (StatusCode::OK, "Good.")
            }
        }
        Err(_) => (StatusCode::INTERNAL_SERVER_ERROR, "Something went wrong."),
    }
}

async fn handler_404() -> impl IntoResponse {
    (StatusCode::NOT_FOUND, "Route not found.")
}

// #[derive(Default)]
struct RateLimitConfig {
    id: String,
    limit: i64,       // default 5
    window_size: i64, // default 1
}

/// .
///
/// # Errors
///
/// This function will return an error if .
// sliding window algo. Rate limited returns true and not rate limited is false.
fn rate_limit(
    mut redis_con: redis::Connection,
    rc: &RateLimitConfig,
) -> Result<bool, redis::RedisError> {
    println!("Made it into rate_limt");
    let now = chrono::Utc::now().timestamp();
    redis_con.zrembyscore(&rc.id, 0, now)?;

    match redis_con.zcard::<&String, i64>(&rc.id) {
        Ok(cur) => {
            // check if reqs is within limit
            if cur > rc.limit {
                // ratelimited
                println!("rate limited: {}", cur);
                return Ok(true);
            }
        }
        Err(e) => return Err(e),
    };

    // req not ratelimited, add it to the curr window
    let _ = redis_con.zadd::<&String, i64, f64, i64>(&rc.id, now as f64, now);

    // whole set should expire after window size
    let _ = redis_con.expire::<&String, i64>(
        &rc.id,
        chrono::Duration::seconds(rc.window_size).num_seconds(),
    );

    // Runtime error
    let dum = redis_con.get::<&String, i64>(&rc.id).unwrap();
    println!("Key: {} - Value: {}", &rc.id, dum);

    // not rate limited
    Ok(false)
}
