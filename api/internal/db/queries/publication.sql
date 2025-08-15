-- name: ListPublications :many
SELECT * FROM publication ORDER BY year DESC, order_index ASC;