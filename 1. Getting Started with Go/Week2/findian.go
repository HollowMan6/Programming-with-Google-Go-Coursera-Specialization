package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string
    fmt.Print("Please enter a string:\n")
    fmt.Scanln(&str)
    if strings.Contains(str,"a") && (str[0]=='i' || str[0]=='I') && (str[len(str)-1]=='n' || str[len(str)-1]=='N'){
        fmt.Print("Found!")
    }else{
        fmt.Print("Not Found!")
    }
}