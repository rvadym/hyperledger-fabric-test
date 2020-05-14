package usecases

import (
	"github.com/golang/mock/gomock"
	"github.com/rvadym/hyperledger-fabric-test/application/contracts"
	"github.com/rvadym/hyperledger-fabric-test/domain"
	"github.com/rvadym/hyperledger-fabric-test/tests/mock"
	"testing"
)

func TestNewGetItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var ids []string
	ids = append(ids, "SOME_ID")

	var expectedItems []*domain.Item
	expectedItems = append(
		expectedItems,
		&domain.Item{ID: "ID", Content: ""},
	)

	mockFetchItemRepo := mock.NewMockFetchItemsRepoInterface(ctrl)
	mockFetchItemRepo.
		EXPECT().
		FetchItems(gomock.Any()).
		Return(
			expectedItems,
			nil,
		).
		Times(1)

	getItems := NewGetItems(mockFetchItemRepo)
	items, err := getItems.ExecuteGetItems(
		&contracts.Filter{
			IDs: ids,
		},
	)

	if err != nil {
		t.Errorf("Error is not nil.")
	}


	if items[0] != expectedItems[0] {
		t.Errorf("Items are not equal.")
	}
}