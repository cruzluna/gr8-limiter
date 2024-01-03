use axum::{http::StatusCode, response::IntoResponse, routing::post, Router};

// use axum_macros::debug_handler;
use bb8_redis::{bb8, RedisConnectionManager};

// #[derive(Clone)]
// pub struct AppState {
//     pool: bb8::Pool<RedisConnectionManager>,
// }

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

    // let state = AppState { pool };

    let app = Router::new()
        .route("/api/v1/ratelimit", post(String::from("DEEZ")))
        .fallback(handler_404);
    // .with_state(state);
    // .with_state(state.clone());

    // run our app with hyper, listening on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
    // Ok(())
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
