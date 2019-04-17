package main

import "fmt"

func main() {
    var values []int
    fmt.Println("Please input 10 numbers:")
    for i:=0;i<10;i++{
        in:=0
        fmt.Scanln(&in)
        values=append(values,in)
    }
    BubbleSort(values)
    fmt.Println(values)
}

func BubbleSort(a []int) {    
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a)-1-i; j++  {
            if a[j+1] < a[j] {
                Swap(a,j)
            }
        }
    }
}

func Swap(a []int, j int) {
    var temp int
    temp = a[j]
    a[j] = a[j+1]
    a[j+1] = temp
}