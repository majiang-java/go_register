package main

import (
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()
	// wg.Wait()
	//o count("fish")
	c := make(chan string)
	go count("sheep", c)
	msg := <-c
	fmt.Println(msg)
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
}
