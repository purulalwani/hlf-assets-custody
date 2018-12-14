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


const prefixBank = "BNK"

// METHOD TO INITIATE RECORDS IN BANK MASTER
func InitBank(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("******************************")
    fmt.Println("---------- INIT BANK----------")
    fmt.Println("******************************")

    // INITIALIZE BANK MASTER WITH DATA
    _bankMaster := []entity.BankMaster {
                     entity.BankMaster {UserName: "hvsrikanth",BankAC: "HDFC00001", Balance:10000},
                     entity.BankMaster {UserName: "madhurib",  BankAC: "HDFC00002", Balance:20000},
                     entity.BankMaster {UserName: "naveens",   BankAC: "HDFC00003", Balance:30000},
                    }
    
    for i := 0; i < len(_bankMaster); i++ {

        fmt.Println("i is : ", i)

        // PREPARE THE KEY VALUE PAIR TO PERSIST THE INVESTOR
        _bankKey, err := stub.CreateCompositeKey(prefixBank, []string{_bankMaster[i].UserName})
        // CHECK FOR ERROR IN CREATING COMPOSITE KEY
        if err != nil {
            return shim.Error(err.Error())
        }

        // MARSHAL THE BANK MASTER RECORD 
        _bankMasterAsBytes, err := json.Marshal(_bankMaster[i])
        // CHECK FOR ERROR IN MARSHALING
        if err != nil {
            return shim.Error(err.Error())
        }

        // NOW WRITE THE BANK MASTER RECORD
        err = stub.PutState(_bankKey, _bankMasterAsBytes)
        // CHECK FOR ERROR
        if err != nil {
            return shim.Error(err.Error())
        }
        
    }

    fmt.Println("***********************************")
    fmt.Println("---------- OUT INIT BANK ----------")
    fmt.Println("***********************************")

    // RETURN SUCCESS
    return shim.Success(nil)
}

