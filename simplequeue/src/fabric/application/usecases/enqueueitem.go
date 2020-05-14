package usecases

import (
	"fmt"
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

type EnqueueItemInterface interface {
	ExecuteEnqueueItem(content string) (*domain.Item, error)
}

func NewEnqueueItem(getItemUC GetItemInterface, saveItemUC SaveItemInterface) EnqueueItemInterface {
	return &enqueueItem{
		getItemUC,
		saveItemUC,
	}
}

type enqueueItem struct {
	getItemUC  GetItemInterface
	saveItemUC SaveItemInterface
}

func (uc *enqueueItem) ExecuteEnqueueItem(content string) (*domain.Item, error) {
	var err error
	var lastItem *domain.Item
	var newOrder = 0

	lastItem, err = uc.getItemUC.ExecuteGetLastItem()
	if err != nil {
		fmt.Println("Error getting last item " + err.Error())
		// FIXME better error handling
		if err.Error() != "failed to get last item" {
			return nil, err
		}

		err = nil
	}

	if lastItem != nil {
		fmt.Println("Last item is nil")
		newOrder = lastItem.Order + 1
	}

	lastItem, err = uc.saveItemUC.ExecuteSaveItem("", newOrder, content)
	if err != nil {
		return nil, err
	}

	return lastItem, err
}
