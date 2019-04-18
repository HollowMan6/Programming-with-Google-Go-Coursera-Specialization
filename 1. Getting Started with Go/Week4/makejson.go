package main

import (
    "fmt"
    "encoding/json"
    "os"
    "bufio"
)

func main() {
    var m map[string]string
    m = make(map[string]string)
    fmt.Println("Please input your name:")
    br := bufio.NewReader(os.Stdin)
    bname, _, _ := br.ReadLine()
    name:=string(bname)
    m["name"]=name
    fmt.Println("Then input your adress:")
    br1 := bufio.NewReader(os.Stdin)
    badress, _, _ := br1.ReadLine()
    adress:=string(badress)
    m["adress"]=adress
    b, err := json.Marshal(m)
    if err != nil {
        fmt.Println("Encoding faild")
    } else {
        fmt.Println("Encoded data : ")
        fmt.Println(b)
        fmt.Println("Decoded data : ")
        fmt.Println(string(b))
    }
}