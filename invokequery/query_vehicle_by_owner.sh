#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 OwnerName"
    exit 1
fi

# Assign input argument to variable
OwnerName=$1

# Execute the chaincode query command
peer chaincode query -C vehiclechannel \
  -n vehicle \
  -c "{\"function\":\"QueryVehicleByOwner\",\"Args\":[\"$OwnerName\"]}"
