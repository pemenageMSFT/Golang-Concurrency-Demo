package main

import (
	"fmt"
	"sync"
)

func main() {
	// channel for even numbers
	even := make(chan int)
	// channel for odd numbers
	odd := make(chan int)
	// channel for values
	valchan := make(chan int)
	// logic for channels
	go chanlogic(even, odd)
	// transfer values
	go chansfer(even, odd, valchan)

	// range over values to see if the number is a multiple of three
	for v := range valchan {
		if v%3 == 0 && v != 0 {
			fmt.Println(v, "Is a multiple of three")
		} else {
			fmt.Println(v, "Is not a multiple of three")
		}

	}

	fmt.Println("about to exit")
}

// channel logic - chanlogic even values to even channel, chanlogic odd values to odd channel
func chanlogic(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// Transfer values from even and odd channels to valchan
func chansfer(even, odd <-chan int, valchan chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			valchan <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			valchan <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(valchan)
}
