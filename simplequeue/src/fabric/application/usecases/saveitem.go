package usecases

import (
	"github.com/google/uuid"
	"github.com/rvadym/hyperledger-fabric-test/application/contracts"
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

type SaveItemInterface interface {
	ExecuteSaveItem(ID string, order int, content string) (*domain.Item, error)
	ExecuteSaveItems(items []*domain.Item) error
}

func NewSaveItem(saveItemRepo contracts.SaveItemRepoInterface) SaveItemInterface {
	return &saveItem{
		saveItemRepo,
	}
}

type saveItem struct {
	saveItemRepo contracts.SaveItemRepoInterface
}

func (uc *saveItem) ExecuteSaveItem(ID string, order int, content string) (*domain.Item, error) {
	if ID == "" {
		ID = uuid.New().String()
	}

	item := &domain.Item{
		ObjectType: "item",
		ID:         ID,
		Content:    content,
		Order:      order,
	}

	return item, uc.saveItemRepo.SaveItem(item)
}

func (uc *saveItem) ExecuteSaveItems(items []*domain.Item) error {
	return uc.saveItemRepo.SaveItems(items)
}
