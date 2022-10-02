package repository

import (
	database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	commenRepository := NewCommentRepository(db)

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@gmail.com",
		Comment: "Test Repository",
	}

	result, err := commenRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	commenRepository := NewCommentRepository(db)

	comment, err := commenRepository.FindById(context.Background(), 90)

	if err != nil {
		panic(err)
	}
	fmt.Print(comment)
}

func TestFindAll(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	commenRepository := NewCommentRepository(db)

	comments, err := commenRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Print(comments)
}
