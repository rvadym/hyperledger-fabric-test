package usecases


import (
	"errors"
	"github.com/rvadym/hyperledger-fabric-test/application/contracts"
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

type GetAllItemsInterface interface {
	ExecuteGetAllItems() ([]*domain.Item, error)
}

func NewGetAllItems(getItemsRepo contracts.GetItemsRepoInterface) GetAllItemsInterface {
	return &getAllItems{
		getItemsRepo,
	}
}

type getAllItems struct {
	getItemsRepo contracts.GetItemsRepoInterface
}

func (uc *getAllItems) ExecuteGetAllItems() ([]*domain.Item, error) {
	var err error
	var items []*domain.Item

	items, err = uc.getItemsRepo.GetAllItems()
	if err != nil {
		err = errors.New("failed to get all items")
		return nil, err
	}

	return items, err
}