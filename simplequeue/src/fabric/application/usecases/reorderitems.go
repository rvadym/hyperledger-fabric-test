package usecases

import (
	"errors"
	"fmt"
	"github.com/rvadym/hyperledger-fabric-test/domain"
	"strconv"
)

type ReorderItemsInterface interface {
	ExecuteReorderItems(itemId string, steps int) error
}

func NewReorderItems(
	getItemUC GetItemInterface,
	getAllItemsUC GetAllItemsInterface,
	saveItemUC SaveItemInterface,
) ReorderItemsInterface {
	return &reorderItems{
		getItemUC,
		getAllItemsUC,
		saveItemUC,
	}
}

type reorderItems struct {
	getItemUC     GetItemInterface
	getAllItemsUC GetAllItemsInterface
	saveItemUC    SaveItemInterface
}

func (uc *reorderItems) ExecuteReorderItems(itemId string, steps int) error {
	fmt.Println("About to move item with ID " + itemId + " " + strconv.Itoa(steps) + " step(s) up/down")
	var err error
	var item *domain.Item
	var items []*domain.Item
	var reordered bool

	item, err = uc.getItemUC.ExecuteGetItem(itemId)
	if err != nil {
		err = errors.New("failed to get item")
		return err
	}

	items, err = uc.getAllItemsUC.ExecuteGetAllItems()
	if err != nil {
		err = errors.New("failed to get all items")
		return err
	}

	items, reordered, err = uc.reorderItems(items, item, steps)
	if err != nil {
		err = errors.New("failed to reorder items")
		return err
	}

	if reordered == true {
		err = uc.saveItemUC.ExecuteSaveItems(items)
		if err != nil {
			err = errors.New("failed to save all items")
			return err
		}
	} else {
		fmt.Println("Info: Nothing to save, items stay in the same order.")
	}

	return nil
}

func (uc *reorderItems) reorderItems(items []*domain.Item, item *domain.Item, steps int) ([]*domain.Item, bool, error) {
	var err error
	itemFoundAtPosition := 0
	newItemPosition := 0
	itemFound := false
	reorderedItems := items
	reordered := false

	for currPos, currItem := range items {
		reorderedItems[currPos].Order = currPos

		if currItem.ID == item.ID {
			itemFoundAtPosition = currItem.Order
			itemFound = true
		}
	}

	// swap
	if itemFound == true {
		newItemPosition = itemFoundAtPosition + steps

		if newItemPosition < 0 {
			newItemPosition = 0
		}

		if newItemPosition > len(reorderedItems)-1 {
			newItemPosition = len(reorderedItems) - 1
		}

		if itemFoundAtPosition != newItemPosition {
			itemToSwapWith := reorderedItems[newItemPosition]
			itemToSwapWith.Order = item.Order
			reorderedItems[itemFoundAtPosition] = itemToSwapWith
			reorderedItems[newItemPosition] = item
			item.Order = newItemPosition

			reordered = true
		}
	}

	return reorderedItems, reordered, err
}
