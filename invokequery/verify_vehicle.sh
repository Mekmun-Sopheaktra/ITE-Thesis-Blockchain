#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 VehicleID"
    exit 1
fi

# Assign input argument to variable
VehicleID=$1

# Execute the chaincode invoke command
peer chaincode invoke -o localhost:7050 \
  --ordererTLSHostnameOverride orderer.example.com \
  --tls $CORE_PEER_TLS_ENABLED \
  --cafile $ORDERER_CA \
  -C vehiclechannel \
  -n vehicle \
  --peerAddresses localhost:7051 \
  --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 \
  --peerAddresses localhost:9051 \
  --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 \
  -c "{\"function\":\"VerifyVehicle\",\"Args\":[\"$VehicleID\"]}"
