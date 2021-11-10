package main

import (
	"fmt"
)

func firstFunc(msg string) {
	// fmt.Println("First Function ")
	for i := 0; i < 16; i++ {
		fmt.Println(msg, ":", i)
	}
}

// func secondFunc(msg string) {
// 	// fmt.Println("Second Function ")
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(msg, ":", i)
// 	}
// }
func main() {
	fmt.Println("go rooutine and channel learning \n")

	message := make(chan string)
	go firstFunc("First Function")

	go func() {
		fmt.Println("Second Function ")
		message <- "ping second func by channel"
	}()

	m := <-message
	fmt.Println(m)

	fmt.Println("done")
}
