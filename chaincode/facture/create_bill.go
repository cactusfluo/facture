package	main

import	"fmt"
import	"encoding/json"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	getBill(ownerId string, items []Item) Bill {
	var	bill		Bill
	var	totalAmount	uint64
	var	item		Item

	totalAmount = 0
	for _, item = range items {
		totalAmount += item.Amount * item.Count
	}
	bill.OwnerId = ownerId
	bill.Items = items
	bill.TotalAmount = totalAmount
	return bill
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	createBill(args []string) (string, error) {
	var	err			error
	var	bill		Bill
	var	billBytes	[]byte
	var	billId		string
	var	ownerId		string
	var	items		[]Item

	/// CHECK ARGUMENTS
	if len(args) != 1 {
		return "", fmt.Errorf("createBill requires one argument. An item list")
	}

	/// GET ARGUMENTS
	err = json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		return "", fmt.Errorf("Cannot unmarshal item list: %s", err)
	}

	/// COMPUTE BILL
	ownerId, _ = getPublicKey()
	billId = STUB.GetTxID()
	bill = getBill(ownerId, items)
	billBytes, err = json.Marshal(bill)
	if err != nil {
		return "", fmt.Errorf("Cannot marshal resulting bill: %s", err)
	}

	println("Bill ID:", billId)
	println("Owner ID:", ownerId)
	println("Bill Items: ", args[0])
	/// PUT STATE
	err = STUB.PutState(billId, []byte(billBytes))
	if err != nil {
		return "", fmt.Errorf("Failed to set bill into ledger: %s", err)
	}

	return string(billBytes), nil
}
