package main

import (
	"fmt"
	"sync"
)


func makeChan(ar ... int) <-chan int {

	x := make(chan int)
	go func() {
		for _, i := range ans {
			x <- i
		}
		close(x)
	}()
	return x
}
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	channel := make(chan int)
	
  for _, c := range cs {
		wg.Add(1)
		
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				channel <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	return channel
}

func main() {
	a1 := makeChan(1, 6, 1, 5)
	a2 := makeChan(8, 3, 56, 4)
	a3 := makeChan(53, 9, 10, 11)
	for i := range merge(a1, a2, a3) {
		fmt.Println(i)
	}
}
