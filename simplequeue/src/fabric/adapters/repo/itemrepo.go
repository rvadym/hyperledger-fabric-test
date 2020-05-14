package repo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/rvadym/hyperledger-fabric-test/domain"
	"strconv"
	"strings"
)

type dbItem struct {
	Key string `json:"Key"`
	Record domain.Item `json:"Record"`
}

var orderIndexName = "order~itemId"

type ItemRepo struct {
	Ctx contractapi.TransactionContextInterface
}

func (r *ItemRepo) SaveItems(items []*domain.Item) error {
	var err error

	for _, item := range items {
		err = r.SaveItem(item)
		if err != nil {
			fmt.Println("Error: ItemRepo.SaveItems (1) " + err.Error())
			return err
		}
	}

	return nil
}

func (r *ItemRepo) SaveItem(item *domain.Item) error {
	var err error
	var stub = r.Ctx.GetStub()

	fmt.Printf(
		"Info: SaveItem :: Item ID = %s, Order = %d, Content = %s \n",
		item.ID,
		item.Order,
		item.Content,
	)

	itemJson, err := json.Marshal(item)
	if err != nil {
		fmt.Println("Error: ItemRepo.SaveItem (1) " + err.Error())
		return err
	}

	err = stub.PutState(item.ID, itemJson) // []byte(itemJson)
	if err != nil {
		fmt.Println("Error: ItemRepo.SaveItem (2) " + err.Error())
		return err
	}

	orderItemIdIndexKey, err := r.generateOrderItemIdCompositeKeyName(item.Order, item.ID)
	if err != nil {
		fmt.Println("Error: ItemRepo.SaveItem (3) " + err.Error())
		return err
	}

	err = r.deleteAllOrderItemIdCompositeKeys(item.ID)
	if err != nil {
		fmt.Println("Error: ItemRepo.SaveItem (4) " + err.Error())
		return err
	}

	err = stub.PutState(orderItemIdIndexKey, []byte{0x00})
	if err != nil {
		fmt.Println("Error: ItemRepo.SaveItem (5) " + err.Error())
		return err
	}

	return nil
}

func (r *ItemRepo) GetItem(id string) (*domain.Item, error) {
	var err error
	var item *domain.Item
	var avalbytes []byte
	var stub = r.Ctx.GetStub()

	fmt.Printf(
		"GetItem :: Item ID = %s \n",
		id,
	)

	avalbytes, err = stub.GetState(id)
	if err != nil {
		fmt.Println("Error: ItemRepo.GetItem (1) " + err.Error())
		return nil, err
	}
	if avalbytes == nil {
		fmt.Println("Error: ItemRepo.GetItem (2) Nothing found")
		return nil, errors.New("nothing found")
	}

	err = json.Unmarshal(avalbytes, &item)
	if err != nil {
		fmt.Println("Error: ItemRepo.GetItem (3) " + err.Error())
		return nil, err
	}

	return item, err
}

func (r *ItemRepo) GetFirstItem() (*domain.Item, error) {
	var err error
	var orderItemIdIndexKeys []string
	var item *domain.Item
	var firstOrderItemIdIndexKey string

	orderItemIdIndexKeys, err = r.getAllOrderItemIdCompositeKeys()
	if err != nil {
		fmt.Println("Error: ItemRepo.GetFirstItem (1) " + err.Error())
		return nil, err
	}
	if len(orderItemIdIndexKeys) == 0 {
		fmt.Println("Error: ItemRepo.GetFirstItem (2) Nothing found")
		return nil, errors.New("nothing found")
	}

	firstOrderItemIdIndexKey = orderItemIdIndexKeys[0]

	_, itemId, err := r.splitItemCompositeKey(firstOrderItemIdIndexKey)
	if err != nil {
		fmt.Println("Error: ItemRepo.GetFirstItem (3) " + err.Error())
		return nil, err
	}

	item, err = r.GetItem(itemId)
	if err != nil {
		fmt.Println("Error: ItemRepo.GetFirstItem (4) " + err.Error())
		return nil, err
	}

	return item, err
}

func (r *ItemRepo) GetLastItem() (*domain.Item, error) {
	var err error
	var orderItemIdIndexKeys []string
	var item *domain.Item
	var lastOrderItemIdIndexKey string

	orderItemIdIndexKeys, err = r.getAllOrderItemIdCompositeKeys()
	if err != nil {
		fmt.Println("Error: ItemRepo.GetLastItem (1) " + err.Error())
		return nil, err
	}
	if len(orderItemIdIndexKeys) == 0 {
		fmt.Println("Error: ItemRepo.GetLastItem (2) Nothing found")
		return nil, errors.New("nothing found")
	}

	lastOrderItemIdIndexKey = orderItemIdIndexKeys[len(orderItemIdIndexKeys)-1]

	_, itemId, err := r.splitItemCompositeKey(lastOrderItemIdIndexKey)
	if err != nil {
		fmt.Println("Error: ItemRepo.GetLastItem (3) " + err.Error())
		return nil, err
	}

	fmt.Println("Info: ItemRepo.GetLastItem | itemId" + itemId)

	item, err = r.GetItem(itemId)
	if err != nil {
		fmt.Println("Error: ItemRepo.GetLastItem (4) " + err.Error())
		return nil, err
	}

	return item, err
}

