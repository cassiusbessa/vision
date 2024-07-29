// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addCommentCount = `-- name: AddCommentCount :exec
UPDATE posts SET comment_count = comment_count + 1 WHERE id = $1
`

func (q *Queries) AddCommentCount(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, addCommentCount, id)
	return err
}

const addReactionCount = `-- name: AddReactionCount :exec
UPDATE posts SET like_count = like_count + 1 WHERE id = $1
`

func (q *Queries) AddReactionCount(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, addReactionCount, id)
	return err
}

const createComment = `-- name: CreateComment :exec
INSERT INTO comments (id, post_id, parent_id, user_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateCommentParams struct {
	ID        uuid.UUID
	PostID    uuid.UUID
	ParentID  uuid.NullUUID
	UserID    uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) error {
	_, err := q.db.Exec(ctx, createComment,
		arg.ID,
		arg.PostID,
		arg.ParentID,
		arg.UserID,
		arg.Content,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createPost = `-- name: CreatePost :exec
INSERT INTO posts (id, project_id, author_id, title, content, repo_link, demo_link, post_image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type CreatePostParams struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	AuthorID  uuid.UUID
	Title     string
	Content   string
	RepoLink  sql.NullString
	DemoLink  sql.NullString
	PostImage sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) error {
	_, err := q.db.Exec(ctx, createPost,
		arg.ID,
		arg.ProjectID,
		arg.AuthorID,
		arg.Title,
		arg.Content,
		arg.RepoLink,
		arg.DemoLink,
		arg.PostImage,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createReaction = `-- name: CreateReaction :exec
INSERT INTO reactions (id, post_id, comment_id, user_id, reaction_type, created_at) VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateReactionParams struct {
	ID           uuid.UUID
	PostID       uuid.UUID
	CommentID    uuid.NullUUID
	UserID       uuid.UUID
	ReactionType string
	CreatedAt    time.Time
}

func (q *Queries) CreateReaction(ctx context.Context, arg CreateReactionParams) error {
	_, err := q.db.Exec(ctx, createReaction,
		arg.ID,
		arg.PostID,
		arg.CommentID,
		arg.UserID,
		arg.ReactionType,
		arg.CreatedAt,
	)
	return err
}

const deleteCommentById = `-- name: DeleteCommentById :exec
DELETE FROM comments WHERE id = $1
`

func (q *Queries) DeleteCommentById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCommentById, id)
	return err
}

const deleteReactionById = `-- name: DeleteReactionById :exec
DELETE FROM reactions WHERE id = $1
`

func (q *Queries) DeleteReactionById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteReactionById, id)
	return err
}

const getCommentsByPostID = `-- name: GetCommentsByPostID :many
SELECT id, post_id, parent_id, user_id, content, created_at, updated_at FROM comments WHERE post_id = $1
`

