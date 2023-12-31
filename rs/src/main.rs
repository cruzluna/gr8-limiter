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
#[allow(dead_code)]
pub struct AppState {
    redis_conn: redis::Connection,
}

#[tokio::main]
async fn main() {
    println!("Server started...");
    // TODO: Logger

    // redis
    let client = redis::Client::open("redis://127.0.0.1/")
        .expect("Invalid connection URL")
        .get_connection()
        .expect("Couldn't connect to redis");

    let state = AppState { redis_conn: client };
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
    for (key, value) in headers.iter() {
        println!("Header: {:?} - {:?}", key, value);
    }

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

    match rate_limit(state.redis_conn, &rc) {
        Ok(_) => (StatusCode::TOO_MANY_REQUESTS, "Rate limited"),
        Err(_) => (StatusCode::OK, "Good."),
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
fn rate_limit(client: redis::Connection, rc: &RateLimitConfig) -> Result<bool, redis::RedisError> {
    let now = chrono::Utc::now().timestamp();
    client.zrembyscore(&rc.id, 0, now)?;

    match client.zcard::<&String, i64>(&rc.id) {
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
    let _ = client.zadd::<&String, i64, f64, i64>(&rc.id, now as f64, now);

    // whole set should expire after window size
    let _ = client.expire::<&String, i64>(
        &rc.id,
        chrono::Duration::seconds(rc.window_size).num_seconds(),
    );

    // not rate limited
    Ok(false)
}
