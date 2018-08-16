package main

import (
    "fmt"
    "time"
    "strconv"
)

var c1 = make(chan string)

func main() {
    go produce()
    for i := 0; i < 10; i++ {
        go consume(i)
    }
    var input string
    fmt.Scanln(&input)
}

func produce() {
    for {
       c1 <- "message"
       time.Sleep(time.Second * 1)
    }
}

func consume(id int) {
    for {
        msg := <- c1 + " received by " + strconv.Itoa(id) + " at " +  time.Now().String()
        fmt.Println(msg)
        time.Sleep(time.Second * 2)
    }
}
