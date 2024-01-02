use axum::{
    extract::State,
    http::{HeaderMap, StatusCode},
    response::IntoResponse,
    routing::post,
    Router,
};

use axum_macros::debug_handler;
use bb8_redis::{bb8, RedisConnectionManager};

#[derive(Clone)]
pub struct AppState {
    pool: bb8::Pool<RedisConnectionManager>,
}

#[tokio::main]
async fn main() {
    println!("Server started...");
    // TODO: Logger

    // redis
    let manager = RedisConnectionManager::new("redis://127.0.0.1/").unwrap();
    let pool = bb8::Pool::builder().build(manager).await.unwrap();

    // let mut client = redis::Client::open("redis://127.0.0.1/")
    //     .expect("Invalid connection URL")
    //     .get_connection()
    //     .expect("Couldn't connect to redis");

    // let client = redis::Client::open("redis://127.0.0.1/").expect("Invalid connection URL");

    let state = AppState { pool };

    let app = Router::new()
        .route("/api/v1/ratelimit", post(post_rate_limit))
        .fallback(handler_404)
        .with_state(state);
    // .with_state(state.clone());

    // run our app with hyper, listening on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
    // Ok(())
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
    let _rc = RateLimitConfig {
        id: api_key,
        limit: 5,
        window_size: 1,
    };

    (StatusCode::TOO_MANY_REQUESTS, "Rate limited")

    // match rate_limit(conn, &rc) {
    //     Ok(should_rate_limit) => {
    //         if should_rate_limit {
    //             (StatusCode::TOO_MANY_REQUESTS, "Rate limited")
    //         } else {
    //             (StatusCode::OK, "Good.")
    //         }
    //     }
    //     Err(_) => (StatusCode::INTERNAL_SERVER_ERROR, "Something went wrong."),
    // }
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
// sliding log algo. Rate limited returns true and not rate limited is false.
async fn rate_limit(
    redis_con: redis::Connection,
    rc: &RateLimitConfig,
) -> Result<bool, redis::RedisError> {
    println!("Made it into rate_limt");
    // MULTI
    // ZREMRANGEBYSCORE {key} 0 {now - window_size}
    // ZADD {key} {now} {now}
    // ZCARD {key}
    // EXPIRE {key} {window_size}
    // EXEC

    // not rate limited
    Ok(false)
}
