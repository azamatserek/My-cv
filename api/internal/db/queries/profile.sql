-- name: GetProfile :one
SELECT * FROM profile ORDER BY id DESC LIMIT 1;

-- name: UpsertProfile :one
INSERT INTO profile (full_name, title, email, phone, location, about, avatar_url)
VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT (id) DO NOTHING
RETURNING *;