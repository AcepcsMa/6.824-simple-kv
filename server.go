package main

import (
	"./db"
	"fmt"
	"net"
	"net/rpc"
)

func StartServer() {
	kv := db.KVDB{}
	kv.Data = make(map[string]string)

	kvHandler := db.KVHandler{}
	kvHandler.Instance = &kv

	rpcServer := rpc.NewServer()
	rpcServer.Register(&kvHandler)

	l, err := net.Listen("tcp", ":9588")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				break
			} else {
				go rpcServer.ServeConn(conn)
			}
		}
		l.Close()
	}()
}

func main() {
	StartServer()
	fmt.Println("server starts.")

	Put("hh", "123")
	value, err := Get("hh")
	if err != nil {
		panic(err)
	}
	fmt.Printf("try to get \"hh\": %s\n", value)
}
