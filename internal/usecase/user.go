package usecase

import "WB_Intern_L0/internal/repo"

type UserUseCase struct {
	repo repo.User
}

func NewUserUseCase(repo repo.User) *UserUseCase {
	return &UserUseCase{repo: repo}
}
