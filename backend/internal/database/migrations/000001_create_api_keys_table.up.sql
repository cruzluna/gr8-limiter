CREATE TABLE IF NOT EXISTS api_keys (
    id SERIAL PRIMARY KEY,
    api_key UUID NOT NULL,
    user_id text NOT NULL,
    created_at timestamp DEFAULT current_timestamp
);
