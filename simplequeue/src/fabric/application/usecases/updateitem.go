package usecases

import (
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

type UpdateItemInterface interface {
	ExecuteUpdateItem(id, content string) (*domain.Item, error)
}

func NewUpdateItem(getItemUC GetItemInterface, saveItemUC SaveItemInterface) UpdateItemInterface {
	return &updateItem{
		getItemUC,
		saveItemUC,
	}
}

type updateItem struct {
	getItemUC  GetItemInterface
	saveItemUC SaveItemInterface
}

func (uc *updateItem) ExecuteUpdateItem(id, content string) (*domain.Item, error) {
	var err error
	var item *domain.Item

	item, err = uc.getItemUC.ExecuteGetItem(id)
	if err != nil {
		return nil, err
	}

	item, err = uc.saveItemUC.ExecuteSaveItem(item.ID, item.Order, content)
	if err != nil {
		return nil, err
	}

	return item, err
}
