package utils

import (
	"edu_v2/graph/model"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
)

func AbsResponseChecking(err error, msg string) (*model.Response, error) {
	if err != nil {
		return &model.Response{
			StatusCode: 409,
			Message:    err.Error(),
		}, nil
	}
	return &model.Response{
		StatusCode: 200,
		Message:    msg,
	}, nil
}

type Response struct {
	UserID int `json:"user_id"`
	Code   int `json:"code"`
}

func UploadQuestionImages(files []*graphql.Upload, name string) (model.Collection, error) {
	var collection model.Collection
	collection.Title = name

	// Use an absolute path to avoid issues with relative paths
	absPath := "/app/question_images"
	err := os.MkdirAll(absPath, os.ModePerm)
	if err != nil {
		return collection, fmt.Errorf("failed to create question_images directory: %w", err)
	}

	if files == nil {
		collection.ImageURL = nil
		return collection, nil
	}

	for _, file := range files {
		newFileName := uuid.New().String() + filepath.Ext(file.Filename)
		dstPath := filepath.Join(absPath, newFileName)

		dst, err := os.Create(dstPath)
		if err != nil {
			return collection, fmt.Errorf("failed to create file: %w", err)
		}

		if _, err := io.Copy(dst, file.File); err != nil {
			dst.Close()
			return collection, fmt.Errorf("failed to write file: %w", err)
		}

		if err := dst.Close(); err != nil {
			return collection, fmt.Errorf("failed to close file: %w", err)
		}

		collection.ImageURL = append(collection.ImageURL, newFileName)
	}

	return collection, nil
}
