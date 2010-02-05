package main

import "fmt"

/*
 * unbuffered channel 
 */
var c = make(chan int, 0)

/*
 * buffered channel You can change the 1 to 10, 50, 100, or 500
 */
//var c = make(chan int, 1)

var a string

func callstart(name string) {
   fmt.Printf("%s Start\n", name) 
}

func callend(name string) {
   fmt.Printf("%s End\n", name) 
}

func f() {
    callstart("func")
    a = "\nHello, world\n"
    for i:=0; i<300; i++ {
        c <- i  // Receive on C
    }
    callend("\nfunc")
}

func main() {
    callstart("main")
    go f()
    for i:=0; i<300; i++ {
            fmt.Printf("%d ", <-c)  // Send on C
    }
    fmt.Printf(a)
    callend("main")
}
