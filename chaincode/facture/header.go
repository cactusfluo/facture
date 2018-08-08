package	main

import	"github.com/hyperledger/fabric/core/chaincode/shim"

type	SimpleAsset	struct {
}

type	Events		struct {
		Amount		uint64
		Allowances	map[string]uint64
		From		string
		To			string
		Value		uint64
}

var		STUB shim.ChaincodeStubInterface
var		LOG *shim.ChaincodeLogger
