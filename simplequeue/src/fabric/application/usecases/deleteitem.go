package usecases

import (
	"errors"
	"github.com/rvadym/hyperledger-fabric-test/application/contracts"
)

type DeleteItemInterface interface {
	ExecuteDeleteItem(id string) error
}

func NewDeleteItem(deleteItemRepo contracts.DeleteItemRepoInterface) DeleteItemInterface {
	return &deleteItem{
		deleteItemRepo,
	}
}

type deleteItem struct {
	deleteItemRepo contracts.DeleteItemRepoInterface
}

func (uc *deleteItem) ExecuteDeleteItem(id string) error {
	var err error

	if id == "" {
		err = errors.New("ID cannot be empty")
		return err
	}

	return uc.deleteItemRepo.DeleteItem(id)
}