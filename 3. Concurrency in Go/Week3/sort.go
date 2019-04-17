package main

import (
	"fmt"
	"sync"
)

func QuickSort(data []int) []int {
    if len(data) <= 1 {
        return data
    }
    var wg sync.WaitGroup
    c := data[0]
    var s1, s2 []int

    for k, v := range data {
        if k == 0 {
            continue
        }
        if c < v {
            s2 = append(s2, v)
        } else {
            s1 = append(s1, v)
        }
    }

    wg.Add(2)
    go func() {
        s1 = QuickSort(s1)
        wg.Done()
    }()
    go func() {
        s2 = QuickSort(s2)
        wg.Done()
    }()
    wg.Wait()

    data = []int{c}
    if len(s1) > 0 {
        data = append(s1, data...)
    }
    if len(s2) > 0 {
        data = append(data, s2...)
    }
    return data
}

func main() {
    var values []int
    fmt.Println("Please input 8 numbers:")
    for i:=0;i<8;i++{
        in:=0
        fmt.Scanln(&in)
        values=append(values,in)
	}
	data1:=values[0:len(values)/4]
	data2:=values[len(values)/4:len(values)/2]
	data3:=values[len(values)/2:len(values)/4*3]
	data4:=values[len(values)/4*3:len(values)]
	fmt.Println(QuickSort(data1))
	fmt.Println(QuickSort(data2))
	fmt.Println(QuickSort(data3))
	fmt.Println(QuickSort(data4))
	fmt.Println(QuickSort(values))
}