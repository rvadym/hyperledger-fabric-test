package usecases

import (
	"errors"
	"github.com/rvadym/hyperledger-fabric-test/application/contracts"
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

type GetItemInterface interface {
	ExecuteGetItem(ID string) (*domain.Item, error)
	ExecuteGetFirstItem() (*domain.Item, error)
	ExecuteGetLastItem() (*domain.Item, error)
}

func NewGetItem(getItemRepo contracts.GetItemRepoInterface) GetItemInterface {
	return &getItem{
		getItemRepo,
	}
}

type getItem struct {
	getItemRepo contracts.GetItemRepoInterface
}

func (uc *getItem) ExecuteGetItem(id string) (*domain.Item, error) {
	var err error
	var item *domain.Item

	if id == "" {
		err = errors.New("id cannot be empty")
		return item, err
	}

	item, err = uc.getItemRepo.GetItem(id)
	if err != nil {
		err = errors.New("failed to get item with ID " + id)
		return nil, err
	}
	if item == nil {
		err = errors.New("item with ID " + id + " not found")
		return nil, err
	}

	return item, err
}

func (uc *getItem) ExecuteGetFirstItem() (*domain.Item, error) {
	var err error
	var item *domain.Item

	item, err = uc.getItemRepo.GetFirstItem()
	if err != nil {
		err = errors.New("failed to get first item")
		return nil, err
	}
	if item == nil {
		err = errors.New("first item not found")
		return nil, err
	}

	return item, err
}

func (uc *getItem) ExecuteGetLastItem() (*domain.Item, error) {
	var err error
	var item *domain.Item

	item, err = uc.getItemRepo.GetLastItem()
	if err != nil {
		err = errors.New("failed to get last item")
		return nil, err
	}
	if item == nil {
		err = errors.New("last item not found")
		return nil, err
	}

	return item, err
}