func (r *ItemRepo) DeleteItem(id string) error {
	var err error
	var item *domain.Item
	var stub = r.Ctx.GetStub()

	item, err = r.GetItem(id)
	if err != nil {
		fmt.Println("Error: ItemRepo.DeleteItem (1) " + err.Error())
		return err
	}

	err = stub.DelState(item.ID)
	if err != nil {
		fmt.Println("Error: ItemRepo.DeleteItem (2) " + err.Error())
		return err
	}

	err = r.deleteOrderItemIdCompositeKey(item.Order, item.ID)
	if err != nil {
		fmt.Println("Error: ItemRepo.DeleteItem (3) " + err.Error())
		return err
	}

	return nil
}

func (r *ItemRepo) GetAllItems() ([]*domain.Item, error) {
	var err error
	var orderItemIdIndexKeys []string
	var items []*domain.Item

	orderItemIdIndexKeys, err = r.getAllOrderItemIdCompositeKeys()
	if err != nil {
		fmt.Println("Error: ItemRepo.GetAllItems (1) " + err.Error())
		return nil, err
	}

	for _, orderItemIdIndexKey := range orderItemIdIndexKeys {
		_, itemId, err := r.splitItemCompositeKey(orderItemIdIndexKey)
		if err != nil {
			fmt.Println("Error: ItemRepo.GetAllItems (2) " + err.Error())
			return nil, err
		}

		item, err := r.GetItem(itemId)
		if err != nil {
			fmt.Println("Error: ItemRepo.GetAllItems (3) " + err.Error())
			return nil, err
		}

		items = append(items, item)
	}


	return items, err
}

func (r *ItemRepo) SearchItems(search string, perPage int32, sort string) ([]*domain.Item, error) {
	var err error
	var queryResults []byte
	var items []*domain.Item
	var dbItems []dbItem

	sortString := ""
	searchString := ""
	queryString := ""
	sortTemplate := ",\"sort\":%s"
	searchTemplate := ",\"content\":{\"$regex\":\"%s\""
	queryTemplate := "{\"selector\":{\"docType\":\"item\"%s}}%s}"

	if sort != "" {
		sortString = fmt.Sprintf(sortTemplate, sort)
	}

	if search != "" {
		searchString = fmt.Sprintf(searchTemplate, search)
	}

	queryString = fmt.Sprintf(queryTemplate, searchString, sortString)

	queryResults, err = r.getQueryResultForQueryString(queryString, perPage, "")
	if err != nil {
		fmt.Println("Error: ItemRepo.SearchItems (2) " + err.Error())
		return nil, err
	}

	err = json.Unmarshal(queryResults, &dbItems)
	if err != nil {
		fmt.Println("Error: ItemRepo.SearchItems (3) " + err.Error())
		return nil, err
	}

	for _, dbItem := range dbItems {
		items = append(items, &dbItem.Record)
	}

	return items, nil
}

