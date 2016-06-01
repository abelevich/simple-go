package main

import(
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    
    //Print arguments
    for _, arg := range args {
        fmt.Println(arg)
    }
}
