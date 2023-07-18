package main

import (
	"fmt"
	"sync"
)

func main() {
	chans := makeWritingChanSlice()

	oneChan := chanalsToOne(chans...)

	for v := range oneChan {
		fmt.Println(v)
	}
}

func chanalsToOne(chans ...chan int) <-chan int {
	oneChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}

		for _, v := range chans {
			wg.Add(1)
			go func(ch chan int) {
				for val := range ch {
					oneChan <- val
				}
				wg.Done()
			}(v)
		}
		wg.Wait()
		close(oneChan)
	}()

	return oneChan
}


func makeWritingChanSlice() []chan int {
	chan1 := make(chan int)
	go func() {
		for _, v := range []int{1, 2, 3, 4, 5} {
			chan1 <- v
		}
		close(chan1)
	}()
	chan2 := make(chan int)
	go func() {
		for _, v := range []int{1, 2, 3, 4, 5} {
			chan2 <- v
		}
		close(chan2)
	}()
	chan3 := make(chan int)
	go func() {
		for _, v := range []int{1, 2, 3, 4, 5} {
			chan3 <- v
		}
		close(chan3)
	}()

	return []chan int{chan1, chan2, chan3}
}