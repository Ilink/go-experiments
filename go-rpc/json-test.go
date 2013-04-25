package main

import (
    simplejson "github.com/bitly/go-simplejson"
    "fmt"
    "os"
)

func main(){
    jsonStr := []byte(`{ 
        "test": { 
            "string_array": ["asdf", "ghjk", "zxcv"],
            "array": [1, "2", 3],
            "arraywithsubs": [{"subkeyone": 1},
            {"subkeytwo": 2, "subkeythree": 3}],
            "int": 10,
            "float": 5.150,
            "bignum": 9223372036854775807,
            "string": "simplejson",
            "bool": true 
        }
    }`)

    js, err := simplejson.NewJson(jsonStr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Bad json!")
        os.Exit(1)
    }

    for k, v := range js.Get("test") {
        fmt.Printf("%v\n", k)
    }

    // for key in js['test']:
        // print something
    test, _ := js.Get("test").Get("int").Int()
    fmt.Printf("%v\n", test)

}