func (r *ItemRepo) getQueryResultForQueryString(queryString string, perPage int32, bookmark string) ([]byte, error) {
	fmt.Printf("Info: getQueryResultForQueryString queryString:\n%s\n", queryString)

	var err error
	var resultsIterator shim.StateQueryIteratorInterface
	var stub = r.Ctx.GetStub()

	if perPage > 0 {
		resultsIterator, _, err = stub.GetQueryResultWithPagination(queryString, perPage, bookmark)
	} else {
		resultsIterator, err = stub.GetQueryResult(queryString)
	}

	if err != nil {
		return nil, err
	}

	defer resultsIterator.Close()

	buffer, err := r.constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Info: getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (r *ItemRepo) constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func (r *ItemRepo) deleteAllOrderItemIdCompositeKeys(itemId string) error {
	var err error
	var stub = r.Ctx.GetStub()

	orderItemIdIndexKeys, err := r.getOrderItemIdCompositeKeysForId(itemId)
	if err != nil {
		fmt.Println("Error: ItemRepo.deleteAllOrderItemIdCompositeKeys (1) " + err.Error())
		return err
	}

	for _, orderItemIdIndexKey := range orderItemIdIndexKeys {
		fmt.Println("Info: I'm going to delete key: " + orderItemIdIndexKey)
		err = stub.DelState(orderItemIdIndexKey)
		if err != nil {
			fmt.Println("Error: ItemRepo.deleteAllOrderItemIdCompositeKeys (6) " + err.Error())
			fmt.Println("     orderItemIdIndexKey: " + orderItemIdIndexKey)
			fmt.Println("                  itemId: " + itemId)
			return err
		}
	}

	return nil
}


func (r *ItemRepo) getAllOrderItemIdCompositeKeys() ([]string, error) {
	return r.getOrderItemIdCompositeKeysForId("")
}

func (r *ItemRepo) getOrderItemIdCompositeKeysForId(id string) ([]string, error) {
	var err error
	var orderItemIdIndexKeys = make([]string, 0)
	var stub = r.Ctx.GetStub()
	var resultsIterator shim.StateQueryIteratorInterface

	resultsIterator, err = stub.GetStateByPartialCompositeKey(orderIndexName, []string{})
	if err != nil {
		fmt.Println("Error: ItemRepo.getAllOrderItemIdCompositeKeys (1) " + err.Error())
		return nil, err
	}
	defer resultsIterator.Close()

	var i int
	for i = 0; resultsIterator.HasNext(); i++ {
		responseRange, err := resultsIterator.Next()
		if err != nil {
			fmt.Println("Error: ItemRepo.getAllOrderItemIdCompositeKeys (2) " + err.Error())
			return nil, err
		}

		itemOrder, itemId, err := r.splitItemCompositeKey(responseRange.Key)
		if err != nil {
			fmt.Println("Error: ItemRepo.getAllOrderItemIdCompositeKeys (3) " + err.Error())
			return nil, err
		}

		orderItemIdIndexKey, err := r.generateOrderItemIdCompositeKeyName(
			itemOrder,
			itemId,
		)
		if err != nil {
			fmt.Println("Error: ItemRepo.getAllOrderItemIdCompositeKeys (5) " + err.Error())
			fmt.Println("     orderItemIdIndexKey: " + orderItemIdIndexKey)
			fmt.Println("               itemOrder: " + strconv.Itoa(itemOrder))
			fmt.Println("                  itemId: " + itemId)
			return nil, err
		}

		//fmt.Println("--------> orderItemIdIndexKey: " + orderItemIdIndexKey)
		//fmt.Println("--------> itemOrder:           " + compositeKeyParts[0])
		//fmt.Println("--------> itemId:              " + compositeKeyParts[1])

		if id != "" {
			if id == itemId {
				orderItemIdIndexKeys = append(orderItemIdIndexKeys, orderItemIdIndexKey)
			}
		} else {
			orderItemIdIndexKeys = append(orderItemIdIndexKeys, orderItemIdIndexKey)
		}
	}

	return orderItemIdIndexKeys, err
}

func (r *ItemRepo) deleteOrderItemIdCompositeKey(order int, itemId string) error {
	var err error
	var stub = r.Ctx.GetStub()

	orderItemIdIndexKey, err := r.generateOrderItemIdCompositeKeyName(order, itemId)
	if err != nil {
		fmt.Println("Error: ItemRepo.deleteOrderItemIdCompositeKey (1) " + err.Error())
		return err
	}

	err = stub.DelState(orderItemIdIndexKey)
	if err != nil {
		fmt.Println("Error: ItemRepo.deleteOrderItemIdCompositeKey (2) " + err.Error())
		return err
	}

	return nil
}

func (r *ItemRepo) generateOrderItemIdCompositeKeyName(order int, itemId string) (string, error) {
	var err error
	var stub = r.Ctx.GetStub()

	orderItemIdIndexKey, err := stub.CreateCompositeKey(
		orderIndexName,
		[]string{
			r.generateFixedLengthOrderString(order),
			"ID" + itemId,
		},
	)
	if err != nil {
		fmt.Println("Error: ItemRepo.generateOrderItemIdCompositeKeyName (1) " + err.Error())
		return "", err
	}

	return orderItemIdIndexKey, err
}

func (r *ItemRepo) generateFixedLengthOrderString(value int) string {
	requiredLength := 10

	valueString := strconv.Itoa(value)
	actualLentgh := len(valueString)
	missingLength := requiredLength - actualLentgh

	for i := 0; i < missingLength; i++ {
		valueString = "0" + valueString
	}

	//fmt.Println(strconv.Itoa(missingLength))
	//fmt.Println("generateFixedLengthOrderString: " + valueString)

	return "ORD" + valueString
}

func (r *ItemRepo) splitItemCompositeKey(compositeKey string) (int, string, error) {
	var err error
	var stub = r.Ctx.GetStub()

	_, compositeKeyParts, err := stub.SplitCompositeKey(compositeKey)
	if err != nil {
		fmt.Println("Error: ItemRepo.splitItemCompositeKey (1) " + err.Error())
		return 0, "", err
	}

	orderPart := compositeKeyParts[0]
	itemIdPart := compositeKeyParts[1]

	//fmt.Println("splitItemCompositeKey orderPart   " + orderPart)
	//fmt.Println("splitItemCompositeKey itemIdPart  " + itemIdPart)

	orderString := strings.Replace(orderPart, "ORD", "", 1)
	//fmt.Println("splitItemCompositeKey orderString " + orderString)

	itemId := strings.Replace(itemIdPart, "ID", "", 1)
	//fmt.Println("splitItemCompositeKey itemId      " + itemId)

	order, err := strconv.Atoi(orderString)
	if err != nil {
		fmt.Println("Error: ItemRepo.splitItemCompositeKey (2) " + err.Error())
		return 0, "", err
	}

	//fmt.Println("splitItemCompositeKey order       " + strconv.Itoa(order))

	return order, itemId, nil
}