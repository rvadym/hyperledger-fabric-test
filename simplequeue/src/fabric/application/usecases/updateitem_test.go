package usecases

import (
	"github.com/golang/mock/gomock"
	"github.com/rvadym/hyperledger-fabric-test/domain"
	"github.com/rvadym/hyperledger-fabric-test/tests/mock"
	"testing"
)

func TestNewUpdateItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSaveItemRepo := mock.NewMockSaveItemRepoInterface(ctrl)
	mockSaveItemRepo.
		EXPECT().
		SaveItem(gomock.Any()).
		Return(nil).
		Times(1)

	updateItem := NewUpdateItem(mockSaveItemRepo)
	err := updateItem.ExecuteUpdateItem(&domain.Item{ID: "ID", Content: ""})

	if err != nil {
		t.Errorf("Error is not nil.")
	}
}