func (q *Queries) GetCommentsByPostID(ctx context.Context, postID uuid.UUID) ([]Comment, error) {
	rows, err := q.db.Query(ctx, getCommentsByPostID, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.ParentID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostByID = `-- name: GetPostByID :one
SELECT id, project_id, author_id, title, content, repo_link, demo_link, post_image, like_count, comment_count, created_at, updated_at FROM posts WHERE id = $1
`

func (q *Queries) GetPostByID(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRow(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.AuthorID,
		&i.Title,
		&i.Content,
		&i.RepoLink,
		&i.DemoLink,
		&i.PostImage,
		&i.LikeCount,
		&i.CommentCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReactionsByPostID = `-- name: GetReactionsByPostID :many
SELECT id, post_id, comment_id, user_id, reaction_type, created_at FROM reactions WHERE post_id = $1
`

func (q *Queries) GetReactionsByPostID(ctx context.Context, postID uuid.UUID) ([]Reaction, error) {
	rows, err := q.db.Query(ctx, getReactionsByPostID, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reaction
	for rows.Next() {
		var i Reaction
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.CommentID,
			&i.UserID,
			&i.ReactionType,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const loadOrderedPosts = `-- name: LoadOrderedPosts :many
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
    c.id AS comment_id,
    c.post_id AS comment_post_id,
    c.parent_id AS comment_parent_id,
    c.user_id AS comment_user_id,
    c.content AS comment_content,
    c.created_at AS comment_created_at,
    c.updated_at AS comment_updated_at,
    r.id AS reaction_id,
    r.post_id AS reaction_post_id,
    r.comment_id AS reaction_comment_id,
    r.user_id AS reaction_user_id,
    r.reaction_type,
    r.created_at AS reaction_created_at
FROM
    posts p
LEFT JOIN comments c ON p.id = c.post_id
LEFT JOIN reactions r ON p.id = r.post_id
ORDER BY
    p.created_at DESC,
    c.created_at,
    r.created_at
`

type LoadOrderedPostsRow struct {
	PostID            uuid.UUID
	AuthorID          uuid.UUID
	ProjectID         uuid.UUID
	Title             string
	PostContent       string
	RepoLink          sql.NullString
	DemoLink          sql.NullString
	PostImage         sql.NullString
	LikeCount         int32
	CommentCount      int32
	PostCreatedAt     time.Time
	PostUpdatedAt     time.Time
	CommentID         uuid.NullUUID
	CommentPostID     uuid.NullUUID
	CommentParentID   uuid.NullUUID
	CommentUserID     uuid.NullUUID
	CommentContent    sql.NullString
	CommentCreatedAt  sql.NullTime
	CommentUpdatedAt  sql.NullTime
	ReactionID        uuid.NullUUID
	ReactionPostID    uuid.NullUUID
	ReactionCommentID uuid.NullUUID
	ReactionUserID    uuid.NullUUID
	ReactionType      sql.NullString
	ReactionCreatedAt sql.NullTime
}

func (q *Queries) LoadOrderedPosts(ctx context.Context) ([]LoadOrderedPostsRow, error) {
	rows, err := q.db.Query(ctx, loadOrderedPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LoadOrderedPostsRow
	for rows.Next() {
		var i LoadOrderedPostsRow
		if err := rows.Scan(
			&i.PostID,
			&i.AuthorID,
			&i.ProjectID,
			&i.Title,
			&i.PostContent,
			&i.RepoLink,
			&i.DemoLink,
			&i.PostImage,
			&i.LikeCount,
			&i.CommentCount,
			&i.PostCreatedAt,
			&i.PostUpdatedAt,
			&i.CommentID,
			&i.CommentPostID,
			&i.CommentParentID,
			&i.CommentUserID,
			&i.CommentContent,
			&i.CommentCreatedAt,
			&i.CommentUpdatedAt,
			&i.ReactionID,
			&i.ReactionPostID,
			&i.ReactionCommentID,
			&i.ReactionUserID,
			&i.ReactionType,
			&i.ReactionCreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeCommentCount = `-- name: RemoveCommentCount :exec
UPDATE posts SET comment_count = comment_count - 1 WHERE id = $1
`

func (q *Queries) RemoveCommentCount(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, removeCommentCount, id)
	return err
}

const removeReactionCount = `-- name: RemoveReactionCount :exec
UPDATE posts SET like_count = like_count - 1 WHERE id = $1
`

func (q *Queries) RemoveReactionCount(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, removeReactionCount, id)
	return err
}

const updatePost = `-- name: UpdatePost :exec
UPDATE posts SET title = $2, content = $3, repo_link = $4, demo_link = $5, post_image = $6, updated_at = $7 WHERE id = $1
`

type UpdatePostParams struct {
	ID        uuid.UUID
	Title     string
	Content   string
	RepoLink  sql.NullString
	DemoLink  sql.NullString
	PostImage sql.NullString
	UpdatedAt time.Time
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.Exec(ctx, updatePost,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.RepoLink,
		arg.DemoLink,
		arg.PostImage,
		arg.UpdatedAt,
	)
	return err
}
