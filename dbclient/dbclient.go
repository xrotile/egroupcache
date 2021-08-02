package client

import (
	"fmt"
	"log"
	"net/rpc"
)

type Client struct {
	Host string
	Port int
}

func NewClient(host string, port int) *Client {
	return new(Client{
		Host: host,
		Port: port,
	})
}

func (client *Client) Get(key string) string {
	address := fmt.Sprintf("%s:%d", client.Host, client.Port)
	cl, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	args := new(api.Load{
		Key: key,
	})
	err = cl.Call("DBServer.Get", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func (client *Client) Put(key string, value string) {
	address := fmt.Sprintf("%s:%d", client.Host, client.Port)
	cl, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply int
	args := new(api.Store{
		Key:   key,
		Value: value,
	})
	err = cl.Call("DBServer.Put", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}
