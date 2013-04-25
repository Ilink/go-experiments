package main

import (
    "net/rpc/jsonrpc"
    "net/rpc"
    "net"
    "log"
    // "encoding/json"
    // simplejson "github.com/bitly/go-simplejson"
)

type RPC uint8

type Resp struct {
    Page   string      `json:"page"`
    Fruits []string `json:"fruits"`
}

func(*RPC) Echo(arg *string, result *interface{}) error {
    m := map[string]int {
        "a": 1,
        "b": 2,
    }

    // res2D := &Resp{
    //     Page:   "1",
    //     Fruits: []string{"apple", "peach", "pear"}}
    // res2B, _ := json.Marshal(res2D)

    // encoded, _ := json.Marshal(m)
    // log.Print(string(encoded))
    // *result = string(encoded)
    *result = m
    log.Printf("Arg: %v, Return: %v", *arg, *result)

    return nil
}

func main(){
    sock, err := net.Listen("tcp", "localhost:1234")
    defer sock.Close()
    if err != nil {
        log.Fatal(err)
    }
    log.Print("listening on: ", sock.Addr())
    rpc.Register(new (RPC))

    for {
        log.Print("waiting for connections")
        conn, err := sock.Accept()
        if err != nil {
            log.Printf("connection error: %s", conn)
            continue
        }
        
        log.Printf("connection started: %v", conn.RemoteAddr())
        go jsonrpc.ServeConn(conn)
    }
}