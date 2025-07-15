#!/bin/bash

echo "===================================="
echo "Hauska Go SDK - GoChain Examples"
echo "===================================="
echo ""

echo "Network Configuration:"
echo "- Network Name: GoChain"
echo "- RPC URL: http://localhost:8545"
echo "- Chain ID: 31337"
echo "- Currency: GO"
echo "- Test Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
echo ""

echo "Running Balance Example..."
echo "-------------------------"
cd balance
go run main.go
