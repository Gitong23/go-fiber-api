package core

import (
	"context"
	"errors"
)

type UserService struct {
	repo IuserRepository
}

func NewUserService(repo IuserRepository) IuserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *User) error {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return errors.New("first name, last name and email are required")
	}
	return s.repo.Create(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *User) error {
	if user.ID.String() == "" {
		return errors.New("user ID is required")
	}
	return s.repo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
