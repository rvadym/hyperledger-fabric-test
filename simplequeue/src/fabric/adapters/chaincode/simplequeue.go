package chaincode

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/rvadym/hyperledger-fabric-test/adapters/repo"
	"github.com/rvadym/hyperledger-fabric-test/application/contracts"
	"github.com/rvadym/hyperledger-fabric-test/application/usecases"
	"github.com/rvadym/hyperledger-fabric-test/domain"
)

// SimpleQueue Chaincode implementation
type SimpleQueue struct {
	contractapi.Contract
}

func (t *SimpleQueue) Init(ctx contractapi.TransactionContextInterface) error {
	return nil
}

func (t *SimpleQueue) Enqueue(ctx contractapi.TransactionContextInterface, content string) (string, error) {
	fmt.Println("SimpleQueue > Enqueue item")
	var err error
	var item *domain.Item

	itemsRepo := &repo.ItemRepo{Ctx: ctx}
	getItemUC := usecases.NewGetItem(itemsRepo)
	saveItemUC := usecases.NewSaveItem(itemsRepo)
	enqueueItemUC := usecases.NewEnqueueItem(getItemUC, saveItemUC)

	item, err = enqueueItemUC.ExecuteEnqueueItem(content)
	if err != nil {
		jsonResp := "{\"Error\": \"" + err.Error() + "\"}"
		return "", errors.New(jsonResp)
	}

	itemJson, err := json.Marshal(item)
	if err != nil {
		fmt.Println("SimpleQueue.Enqueue (1) " + err.Error())
		jsonResp := "{\"Error\":\"Item was not queued\"}"
		return "", errors.New(jsonResp)
	}

	return string(itemJson), nil
}

func (t *SimpleQueue) Dequeue(ctx contractapi.TransactionContextInterface) (string, error) {
	fmt.Println("SimpleQueue > Dequeue item")
	var err error
	var item *domain.Item

	itemsRepo := &repo.ItemRepo{Ctx: ctx}
	getItemUC := usecases.NewGetItem(itemsRepo)
	deleteItemUC := usecases.NewDeleteItem(itemsRepo)
	dequeueItemUC := usecases.NewDequeueItem(getItemUC, deleteItemUC)

	item, err = dequeueItemUC.ExecuteDequeueItem()
	if err != nil {
		jsonResp := "{\"Error\": \"" + err.Error() + "\"}"
		return "", errors.New(jsonResp)
	}

	itemJson, err := json.Marshal(item)
	if err != nil {
		fmt.Println("SimpleQueue.Dequeue (1) " + err.Error())
		jsonResp := "{\"Error\":\"Item was not dequeued\"}"
		return "", errors.New(jsonResp)
	}

	return string(itemJson), nil
}

func (t *SimpleQueue) Move(ctx contractapi.TransactionContextInterface, id string, steps int) error {
	fmt.Println("SimpleQueue > Move item")
	var err error

	itemsRepo := &repo.ItemRepo{Ctx: ctx}

	getItemUC := usecases.NewGetItem(itemsRepo)
	getAllItemsUC := usecases.NewGetAllItems(itemsRepo)
	saveItemUC := usecases.NewSaveItem(itemsRepo)
	reorderItemsUC := usecases.NewReorderItems(getItemUC, getAllItemsUC, saveItemUC)

	err = reorderItemsUC.ExecuteReorderItems(id, steps)

	return err
}

func (t *SimpleQueue) Get(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	fmt.Println("SimpleQueue > Get item from queue")
	var err error
	var item *domain.Item

	getItemsRepo := &repo.ItemRepo{Ctx: ctx}
	addItemUC := usecases.NewGetItem(getItemsRepo)

	item, err = addItemUC.ExecuteGetItem(id)
	if err != nil {
		jsonResp := "{\"Error\": \"" + err.Error() + "\"}"
		return "", errors.New(jsonResp)
	}

	jsonResp, err := json.Marshal(item)
	if err != nil {
		fmt.Println("SimpleQueue.Get (1) " + err.Error())
		jsonResp := "{\"Error\":\"Item with ID  " + id + " is corrupted\"}"
		return "", errors.New(jsonResp)
	}

	return string(jsonResp), err
}

func (t *SimpleQueue) Update(ctx contractapi.TransactionContextInterface, id, content string) (string, error) {
	fmt.Println("SimpleQueue > Update item in queue")
	var err error
	var item *domain.Item

	itemsRepo := &repo.ItemRepo{Ctx: ctx}

	getItemUC := usecases.NewGetItem(itemsRepo)
	saveItemUC := usecases.NewSaveItem(itemsRepo)
	updateItemUC := usecases.NewUpdateItem(getItemUC, saveItemUC)

	item, err = updateItemUC.ExecuteUpdateItem(id, content)
	if err != nil {
		jsonResp := "{\"Error\": \"" + err.Error() + "\"}"
		return "", errors.New(jsonResp)
	}

	jsonResp, err := json.Marshal(item)
	if err != nil {
		fmt.Println("SimpleQueue.Update (1) " + err.Error())
		jsonResp := "{\"Error\":\"Item with ID  " + id + " is corrupted\"}"
		return "", errors.New(jsonResp)
	}

	return string(jsonResp), err
}

func (t *SimpleQueue) Delete(ctx contractapi.TransactionContextInterface, id string) error {
	fmt.Println("SimpleQueue > Delete item from queue")
	var err error

	deleteRepo := &repo.ItemRepo{Ctx: ctx}
	deleteItemUC := usecases.NewDeleteItem(deleteRepo)
	err = deleteItemUC.ExecuteDeleteItem(id)
	if err != nil {
		jsonResp := "{\"Error\": \"" + err.Error() + "\"}"
		return errors.New(jsonResp)
	}

	return err
}

func (t *SimpleQueue) Search(ctx contractapi.TransactionContextInterface, search string) (string, error) {
	fmt.Println("SimpleQueue > Search items")
	var err error
	var items []*domain.Item
	var filter contracts.Filter

	filter.Limit = 100
	filter.Sort = "[{\"order\": \"desc\"}]"
	filter.Search = search

	itemsRepo := &repo.ItemRepo{Ctx: ctx}
	searchItemsUC := usecases.NewSearchItems(itemsRepo)

	items, err = searchItemsUC.ExecuteSearchItems(&filter)
	if err != nil {
		jsonResp := "{\"Error\": \"" + err.Error() + "\"}"
		return "", errors.New(jsonResp)
	}

	itemJson, err := json.Marshal(items)
	if err != nil {
		fmt.Println("SimpleQueue.Search (1) " + err.Error())
		jsonResp := "{\"Error\":\"Items search failed\"}"
		return "", errors.New(jsonResp)
	}

	return string(itemJson), nil
}

func (t *SimpleQueue) Test(ctx contractapi.TransactionContextInterface) (string, error) {
	fmt.Println("SimpleQueue > Test")
	var err error

	return "", err
}