package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	token "./contracts/token"

	"./models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"gopkg.in/mgo.v2/bson"
)

func get_list_coins(writer http.ResponseWriter, request *http.Request) {
	//TODO: fetch templates from mongo
	//TODO: send to writer
	log.Printf("Getting coins list")
	err, coins := models.GetCoins(bson.M{})
	if err != nil {
		log.Printf("Error in listing coins: %v", err)
		http.Error(writer, `{"status":"ERROR","message":"couldn't list templates'"}`, http.StatusInternalServerError)
	}
	log.Printf("Got list: %v", coins)
	response := make(map[string]interface{})
	response["data"] = coins
	response["status"] = "SUCCESS"
	writer.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(writer)
	enc.Encode(response)
	return
}

func post_coin(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Could not parse the body due to: %v", err)
		http.Error(writer, "Could not parse the body", http.StatusBadRequest)
		return
	}
	var req struct {
		Name            string `json:"name"`
		Symbol          string `json:"symbol"`
		Supply          uint64 `json:"supply"`
		ContractAddress string `bson:"contract_address"`
		PrivateKey      string `json:"private_key"`
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("could not unmarshall the body due to: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"status":"ERROR","message":"Could not process the request due to an error unmarshalling the body"}`))
		return
	}

	// Connect to Ethereum Network
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	// Setting privateKey
	privateKey, err := crypto.HexToECDSA(req.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	var supply uint64 = req.Supply
	name := req.Name
	symbol := req.Symbol

	address, tx, instance, err := token.DeployToken(auth, client, supply, name, symbol)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Printf("Failed to deploy the contract: %v", err)
	}

	fmt.Println()
	fmt.Println(address.Hex())
	fmt.Println(tx)

	if len(address.Bytes()) == 0 {
		log.Printf("Expected a valid deployment address. Received empty address byte array instead")
	}
	writer.WriteHeader(http.StatusOK)
	_ = instance

	err = models.CreateCoin(
		req.Name,
		req.Symbol,
		req.Supply,
		address.Hex(),
	)
	if err != nil {
		log.Printf("could not save the body due to: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"status":"ERROR","message":"Could not process the request due to an error saving the body"}`))
		return
	}

	writer.Write([]byte(`{"status":"SUCCESS","message":"Created new template"}`))
	return

}
