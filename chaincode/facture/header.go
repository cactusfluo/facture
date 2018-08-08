package	main

import	"github.com/hyperledger/fabric/core/chaincode/shim"

type	SimpleAsset	struct {
}

type	Item		struct {
		Name		string
		Amount		uint64
		Count		uint64
}

type	Bill		struct {
		OwnerId		string
		Items		[]Item
		TotalAmount	uint64
}

var		STUB		shim.ChaincodeStubInterface
var		LOG			*shim.ChaincodeLogger
