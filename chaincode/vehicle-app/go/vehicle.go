package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// VehicleTransferSmartContract implements a smart contract to manage vehicles on a blockchain
type VehicleTransferSmartContract struct {
	contractapi.Contract
}

// Vehicle represents a vehicle asset
type Vehicle struct {
	ID                 string   `json:"id"`
	Brand              string   `json:"brand"`
	Model              string   `json:"model"`
	ModelCode          string   `json:"modelcode"`
	BodyNumber         string   `json:"bodynumber"`
	EngineNumber       string   `json:"enginenumber"`
	Color              string   `json:"color"`
	MadeYear           string   `json:"madeyear"`
	Type               string   `json:"type"`
	VehicleType        string   `json:"vehicletype"`
	OwnerName          string   `json:"ownerName"`
	OwnerAddress       string   `json:"ownerAddress"`
	PlateNumber        string   `json:"platenumber"`
	FirstRegisterDate  string   `json:"firstregisterdate"`
	LastTransferDate   string   `json:"lasttransferdate"`
	IsActive           bool     `json:"isActive"`
	PreviousOwners     []string `json:"previousOwners"`
	VIN                string   `json:"vin"`
	RegistrationStatus string   `json:"registrationStatus"`
	VerificationStatus string   `json:"verificationStatus"`
}

// RegisterVehicle registers a new vehicle in the ledger
func (vc *VehicleTransferSmartContract) RegisterVehicle(ctx contractapi.TransactionContextInterface,
	id, brand, model, modelCode, bodyNumber, engineNumber, color, madeYear, vehicleType, ownerName, ownerAddress, plateNumber, firstRegisterDate,
	lastTransferDate, vin string) error {

	vehicleJSON, err := ctx.GetStub().GetState(id) // checks if id already exists
	if err != nil {
		return fmt.Errorf("Failed to read the data from world state: %v", err)
	}
	if vehicleJSON != nil {
		return fmt.Errorf("The vehicle %s already exists", id)
	}

	vehicle := Vehicle{
		ID:                 id,
		Brand:              brand,
		Model:              model,
		ModelCode:          modelCode,
		BodyNumber:         bodyNumber,
		EngineNumber:       engineNumber,
		Color:              color,
		MadeYear:           madeYear,
		Type:               vehicleType,
		VehicleType:        vehicleType,
		OwnerName:          ownerName,
		OwnerAddress:       ownerAddress,
		PlateNumber:        plateNumber,
		FirstRegisterDate:  firstRegisterDate,
		LastTransferDate:   lastTransferDate,
		IsActive:           true,
		PreviousOwners:     []string{},
		VIN:                vin,
		RegistrationStatus: "registered",
		VerificationStatus: "unverified",
	}

	vehicleBytes, err := json.Marshal(vehicle) // create JSON encoding
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, vehicleBytes) // pass JSON to API
}

// VerifyVehicle updates the verification status of a vehicle
func (vc *VehicleTransferSmartContract) VerifyVehicle(ctx contractapi.TransactionContextInterface, id string) error {
	vehicle, err := vc.QueryVehicleByID(ctx, id)
	if err != nil {
		return err
	}

	// Perform verification checks (e.g., VIN check, ownership validation, etc.)

	// Update verification status
	vehicle.VerificationStatus = "verified"

	vehicleJSON, err := json.Marshal(vehicle)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, vehicleJSON)
}

// TransferVehicle transfers ownership of a vehicle
func (vc *VehicleTransferSmartContract) TransferVehicle(ctx contractapi.TransactionContextInterface,
	id, newOwner string) error {

	vehicle, err := vc.QueryVehicleByID(ctx, id)
	if err != nil {
		return err
	}

	// Check if vehicle is active
	if !vehicle.IsActive {
		return fmt.Errorf("vehicle %s is terminated and cannot be transferred", id)
	}

	// Add current owner to previous owners list
	if vehicle.OwnerName != "" {
		vehicle.PreviousOwners = append(vehicle.PreviousOwners, vehicle.OwnerName)
	}

	// Update owner and transfer date
	vehicle.OwnerName = newOwner
	vehicle.LastTransferDate = getCurrentDateTime()

	vehicleJSON, err := json.Marshal(vehicle)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, vehicleJSON)
}

// TerminateVehicle sets the IsActive flag of a vehicle to false
func (vc *VehicleTransferSmartContract) TerminateVehicle(ctx contractapi.TransactionContextInterface, id string) error {
	vehicle, err := vc.QueryVehicleByID(ctx, id)
	if err != nil {
		return err
	}

	// Update vehicle status to inactive
	vehicle.IsActive = false

	vehicleJSON, err := json.Marshal(vehicle)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, vehicleJSON)
}

