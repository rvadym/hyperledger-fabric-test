package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/rvadym/hyperledger-fabric-test/domain"
	"github.com/rvadym/hyperledger-fabric-test/tests/mock"
)

func TestNewSaveItem(t *testing.T) {
	var err error
	var item *domain.Item
	var expectedId string = "ID"
	var expectedContent string = "Some content"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSaveItemRepo := mock.NewMockSaveItemRepoInterface(ctrl)
	mockSaveItemRepo.
		EXPECT().
		SaveItem(gomock.Any()).
		Return(nil).
		Times(1)

	saveItem := NewSaveItem(mockSaveItemRepo)
	item, err = saveItem.ExecuteSaveItem(expectedId, expectedContent)

	if err != nil {
		t.Errorf("Error is not nil")
	}

	if item.ID != expectedId {
		t.Errorf("ID value %s doesn't match %s", item.ID, expectedId)
	}

	if item.Content != expectedContent {
		t.Errorf("ID value %s doesn't match %s", item.Content, expectedContent)
	}
}
