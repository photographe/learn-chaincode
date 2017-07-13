package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	
	err := stub.PutState("keyArg0", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "writeToBlockchain" {
		return t.writeToBlockchain(stub, args)
	} else if function == "writeToBlockchainBranch" {
		return t.writeToBlockchainBranch(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) writeToBlockchain(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println ("Within Write function..")
	
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments passed, it expects 2 - First : name of the key, Second: value to set for the key")
	}
	
	key = args[0]
	value = args[1]
	err = stub.PutState(key, []byte(value))
	
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

func (t *SimpleChaincode) writeToBlockchainBranch(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println ("Within Write To branch blockchain function..")
	
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments passed, it expects 2 - First : name of the key, Second: value to set for the key")
	}
	
	key = args[0]
	value = args[1]
	err = stub.PutState(key, []byte(value))
	
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "dummy_query" {								//read a variable
		fmt.Println("hi there " + function)						//error
		return nil, nil;
	}
	
	if function == "readFromBlockchain" {
		return t.readFromBlockchain(stub, args)
	}
	// fmt.Println("query did not find func: " + function)		//error

	return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) readFromBlockchain(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments passed, it expects 1 - First : name of the key to query")
	}
	
	key = args[0]
	valAsBytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	return valAsBytes, nil
}