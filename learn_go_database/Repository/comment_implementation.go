package repository

import (
	"context"
	"database/sql"
	"errors"
	entity "learn_go_database/Entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}


func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.CommentEntity) (entity.CommentEntity, error) {
	sqlQuery := "INSERT INTO comment (email, comment) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, sqlQuery, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, err
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.CommentEntity, error) {
	sqlQuery := "SELECT id, email, comment FROM comment WHERE id=? LIMIT 1"
	result, err := repo.DB.QueryContext(ctx, sqlQuery, id)
	comment := entity.CommentEntity{}
	if err != nil {
		return comment, err
	}

	defer result.Close()

	if result.Next() {
		result.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repo *commentRepositoryImpl) FIndAll(ctx context.Context) ([]entity.CommentEntity, error) {
	sqlQuery := "SELECT id, email, comment FROM comment"
	var comments []entity.CommentEntity

	result, err := repo.DB.QueryContext(ctx, sqlQuery)

	if err != nil {
		return comments, err
	}

	defer result.Close()

	for result.Next() {
		comment := entity.CommentEntity{}
		result.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil

}
