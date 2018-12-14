package util

import (
    "fmt"
    "encoding/json"
    //"strings"
    //"time"
    "github.com/src/entity"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

const prefixCustodian = "CUSTDN"

// METHOD TO CREATE INVESTOR
func OnboardInvestor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Printf("***************************************\n")
    fmt.Printf("---------- IN ONBOARDINVESTOR----------\n")
    fmt.Printf("***************************************\n")

    // RETURN ERROR IF ARGS IS NOT 7 IN NUMBER
    if len(args) != 7 {
        fmt.Printf("**************************\n")
        fmt.Printf("Too few argments... Need 7\n")
        fmt.Printf("**************************\n")

        return shim.Error("Invalid argument count. Expecting 7.")
    }
/**
    // CREATE A TEMP STRUCTURE TO RECEIVE INVESOR DATA FROM API
    dto := struct {
        userName     string  `json:"user_name"`
        userFName    string  `json:"user_fname"`
        userLName    string  `json:"user_lname"`
        userIdentity string  `json:"user_identity"`
        kycStatus    string  `json:"kyc_status"`
        depositoryAC string  `json:"depository_ac"`
        bankAC       string  `json:"bank_ac"`
    } {}

    // CHECK FOR ERROR IN PARSING INPUT
    err := json.Unmarshal([]byte(args[0]), &dto)
    if err != nil {
        return shim.Error(err.Error())
    }
**/
    // PREPARE THE INPUT VALUES TO WRITE
    _investor := entity.Investor {
        UserName:     args[0], //dto.userName,
        UserFName:    args[1], //dto.userFName,
        UserLName:    args[2], //dto.userLName,
        UserIdentity: args[3], //dto.userIdentity,
        KycStatus:    args[4], //dto.kycStatus,
        DepositoryAC: args[5], //dto.depositoryAC,
        BankAC:       args[6], //dto.bankAC,
    }

    // PREPARE THE KEY VALUE PAIR TO PERSIST THE INVESTOR
    _investorKey, err := stub.CreateCompositeKey(prefixCustodian, []string{_investor.UserName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }

    // MARSHAL INVESTOR RECORD
    _investorBytes, err := json.Marshal(_investor)
    // CHECK FOR ERROR IN MARSHALING 
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE INVESTOR RECORD
    err = stub.PutState(_investorKey, _investorBytes)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }

    fmt.Printf("****************************************\n")
    fmt.Printf("---------- OUT ONBOARDINVESTOR----------\n")
    fmt.Printf("****************************************\n")

    // RETURN SUCCESS
    return shim.Success(nil)
}

