package main

import (
	"log"
	"fmt"
	"net/rpc"
	"crypto/tls"
)

// client call example

func main() {

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:9091", conf)
	if err != nil {
		log.Fatal("tcp error", err)
	}
	conn.Write([]byte("agent-code"))

	var buf = make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(buf[:n]))
	if string(buf[:n]) != "success" {
		log.Fatal("token error!")
		//os.Exit(1);
	}

	client := rpc.NewClient(conn)

	args := map[string]interface {}{
		"a": 7,
		"b": 8,
	}

	var reply int

	err = client.Call("Example.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args["a"], args["b"], reply)
}