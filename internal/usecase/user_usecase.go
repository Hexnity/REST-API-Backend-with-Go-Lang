package usecase

import (
	"context"
	"hexnity/internal/domain"
)

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) domain.UserUseCase {
	return &userUseCase{repo: r}
}

func (uc *userUseCase) Register(ctx context.Context, u *domain.User) error {
	// Logic: You could check if email is valid or user is banned here
	return uc.repo.Create(ctx, u)
}
