package contracts

import (
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

type SaveItemRepoInterface interface {
	SaveItem(item *domain.Item) error
	SaveItems(items []*domain.Item) error
}

type GetItemRepoInterface interface {
	GetItem(id string) (*domain.Item, error)
	GetFirstItem() (*domain.Item, error)
	GetLastItem() (*domain.Item, error)
}

type DeleteItemRepoInterface interface {
	DeleteItem(id string) error
}

type GetItemsRepoInterface interface {
	GetAllItems() ([]*domain.Item, error)
}

type SearchItemsRepoInterface interface {
	SearchItems(search string, perPage int32, sort string) ([]*domain.Item, error)
}
