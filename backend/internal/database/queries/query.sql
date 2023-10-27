-- name: GetApiKeys :many
SELECT * FROM api_keys;

-- name: InsertApiKey :exec
INSERT INTO api_keys (api_key, user_id) VALUES (@apiKey, @userId);

-- name: DeleteByUser :exec
DELETE FROM api_keys WHERE user_id = @userId;

-- name: DeleteByApiKey :exec
DELETE FROM api_keys WHERE api_key = @apiKey;

-- name: DeleteByUserAndApiKey :exec
DELETE FROM api_keys WHERE user_id = @userId;


