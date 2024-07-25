-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;

-- name: GetOrderedPosts :many
SELECT * FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: CreatePost :one
INSERT INTO posts (id, project_id, author_id, title, content, repo_link, demo_link, post_image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;

-- name: UpdatePost :one
UPDATE posts SET title = $2, content = $3, repo_link = $4, demo_link = $5, post_image = $6, updated_at = $7 WHERE id = $1 RETURNING *;