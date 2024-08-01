-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;

-- name: GetPostByProjectID :one
SELECT * FROM posts WHERE project_id = $1;

-- name: GetReactionsByPostID :many
SELECT * FROM reactions WHERE post_id = $1;

-- name: GetCommentsByPostID :many
SELECT * FROM comments WHERE post_id = $1;

-- name: LoadOrderedPosts :many
SELECT
    p.id AS post_id,
    p.author_id,
    p.project_id,
    p.title,
    p.content AS post_content,
    p.repo_link,
    p.demo_link,
    p.post_image,
    p.like_count,
    p.comment_count,
    p.created_at AS post_created_at,
    p.updated_at AS post_updated_at,
    pf.name AS author_name,
    pf.image AS author_image
FROM
    posts p
JOIN accounts a ON p.author_id = a.id
JOIN profiles pf ON a.id = pf.account_id
ORDER BY
    p.created_at DESC
LIMIT $1
OFFSET $2;

-- name: LoadCommentsByPostID :many
SELECT
    c.id AS comment_id,
    c.post_id,
    c.parent_id,
    c.user_id,
    c.content AS comment_content,
    c.created_at AS comment_created_at,
    c.updated_at AS comment_updated_at,
    pf.name AS author_name,
    pf.image AS author_image
FROM
    comments c
JOIN accounts a ON c.user_id = a.id
JOIN profiles pf ON a.id = pf.account_id
WHERE
    c.post_id = $1
ORDER BY
    c.created_at DESC
LIMIT $2
OFFSET $3;

-- name: LoadReactionsByPostID :many
SELECT
    r.id AS reaction_id,
    r.post_id,
    r.comment_id,
    r.user_id,
    r.reaction_type,
    r.created_at AS reaction_created_at,
    pf.name AS author_name,
    pf.image AS author_image
FROM
    reactions r
JOIN accounts a ON r.user_id = a.id
JOIN profiles pf ON a.id = pf.account_id
WHERE
    r.post_id = $1
ORDER BY
    r.created_at DESC
LIMIT $2
OFFSET $3;


-- name: CreatePost :exec
INSERT INTO posts (id, project_id, author_id, title, content, repo_link, demo_link, post_image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: DeletePostById :exec
DELETE FROM posts WHERE id = $1;

-- name: DeletePostByProjectId :exec
DELETE FROM posts WHERE project_id = $1;

-- name: UpdatePost :exec
UPDATE posts SET title = $2, content = $3, repo_link = $4, demo_link = $5, post_image = $6, updated_at = $7 WHERE id = $1;

-- name: UpdatePostByProjectId :exec
UPDATE posts SET title = $2, content = $3, repo_link = $4, demo_link = $5, post_image = $6, updated_at = $7 WHERE project_id = $1;

-- name: CreateReaction :exec
INSERT INTO reactions (id, post_id, comment_id, user_id, reaction_type, created_at) VALUES ($1, $2, $3, $4, $5, $6);

-- name: AddReactionCount :exec
UPDATE posts SET like_count = like_count + 1 WHERE id = $1;

-- name: DeleteReactionById :exec
DELETE FROM reactions WHERE id = $1;

-- name: RemoveReactionCount :exec
UPDATE posts SET like_count = like_count - 1 WHERE id = $1;

-- name: CreateComment :exec
INSERT INTO comments (id, post_id, parent_id, user_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: DeleteCommentById :exec
DELETE FROM comments WHERE id = $1;

-- name: AddCommentCount :exec
UPDATE posts SET comment_count = comment_count + 1 WHERE id = $1;

-- name: RemoveCommentCount :exec
UPDATE posts SET comment_count = comment_count - 1 WHERE id = $1;