package repository

import (
	"context"
	entity "learn_go_database/Entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.CommentEntity) (entity.CommentEntity, error)
	FindById(ctx context.Context, id int32) (entity.CommentEntity, error)
	FIndAll(ctx context.Context) ([]entity.CommentEntity, error)
}
