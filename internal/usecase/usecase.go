package usecase

import "WB_Intern_L0/internal/repo"

type User interface {
}

type UseCase struct {
	User
}

func NewUseCase(repo *repo.Repository) *UseCase {
	return &UseCase{
		User: NewUserUseCase(repo.User),
	}
}
