package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "io"
)
type Name struct {
	fname string
	lname string 
}
func main() {
    path:=""
    fmt.Println("Please input file full Path(Such as C:\\Windows\\win.ini):")
    fmt.Scanln(&path)
    var name[] Name
    f, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n')
        
        if err != nil || io.EOF == err {
            break
        }       
            tname:=strings.Split(line," ")
            tpname:= Name {
                tname[0],
                tname[1],
            } 
            name=append(name,tpname)
            }  
    for i:=0;i<len(name);i++{
        fmt.Println(i+1)
        fmt.Println("First Name:"+name[i].fname)
        fmt.Println("Last Name:"+name[i].lname)
    }
}