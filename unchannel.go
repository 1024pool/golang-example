package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 10)
    c <- 1
    c <- 2
    c <- 3
    close(c)

    for {
        fmt.Println(<-c)
    }
} 
