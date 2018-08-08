#!/usr/bin/env bash

peer chaincode invoke \
	-n facture \
	-c '{"Args":["createBill", "facture", "[{\"Name\": \"test\", \"Amount\": 21, \"Count\": 2}]"], "abdul"}' \
	-C myc
