package service

import (
	"edu_v2/graph/model"
	"edu_v2/internal/repository"
)

type GroupService struct {
	repo *repository.GroupRepository
}

func NewGroupService(repo *repository.GroupRepository) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) CreateGroup(group *model.Group) error {
	return s.repo.Create(group)
}

func (s *GroupService) GetGroup(id *string, orderLevel *bool, page, size *int) ([]*model.Group, error) {
	return s.repo.Get(id, orderLevel)
}

func (s *GroupService) UpdateGroup(group *model.Group) error {
	return s.repo.Update(group)
}

func (s *GroupService) DeleteGroup(id int) error {
	return s.repo.Delete(id)
}
