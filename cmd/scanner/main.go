package main

import(
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        fmt.Fprintf(os.Stdout, scanner.Text())
    }
}
