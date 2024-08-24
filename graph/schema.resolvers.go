package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"edu_v2/graph/model"
	utils "edu_v2/internal"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// CreateCollection is the resolver for the createCollection field.
func (r *mutationResolver) CreateCollection(ctx context.Context, name string, file []*graphql.Upload) (*model.Response, error) {
	collection, err := utils.UploadQuestionImages(file, name)
	if err != nil {
		return nil, err
	}
	err = r.CollService.CreateCollection(&collection)
	return utils.AbsResponseChecking(err, "collection created")
}

// UpdateCollection is the resolver for the updateCollection field.
func (r *mutationResolver) UpdateCollection(ctx context.Context, id string, name string, file []*graphql.Upload) (*model.Response, error) {
	return nil, nil
}

// DeleteCollection is the resolver for the deleteCollection field.
func (r *mutationResolver) DeleteCollection(ctx context.Context, id string) (*model.Response, error) {
	err := r.CollService.DeleteCollection(id)
	return utils.AbsResponseChecking(err, "Deleted")
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, name string, teacherName string, level model.GroupLevel, startAt string, startDate string, daysWeek model.DaysWeek) (*model.Response, error) {
	var group model.Group
	group.Name = name
	group.TeacherName = teacherName
	group.Level = level
	group.DaysWeek = daysWeek
	group.StartedDate = startDate
	group.StartAt = startAt
	err := r.GroupService.CreateGroup(&group)
	return utils.AbsResponseChecking(err, "Group added")
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, id string, name string, teacherName string, level model.GroupLevel, startAt string, startDate string, daysWeek model.DaysWeek) (*model.Response, error) {
	var group model.Group
	group.ID = id
	group.Name = name
	group.TeacherName = teacherName
	group.Level = level
	err := r.GroupService.UpdateGroup(&group)
	return utils.AbsResponseChecking(err, "updated")
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (*model.Response, error) {
	realId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	err = r.GroupService.DeleteGroup(realId)
	return utils.AbsResponseChecking(err, "deleted")
}

// CreateAnswer is the resolver for the createAnswer field.
func (r *mutationResolver) CreateAnswer(ctx context.Context, collectionID string, answers []*string, isUpdated *bool) (*model.Response, error) {
	return utils.AbsResponseChecking(r.AnswerService.CreateAnswer(answers, isUpdated, &collectionID), "Answer Created")
}

// DeleteAnswer is the resolver for the deleteAnswer field.
func (r *mutationResolver) DeleteAnswer(ctx context.Context, collectionID string) (*model.Response, error) {
	return utils.AbsResponseChecking(r.AnswerService.DeleteAnswer(&collectionID), "Answer Deleted")
}

// CreateStudentAnswer is the resolver for the createStudentAnswer field.
func (r *mutationResolver) CreateStudentAnswer(ctx context.Context, collectionID string, answers []*string) (*model.Response, error) {
	panic(fmt.Errorf("not implemented: CreateStudentAnswer - createStudentAnswer"))
}

// GetGroups is the resolver for the getGroups field.
func (r *queryResolver) GetGroups(ctx context.Context, byID *string, orderByLevel *bool) ([]*model.Group, error) {
	group, err := r.GroupService.GetGroup(byID, orderByLevel)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// GetCollection is the resolver for the getCollection field.
func (r *queryResolver) GetCollection(ctx context.Context) ([]*model.Collection, error) {
	collections, err := r.CollService.GetCollections()
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// GetCollectionByID is the resolver for the getCollectionById field.
func (r *queryResolver) GetCollectionByID(ctx context.Context, collectionID string) (*model.Collection, error) {
	collection, err := r.CollService.GetCollection(collectionID)
	if err != nil {
		return nil, err
	}
	return collection, err
}

// GetStudentTestExams is the resolver for the getStudentTestExams field.
func (r *queryResolver) GetStudentTestExams(ctx context.Context, code *string, studentID *string, page *int, size *int) ([]*model.UserCollectionTestExams, error) {
	panic(fmt.Errorf("not implemented: GetStudentTestExams - getStudentTestExams"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
