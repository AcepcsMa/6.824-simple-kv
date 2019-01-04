package main

import (
	"errors"
	"fmt"
	"net/rpc"
	"strconv"
	"./network"
)

var GetError = errors.New("get value error")
var PutError = errors.New("put error")

func Connect(port int) *rpc.Client {
	client, err := rpc.Dial("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	return client
}

func Get(key string) (string, error) {
	client := Connect(9588)
	getArg := network.GetArg{Key: key}
	response := network.GetResponse{}
	fmt.Println("Sending get request.")
	client.Call("KVHandler.Get", &getArg, &response)
	if response.StatusCode == network.SUCCESS {
		return response.Value, nil
	}
	return "", GetError
}

func Put(key string, value string) error {
	client := Connect(9588)
	putArg := network.PutArg{Key: key, Value: value}
	response := network.PutResponse{}
	fmt.Println("Sending put request")
	client.Call("KVHandler.Put", &putArg, &response)
	if response.StatusCode == network.SUCCESS {
		return nil
	}
	return PutError
}

