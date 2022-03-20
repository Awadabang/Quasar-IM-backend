-- name: CreateUser :execresult
INSERT INTO users (username, avatar, hashed_password) VALUES (?,?,?);

-- name: GetUserByName :one
SELECT * FROM users WHERE username = ? LIMIT 1;