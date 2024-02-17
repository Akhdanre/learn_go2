package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	lgd "learn_go_database"
	entity "learn_go_database/Entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(lgd.GetConnection())

	ctx := context.Background()
	comment := entity.CommentEntity{
		Email:   "testRepoins@gmail.com",
		Comment: "test repository insert",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindByID(t *testing.T) {
	CommentRepository := NewCommentRepository(lgd.GetConnection())

	ctx := context.Background()
	result, err := CommentRepository.FindById(ctx, 57)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentAll(t *testing.T) {
	CommentRepository := NewCommentRepository(lgd.GetConnection())

	ctx := context.Background()
	results, err := CommentRepository.FIndAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, result := range results{
		fmt.Println(result)
	}
}
