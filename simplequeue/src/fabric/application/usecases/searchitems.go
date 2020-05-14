package usecases

import (
	"errors"
	"github.com/rvadym/hyperledger-fabric-test/application/contracts"
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

type SearchItemsInterface interface {
	ExecuteSearchItems(filter *contracts.Filter) ([]domain.Item, error)
}

func NewSearchItems(searchItemsRepo contracts.SearchItemsRepoInterface) SearchItemsInterface {
	return &searchItems{
		searchItemsRepo,
	}
}

type searchItems struct {
	searchItemsRepo contracts.SearchItemsRepoInterface
}

func (uc *searchItems) ExecuteSearchItems(filter *contracts.Filter) ([]domain.Item, error) {
	var err error
	var items []domain.Item

	items, err = uc.searchItemsRepo.SearchItems(
		filter.Limit,
		filter.Sort,
	)
	if err != nil {
		err = errors.New("failed to search for items")
		return nil, err
	}

	return items, err
}