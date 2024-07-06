#!/bin/bash

# Execute the chaincode invoke command for QueryAllVehicles
peer chaincode query -C vehiclechannel -n vehicle -c '{"Args":["QueryAllVehicles"]}'
