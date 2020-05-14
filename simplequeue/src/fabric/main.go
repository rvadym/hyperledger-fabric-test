package main

import (
	"fmt"
	"github.com/rvadym/hyperledger-fabric-test/adapters/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	cc, err := contractapi.NewChaincode(new(chaincode.SimpleQueue))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting SimpleQueue chaincode: %s", err)
	}
}
