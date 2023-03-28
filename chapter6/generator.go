package main

import (
	"fmt"
	"time"
)

func fanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			new_ch <- <-ch1
		}
	}()
	go func() {
		for {
			new_ch <- <-ch2
		}
	}()
	return new_ch
}

func generator(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	g1 := generator("Hello")
	g2 := generator("Bye")
	ch := fanIn(g1, g2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
