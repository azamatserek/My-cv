-- name: ListExperience :many
SELECT * FROM experience ORDER BY order_index ASC, start_date DESC;