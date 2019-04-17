package main

import (
    "fmt"
)

func main() {
    var number int
    fmt.Print("Please enter a floating point number:\n")
    fmt.Scanln(&number)
    fmt.Print("\nConverted interger:\n",number)
}