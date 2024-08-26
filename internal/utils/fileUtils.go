package utils

import (
	"edu_v2/graph/model"
	"github.com/99designs/gqlgen/graphql"
)

func UploadQuestionfile(file graphql.Upload, name string) (model.Collection, error) {

	return model.Collection{
		Title:     name,
		Questions: "",
	}, nil
}
