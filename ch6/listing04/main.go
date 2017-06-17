package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	num := runtime.NumCPU()

	fmt.Printf("Numer of cup is: %d\n", num)
	runtime.GOMAXPROCS(num)

	wg.Add(2)

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
