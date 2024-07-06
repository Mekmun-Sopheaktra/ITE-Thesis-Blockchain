#!/bin/bash

# Usage: ./register_vehicle.sh V006 Toyota Camry C123 123456 ABC789 Red 2023 Sedan Reaksa "123 Main St" ABC123 "2024-06-23" "2024-06-23" ABC123456
# Check if the correct number of arguments are provided
if [ "$#" -ne 15 ]; then
    echo "Usage: $0 id brand model modelCode bodyNumber engineNumber color madeYear vehicleType ownerName ownerAddress plateNumber firstRegisterDate lastTransferDate vin"
    exit 1
fi

# Assign input arguments to variables
ID=$1
Brand=$2
Model=$3
ModelCode=$4
BodyNumber=$5
EngineNumber=$6
Color=$7
MadeYear=$8
VehicleType=$9
OwnerName=${10}
OwnerAddress=${11}
PlateNumber=${12}
FirstRegisterDate=${13}
LastTransferDate=${14}
VIN=${15}

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
  --peerAddresses localhost:11051 \
  --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG3 \
  -c "{\"function\":\"RegisterVehicle\",\"Args\":[
  \"$ID\",
  \"$Brand\",
  \"$Model\",
  \"$ModelCode\",
  \"$BodyNumber\",
  \"$EngineNumber\",
  \"$Color\",
  \"$MadeYear\",
  \"$VehicleType\",
  \"$OwnerName\",
  \"$OwnerAddress\",
  \"$PlateNumber\",
  \"$FirstRegisterDate\",
  \"$LastTransferDate\",
  \"$VIN\"]}"