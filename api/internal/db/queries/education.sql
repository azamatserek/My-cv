-- name: ListEducation :many
SELECT * FROM education ORDER BY order_index ASC, end_year DESC;