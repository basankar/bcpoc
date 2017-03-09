package devices

import (
	"fmt"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

var logger = shim.NewLogger("DIChaincode")

type IMEI_Holder struct {
	IMEIs 	[]string `json:"imeis"`
}

type Device struct {
	DeviceName     string `json:"devicename"`
	DeviceModel    string `json:"devicemodel"`
	DateOfManf     string `json:"dateofmanf"`
	DateOfSale     string `json:"dateofsale"`
	OldIMEI        string `json:"oldimei"`
	IMEI	       string `json:"imei"`
	Status         string `json:"status"`
	SoldBy         string `json:"soldby"`
	Owner          string `json:"owner"`
}

type SimpleChainCode struct {
}

func (t *SimpleChainCode) Init(stub shim.ChaincodeStubInterface, function string, args[] string) ([]byte, error ) {
	
	var imeiIds IMEI_Holder
	
	bytes, err := json.Marshal(imeiIds);
	
	err = stub.PutState("imeiIds", bytes)
	
	return nil, nil
	 
} 

func (t *SimpleChainCode) Invoke(stub shim.ChaincodeStubInterface, function string, args[] string) ([]byte, error) {
	
	if function == "create_device" {
		return	t.createDevice(stub, args[0])
	}	
	return nil, nil
}

func (t *SimpleChainCode) Query(stub shim.ChaincodeStubInterface, function string, args[] string) ([]byte, error) {
	var dev  Device
	if function == "get_device_details" {
		dev, err = t.get_device(stub, args[0])
		if err != nil { fmt.Printf("error retrieving device details"); return nil, errors.New("error retrieving device details")}
		return t.get_dev_details(stub, &dev)
	}
	return nil, nil
}

func (t *SimpleChainCode) createDevice(stub shim.ChaincodeStubInterface, imeiId string) ([]byte, error) {
	
	var dev Device
	
	DeviceName  := "\"deviceName\":\"LENOVO\", "
	DeviceModel := "\"devicemodel\":\"VIBE\", "
	DateOfManf  := "\"dateofmanf\":\"''03-12-2016''\" , "
	DateOfSale  := "\"dateofsale\":\"UNDEFINED\", "
	OldIMEI     := "\"oldimei\":\"UNDEFINED\", "
	IMEI_ID     := "\"imei\":\""+imeiId+"\", "
	Status     	:= "\"status\":\"CREATED\", "
	SoldBy     	:= "\"soldby\":\"UNDEFINED\", "
	Owner     	:= "\"owner\":\"MANF\" "
	
	json_device = " {" +deviceName+DeviceModel+DateOfManf+DateOfSale+OldIMEI+IMEI_ID+Status+SoldBy+Owner+"} "
	
	
	if IMEI_ID = null {
		fmt.Printf("Invalid device ID")
	}
	
	err = json.Unmarshal([]byte(json_device), &dev)
	
	record, err = stub.GetState(dev.IMEI)
	
	if record != nil { return nil, errors.New("Device already exists") }
	
	_, err = t.save_changes(stub, &dev)
	
	if err != nil { fmt.Printf("CREATEDEVICE: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	bytes, err := stub.GetState("imeiIds")

	if err != nil { return nil, errors.New("Unable to get imeiIds") }

	var IMEI_Ids IMEI_Holder

	err = json.Unmarshal(bytes, &IMEI_Ids)

	if err != nil {	return nil, errors.New("Corrupt IMEI_Holder record") }

	IMEI_Ids.IMEIs = append(IMEI_Ids.IMEIs, imeiId)

	bytes, err = json.Marshal(IMEI_Ids)

	if err != nil { fmt.Printf("Error creating IMEI_Holder record") }

	err = stub.PutState("imeiIds", bytes)

	if err != nil { return nil, errors.New("Unable to put the state") }

	return nil, nil

}

func (t *SimpleChainCode) save_changes(stub shim.ChaincodeStubInterface, d Device) (bool, error) {

	bytes, err := json.Marshal(d)

	if err != nil { fmt.Printf("SAVE_CHANGES: Error converting Device record: %s", err); return false, errors.New("Error converting Device record") }

	err = stub.PutState(d.IMEI, bytes)

	if err != nil { fmt.Printf("SAVE_CHANGES: Error storing device record: %s", err); return false, errors.New("Error storing device record") }

	return true, nil
}

func (t *SimpleChainCode) get_device(stub shim.ChaincodeStubInterface, imeiId string) (Device, error) {
	  dev, err = stub.GetState(imeiId)
	  if err != nil { fmt.Printf("error while retrieving device"); return dev, errors.New("error retrieving device") }
	  return dev, nil
}

func (t *SimpleChainCode) get_dev_details(stub shim.ChaincodeStubInterface, device Device) ([]byte, error){
	bytes, err := json.Marshal(device)
	
	if err != nil {fmt.Printf("error converting device record "); return bytes, errors.New("Error converting device record")}
	
	return bytes, nil
}

func main() {
	err := shim.Start(new(SimpleChainCode));
	
	if err != null { fmt.Println("error while starting shim code"); 
	} else {
		fmt.Println("chaincode started");
	}
}
