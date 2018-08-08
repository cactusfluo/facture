package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func gethistory(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetHistoryForKey(args[0])

	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}

	var history string
	history = "\n"

	for value.HasNext() {
		history = fmt.Sprintf("%s%s", history, fmt.Sprintln(value.Next()))
	}

	return string(history), nil
}

// toChaincodeArgs converts string args to []byte args
func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var fct string
	var argv []string
	var ret string
	var err error

	fct, argv = stub.GetFunctionAndParameters()
	if fct != "balanceOf" && fct != "whoOwesMe" && fct != "whoOweI" { // TEMP
		fmt.Println("---------------> Invoke <---------------")
	}

	STUB = stub
	LOG = shim.NewLogger("Pcoin")
	LOG.SetLevel(shim.LogInfo)

	switch fct {
	// Temporary, not in production
	case "get":
		ret, err = _get(argv)
	case "history":
		ret, err = gethistory(stub, argv)
	case "remoteinvoke":
		chainCodeArgs := toChaincodeArgs("get", "a")
		response := stub.InvokeChaincode("sacc", chainCodeArgs, "myc")
		if response.Status != shim.OK {
			return shim.Error(response.Message)
		}
		ret, err = string(response.Payload), nil
	default:
		err = fmt.Errorf("Illegal function called \n")
	}

	if err != nil {
		LOG.Error(err)
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(ret))
}
