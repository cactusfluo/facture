package	main

import	"fmt"
import	"encoding/json"
import	"github.com/hyperledger/fabric/core/chaincode/shim"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	toChaincodeArgs(args ...string) [][]byte {
	var	bargs		[][]byte
	bargs = make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	payBill(args []string) (string, error) {
	var	err			error
	var	bill		Bill
	var	billBytes	[]byte
	var	billId		string
	var	ccArgs		[][]byte

	/// CHECK ARGUMENTS
	if len(args) != 1 {
		return "", fmt.Errorf("payBill requires one argument. A bill ID")
	}

	/// GET ARGUMENTS
	billId = args[0]

	/// GET BILL
	billBytes, err = STUB.GetState(billId)
	if err != nil {
		return "", fmt.Errorf("Cannot get bill: %s", err)
	}
	if billBytes == nil {
		return "", fmt.Errorf("Innexistant bill: %s", billBytes)
	}
	err = json.Unmarshal(billBytes, &bill)
	if err != nil {
		return "", fmt.Errorf("Cannot unmarshal bill: %s", err)
	}

	/// CALL CHAINCODE TO PAY BILL
	ccArgs = toChaincodeArgs("transfer", bill.OwnerId, string(bill.TotalAmount))
	response := STUB.InvokeChaincode("ptwist", ccArgs, "ptwist")
	if response.Status != shim.OK {
		return "", fmt.Errorf("Cannot transfer assets for the bill: %s", response.Message)
	}

	/// DELETE BILL
	err = STUB.DelState(string(billBytes))
	if err != nil {
		return "", fmt.Errorf("Cannot delete bill: %s", err)
	}

	return string(billBytes), nil
}
