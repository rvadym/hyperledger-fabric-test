package usecases

import (
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

type DequeueItemInterface interface {
	ExecuteDequeueItem() (*domain.Item, error)
}

func NewDequeueItem(getItemUC GetItemInterface, deleteItemUC DeleteItemInterface) DequeueItemInterface {
	return &dequeueItem{
		getItemUC,
		deleteItemUC,
	}
}

type dequeueItem struct {
	getItemUC GetItemInterface
	deleteItemUC DeleteItemInterface
}

func (uc *dequeueItem) ExecuteDequeueItem() (*domain.Item, error) {
	var err error
	var lastItem *domain.Item

	lastItem, err = uc.getItemUC.ExecuteGetFirstItem()
	if err != nil {
		return nil, err
	}

	err = uc.deleteItemUC.ExecuteDeleteItem(lastItem.ID)
	if err != nil {
		return nil, err
	}

	return lastItem, err
}