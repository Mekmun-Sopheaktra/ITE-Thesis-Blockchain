#!/bin/bash

# Start the Test Network and create a channel
./network.sh up
./network.sh createChannel -c vehiclechannel

# Set Network ENV VARS
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

# Package Chaincode

peer lifecycle chaincode package vehicle.tar.gz --path ../chaincode/vehicle-app/go --lang golang --label vehicle_1
echo "Chaincode vehicle_1 package successfully"

# Install on Org1 peer
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode install vehicle.tar.gz --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
echo "Chaincode vehicle_1 install successfully in Org1"
# Approve on Org1 peer
PACKAGE_ID=$(peer lifecycle chaincode queryinstalled | grep vehicle_1 | awk '{print $3}' | sed 's/,//')

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --sequence 1 --cafile $ORDERER_CA --channelID vehiclechannel --name vehicle --version 1.0 --init-required --package-id $PACKAGE_ID
echo "Chaincode vehicle_1 approveformyorg successfully in Org1"

# Install on Org2 peer
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

peer lifecycle chaincode install vehicle.tar.gz --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
echo "Chaincode vehicle_1 install successfully in Org2"
# Approve on Org2 peer
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --sequence 1 --cafile $ORDERER_CA --channelID vehiclechannel --name vehicle --version 1.0 --init-required --package-id $PACKAGE_ID
echo "Chaincode vehicle_1 approveformyorg successfully in Org2"
# Install on Org3 peer
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
export CORE_PEER_ADDRESS=localhost:11051

peer lifecycle chaincode install vehicle.tar.gz --peerAddresses localhost:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE
echo "Chaincode vehicle_1 install successfully in Org3"
# Approve on Org3 peer
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --sequence 1 --cafile $ORDERER_CA --channelID vehiclechannel --name vehicle --version 1.0 --init-required --package-id $PACKAGE_ID
echo "Chaincode vehicle_1 approveformyorg successfully in Org3"
# Commit Chaincode
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE_ORG1=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE_ORG2=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE_ORG3=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID vehiclechannel --name vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 --peerAddresses localhost:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG3 --version 1.0 --sequence 1 --init-required
echo "Chaincode vehicle_1 commit successfully"
# Invoke Chaincode
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C vehiclechannel -n vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 --isInit -c '{"Args":[]}'
echo "Chaincode vehicle_1 invoke init successfully"
## Add a vehicle
#peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C vehiclechannel -n vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 -c '{"function":"RegisterVehicle","Args":["Toyota","Corolla","XYZ123","Red","BODY1234","ENG1234","2021","Car","Sedan","2021-01-01T15:04:05Z"]}'
#
## Query vehicle by ID
#peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C vehiclechannel -n vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 -c '{"Args":["QueryVehicleByPlateNumber", "XYZ123"]}'
#
## Query all vehicles
#peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C vehiclechannel -n vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 -c '{"Args":["QueryAllVehicle"]}'

# print success message
echo "Chaincode vehicle_1 deployed successfully on vehiclechannel"
