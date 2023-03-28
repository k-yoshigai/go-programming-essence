package main

import (
	"fmt"
	"time"
)

func fanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				new_ch <- s
			case s := <-ch2:
				new_ch <- s
			}
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
