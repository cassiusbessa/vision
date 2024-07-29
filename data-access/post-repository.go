package data

import (
	"context"

	"github.com/cassiusbessa/vision-social-media/data-access/mappers"
	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type PostRepository struct {
	queries *sqlc.Queries
	db      *pgx.Conn
}

func NewPostRepository(queries *sqlc.Queries, db *pgx.Conn) *PostRepository {
	return &PostRepository{
		queries: queries,
		db:      db,
	}
}

func (repo *PostRepository) SavePost(post *entities.ProjectPost) error {
	err := repo.queries.CreatePost(context.Background(), mappers.PostEntityToCreateQueryParams(post))
	return err
}

func (repo *PostRepository) RemovePost(postID uuid.UUID) (bool, error) {
	err := repo.queries.DeletePostById(context.Background(), postID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *PostRepository) UpdatePost(post *entities.ProjectPost) error {
	err := repo.queries.UpdatePost(context.Background(), mappers.PostEntityToUpdateQueryParams(post))
	return err
}

func (repo *PostRepository) GetPostByID(postID uuid.UUID) (*entities.ProjectPost, error) {

	post, err := repo.queries.GetPostByID(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	comments, err := repo.queries.GetCommentsByPostID(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	reactions, err := repo.queries.GetReactionsByPostID(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	return mappers.PostDBEntityToProjectPost(post, comments, reactions), nil
}

func (repo *PostRepository) LoadOrderedPosts() ([]entities.ProjectPost, error) {
	posts, err := repo.queries.LoadOrderedPosts(context.Background())
	if err != nil {
		return nil, err
	}

	entitiesPosts := map[uuid.UUID]*entities.ProjectPost{}

	for _, post := range posts {
		if _, ok := entitiesPosts[post.PostID]; !ok {
			entitiesPosts[post.PostID] = mappers.LoadOrderedPostRowToProjectPosts(post)
		}

		if post.CommentID.Valid {
			entitiesPosts[post.PostID].AddComment(mappers.LoadOrderedPostRowToPostComment(post))
		}

		if post.ReactionID.Valid {
			entitiesPosts[post.PostID].AddReaction(mappers.LoadOrderedPostRowToPostReaction(post))
		}
	}

	result := make([]entities.ProjectPost, 0, len(entitiesPosts))
	for _, post := range entitiesPosts {
		result = append(result, *post)
	}

	return result, nil
}

func (repo *PostRepository) AddReactionToPost(reaction *entities.Reaction) error {

	return withTransaction(context.Background(), repo.db, func(ctx context.Context, qtx *sqlc.Queries) error {
		err := qtx.CreateReaction(ctx, mappers.ReactionEntityToCreateReactionQueryParams(reaction))
		if err != nil {
			return err
		}

		err = qtx.AddReactionCount(ctx, reaction.PostID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (repo *PostRepository) RemovePostReaction(reactionID, postID uuid.UUID) (bool, error) {
	return true, withTransaction(context.Background(), repo.db, func(ctx context.Context, qtx *sqlc.Queries) error {
		err := qtx.DeleteReactionById(ctx, reactionID)
		if err != nil {
			return err
		}

		err = qtx.RemoveReactionCount(ctx, postID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (repo *PostRepository) AddCommentToPost(comment *entities.Comment) error {
	return withTransaction(context.Background(), repo.db, func(ctx context.Context, qtx *sqlc.Queries) error {
		err := qtx.CreateComment(ctx, mappers.CommentEntityToCreateQueryParams(comment))
		if err != nil {
			return err
		}

		err = qtx.AddCommentCount(ctx, comment.PostID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (repo *PostRepository) RemovePostComment(commentID, postID uuid.UUID) (bool, error) {
	return true, withTransaction(context.Background(), repo.db, func(ctx context.Context, qtx *sqlc.Queries) error {
		err := qtx.DeleteCommentById(ctx, commentID)
		if err != nil {
			return err
		}

		err = qtx.RemoveCommentCount(ctx, postID)
		if err != nil {
			return err
		}

		return nil
	})
}

func withTransaction(ctx context.Context, db *pgx.Conn, fn func(context.Context, *sqlc.Queries) error) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			tx.Rollback(ctx)
		}
	}()

	qtx := sqlc.New(tx)

	err = fn(ctx, qtx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
