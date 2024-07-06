# Setup chaincode in chaincode dir

>       go mod init example.com/vehicle
>        go get
>        go build

# cd to test-network dir

# Start Test Network:

>       ./network.sh up 
>	    ./network.sh createChannel -c vehiclechannel

# Set Network ENV VARS

>       export PATH=${PWD}/../bin:$PATH
>	      export FABRIC_CFG_PATH=$PWD/../config/
>       export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem


# Package Chaincode:

>    peer lifecycle chaincode package vehicle.tar.gz --path ../chaincode/ --lang golang --label vehicle_1

# Install on Org1 peer:

    SETENV:

>       export CORE_PEER_TLS_ENABLED=true
>		export CORE_PEER_LOCALMSPID="Org1MSP"
>		export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
>		export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
>		export CORE_PEER_ADDRESS=localhost:7051
    
    Install:

>		peer lifecycle chaincode install vehicle.tar.gz --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# Approve on Org1 peer:
BE SURE TO UPDATE vehicle ID IN COMMAND BELOW!
    Approve:

>        peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --sequence 1 --cafile $ORDERER_CA --channelID vehiclechannel --name vehicle --version 1.0 --init-required --package-id <ID>



# Install on Org2 peer:

    SETENV:

>		export CORE_PEER_LOCALMSPID="Org2MSP"
>		export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
>		export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
>		export CORE_PEER_ADDRESS=localhost:9051
    
    Install:

>        peer lifecycle chaincode install vehicle.tar.gz --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# Approve on Org2 peer:
BE SURE TO UPDATE vehicle ID IN COMMAND BELOW!
    Approve:

>        peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --sequence 1 --cafile $ORDERER_CA --channelID vehiclechannel --name vehicle --version 1.0 --init-required --package-id <ID> 


# Install on Org3 peer:
    SETENV:

>		export CORE_PEER_LOCALMSPID="Org3MSP"
>		export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
>		export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
>		export CORE_PEER_ADDRESS=localhost:11051
    
    Install:

>        peer lifecycle chaincode install vehicle.tar.gz --peerAddresses localhost:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE

# Approve on Org3 peer
BE SURE TO UPDATE vehicle ID IN COMMAND BELOW! pApprove:

>        peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --sequence 1 --cafile $ORDERER_CA --channelID vehiclechannel --name vehicle --version 1.0 --init-required --package-id <ID>

# Commit Chaincode:
    SET ENV:

>		export CORE_PEER_LOCALMSPID="Org1MSP"
>    	export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
>		export CORE_PEER_TLS_ROOTCERT_FILE_ORG1=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
>		export CORE_PEER_TLS_ROOTCERT_FILE_ORG2=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
>		export CORE_PEER_TLS_ROOTCERT_FILE_ORG3=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
>    	export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
>    	export CORE_PEER_ADDRESS=localhost:7051
>		export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

    Commit Chaincode    
>        peer lifecycle chaincode commit -o  localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID vehiclechannel --name vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 --peerAddresses localhost:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG3 --version 1.0 --sequence 1 --init-required

# Invoke Chaincode:
    	
>        peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C vehiclechannel -n vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 --isInit -c '{"Args":[]}'

# QUERY COMMITED CHAINCODE
		//Add vehicle
		
>       peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C vehiclechannel -n vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 -c '{"function":"RegisterVehicle","Args":["Toyota","Corolla","XYZ123","Red","BODY1234","ENG1234","2021","Car","Sedan","2021-01-01T15:04:05Z"]}'

		//Query vehicle by ID
>		peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C vehiclechannel -n vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG2 -c '{"Args":["QueryVehicleByPlateNumber", "XYZ123"]}'

		//Query All Properties
>		peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C vehiclechannel -n vehicle --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE_ORG1 '{"Args":["QueryAllVehicle"]}'

# Shutdown Network:
>        ./network.sh down