// QueryAllVehicles returns all vehicles in the ledger
func (vc *VehicleTransferSmartContract) QueryAllVehicles(ctx contractapi.TransactionContextInterface) ([]*Vehicle, error) {
	vehicleIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer vehicleIterator.Close()

	var vehicles []*Vehicle
	for vehicleIterator.HasNext() {
		vehicleResponse, err := vehicleIterator.Next()
		if err != nil {
			return nil, err
		}

		var vehicle Vehicle
		err = json.Unmarshal(vehicleResponse.Value, &vehicle)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, &vehicle)
	}

	return vehicles, nil
}

// QueryVehicleByID returns a vehicle by its ID from the ledger
func (vc *VehicleTransferSmartContract) QueryVehicleByID(ctx contractapi.TransactionContextInterface, id string) (*Vehicle, error) {
	vehicleJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the data from world state: %v", err)
	}
	if vehicleJSON == nil {
		return nil, fmt.Errorf("The vehicle %s does not exist", id)
	}

	var vehicle Vehicle
	err = json.Unmarshal(vehicleJSON, &vehicle)
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

// QueryVehicleByOwner returns vehicles owned by a specific owner
func (vc *VehicleTransferSmartContract) QueryVehicleByOwner(ctx contractapi.TransactionContextInterface, owner string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByOwner []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.OwnerName == owner {
			vehiclesByOwner = append(vehiclesByOwner, vehicle)
		}
	}

	return vehiclesByOwner, nil
}

// QueryVehicleHistory returns the history of a vehicle by its ID
func (vc *VehicleTransferSmartContract) QueryVehicleHistory(ctx contractapi.TransactionContextInterface, id string) ([]Vehicle, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var vehicleHistory []Vehicle
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var vehicle Vehicle
		err = json.Unmarshal(response.Value, &vehicle)
		if err != nil {
			return nil, err
		}

		vehicleHistory = append(vehicleHistory, vehicle)
	}

	return vehicleHistory, nil
}

// QueryVehicleByRegistrationStatus returns vehicles by their registration status
func (vc *VehicleTransferSmartContract) QueryVehicleByRegistrationStatus(ctx contractapi.TransactionContextInterface, status string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByStatus []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.RegistrationStatus == status {
			vehiclesByStatus = append(vehiclesByStatus, vehicle)
		}
	}

	return vehiclesByStatus, nil
}

// QueryVehicleByPlateNumber returns a vehicle by its plate number
func (vc *VehicleTransferSmartContract) QueryVehicleByPlateNumber(ctx contractapi.TransactionContextInterface, plateNumber string) (*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	for _, vehicle := range vehicles {
		if vehicle.PlateNumber == plateNumber {
			return vehicle, nil
		}
	}

	return nil, fmt.Errorf("The vehicle with plate number %s does not exist", plateNumber)
}

// QueryVehicleByVIN returns a vehicle by its VIN
func (vc *VehicleTransferSmartContract) QueryVehicleByVIN(ctx contractapi.TransactionContextInterface, vin string) (*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	for _, vehicle := range vehicles {
		if vehicle.VIN == vin {
			return vehicle, nil
		}
	}

	return nil, fmt.Errorf("The vehicle with VIN %s does not exist", vin)
}

// QueryVehicleByType returns vehicles by their type
func (vc *VehicleTransferSmartContract) QueryVehicleByType(ctx contractapi.TransactionContextInterface, vehicleType string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByType []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.VehicleType == vehicleType {
			vehiclesByType = append(vehiclesByType, vehicle)
		}
	}

	return vehiclesByType, nil
}

// QueryVehicleByBrand returns vehicles by their brand
func (vc *VehicleTransferSmartContract) QueryVehicleByBrand(ctx contractapi.TransactionContextInterface, brand string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByBrand []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.Brand == brand {
			vehiclesByBrand = append(vehiclesByBrand, vehicle)
		}
	}

	return vehiclesByBrand, nil
}

// QueryVehicleByModel returns vehicles by their model
func (vc *VehicleTransferSmartContract) QueryVehicleByModel(ctx contractapi.TransactionContextInterface, model string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByModel []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.Model == model {
			vehiclesByModel = append(vehiclesByModel, vehicle)
		}
	}

	return vehiclesByModel, nil
}

// QueryVehicleByColor returns vehicles by their color
func (vc *VehicleTransferSmartContract) QueryVehicleByColor(ctx contractapi.TransactionContextInterface, color string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByColor []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.Color == color {
			vehiclesByColor = append(vehiclesByColor, vehicle)
		}
	}

	return vehiclesByColor, nil
}

// QueryVehicleByMadeYear returns vehicles by their made year
func (vc *VehicleTransferSmartContract) QueryVehicleByMadeYear(ctx contractapi.TransactionContextInterface, madeYear string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByMadeYear []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.MadeYear == madeYear {
			vehiclesByMadeYear = append(vehiclesByMadeYear, vehicle)
		}
	}

	return vehiclesByMadeYear, nil
}

