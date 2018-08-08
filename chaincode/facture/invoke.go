package main

import	"fmt"
import	"github.com/hyperledger/fabric/core/chaincode/shim"
import	"github.com/hyperledger/fabric/protos/peer"

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

// toChaincodeArgs converts string args to []byte args
//func toChaincodeArgs(args ...string) [][]byte {
//	bargs := make([][]byte, len(args))
//	for i, arg := range args {
//		bargs[i] = []byte(arg)
//	}
//	return bargs
//}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var fct string
	var argv []string
	var ret string
	var err error

	/// GET FUNCTION AND PARAMETERS
	fct, argv = stub.GetFunctionAndParameters()

	/// INIT ENVIRONNEMENT
	STUB = stub
	LOG = shim.NewLogger("Pcoin")
	LOG.SetLevel(shim.LogInfo)

	/// LAUNCH FUNCTION
	switch fct {
	case "createBill":
		ret, err = createBill(argv)
	case "payBill":
		ret, err = _get(argv)
	case "listBills":
		ret, err = _get(argv)
		//chainCodeArgs := toChaincodeArgs("get", "a")
		//response := stub.InvokeChaincode("sacc", chainCodeArgs, "myc")
		//if response.Status != shim.OK {
		//	return shim.Error(response.Message)
		//}
		//ret, err = string(response.Payload), nil
	default:
		err = fmt.Errorf("Illegal function called \n")
	}

	/// CHECK ERROR
	if err != nil {
		LOG.Error(err)
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(ret))
}
