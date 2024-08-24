package graph

import "edu_v2/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GroupService  *service.GroupService
	CollService   *service.CollectionService
	AnswerService *service.AnswerService
}