// QueryVehicleByTypeAndBrand returns vehicles by their type and brand
func (vc *VehicleTransferSmartContract) QueryVehicleByTypeAndBrand(ctx contractapi.TransactionContextInterface, vehicleType, brand string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByTypeAndBrand []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.VehicleType == vehicleType && vehicle.Brand == brand {
			vehiclesByTypeAndBrand = append(vehiclesByTypeAndBrand, vehicle)
		}
	}

	return vehiclesByTypeAndBrand, nil
}

// QueryVehicleByTypeAndModel returns vehicles by their type and model
func (vc *VehicleTransferSmartContract) QueryVehicleByTypeAndModel(ctx contractapi.TransactionContextInterface, vehicleType, model string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByTypeAndModel []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.VehicleType == vehicleType && vehicle.Model == model {
			vehiclesByTypeAndModel = append(vehiclesByTypeAndModel, vehicle)
		}
	}

	return vehiclesByTypeAndModel, nil
}

// QueryVehicleByTypeAndColor returns vehicles by their type and color
func (vc *VehicleTransferSmartContract) QueryVehicleByTypeAndColor(ctx contractapi.TransactionContextInterface, vehicleType, color string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByTypeAndColor []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.VehicleType == vehicleType && vehicle.Color == color {
			vehiclesByTypeAndColor = append(vehiclesByTypeAndColor, vehicle)
		}
	}

	return vehiclesByTypeAndColor, nil
}

// QueryVehicleByTypeAndMadeYear returns vehicles by their type and made year
func (vc *VehicleTransferSmartContract) QueryVehicleByTypeAndMadeYear(ctx contractapi.TransactionContextInterface, vehicleType, madeYear string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByTypeAndMadeYear []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.VehicleType == vehicleType && vehicle.MadeYear == madeYear {
			vehiclesByTypeAndMadeYear = append(vehiclesByTypeAndMadeYear, vehicle)
		}
	}

	return vehiclesByTypeAndMadeYear, nil
}

// QueryVehicleByBrandAndModel returns vehicles by their brand and model
func (vc *VehicleTransferSmartContract) QueryVehicleByBrandAndModel(ctx contractapi.TransactionContextInterface, brand, model string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByBrandAndModel []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.Brand == brand && vehicle.Model == model {
			vehiclesByBrandAndModel = append(vehiclesByBrandAndModel, vehicle)
		}
	}

	return vehiclesByBrandAndModel, nil
}

// QueryVehicleByBrandAndColor returns vehicles by their brand and color
func (vc *VehicleTransferSmartContract) QueryVehicleByBrandAndColor(ctx contractapi.TransactionContextInterface, brand, color string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByBrandAndColor []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.Brand == brand && vehicle.Color == color {
			vehiclesByBrandAndColor = append(vehiclesByBrandAndColor, vehicle)
		}
	}

	return vehiclesByBrandAndColor, nil
}

// QueryVehicleByBrandAndMadeYear returns vehicles by their brand and made year
func (vc *VehicleTransferSmartContract) QueryVehicleByBrandAndMadeYear(ctx contractapi.TransactionContextInterface, brand, madeYear string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByBrandAndMadeYear []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.Brand == brand && vehicle.MadeYear == madeYear {
			vehiclesByBrandAndMadeYear = append(vehiclesByBrandAndMadeYear, vehicle)
		}
	}

	return vehiclesByBrandAndMadeYear, nil
}

// QueryVehicleByModelAndColor returns vehicles by their model and color
func (vc *VehicleTransferSmartContract) QueryVehicleByModelAndColor(ctx contractapi.TransactionContextInterface, model, color string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByModelAndColor []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.Model == model && vehicle.Color == color {
			vehiclesByModelAndColor = append(vehiclesByModelAndColor, vehicle)
		}
	}

	return vehiclesByModelAndColor, nil
}

// QueryVehicleByModelAndMadeYear returns vehicles by their model and made year
func (vc *VehicleTransferSmartContract) QueryVehicleByModelAndMadeYear(ctx contractapi.TransactionContextInterface, model, madeYear string) ([]*Vehicle, error) {
	vehicles, err := vc.QueryAllVehicles(ctx)
	if err != nil {
		return nil, err
	}

	var vehiclesByModelAndMadeYear []*Vehicle
	for _, vehicle := range vehicles {
		if vehicle.Model == model && vehicle.MadeYear == madeYear {
			vehiclesByModelAndMadeYear = append(vehiclesByModelAndMadeYear, vehicle)
		}
	}

	return vehiclesByModelAndMadeYear, nil
}

// Utility function to get current date and time
func getCurrentDateTime() string {
	return time.Now().Format(time.RFC3339)
}

func main() {
	vehicleTransferSmartContract := new(VehicleTransferSmartContract)

	cc, err := contractapi.NewChaincode(vehicleTransferSmartContract)
	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}
