package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
)

type ItemsUseCase struct {
	repo repo.Items
}

func (p ItemsUseCase) NewItem(item entity.Item, orderID string) error {
	return p.repo.CreateItem(item, orderID)
}

func (p ItemsUseCase) GetItemsByOrderUID(orderUID string) ([]entity.Item, error) {
	return p.repo.GetItemsByOrderUID(orderUID)
}

func NewItemsUseCase(repo repo.Items) *ItemsUseCase {
	return &ItemsUseCase{repo: repo}
}
