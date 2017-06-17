package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	num := int(1)
	var err error
	if len(os.Args) > 1 {
		num, err = strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}
	}

	runtime.GOMAXPROCS(num)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Strat Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
