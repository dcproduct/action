package main

import (
	"fmt"
)

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	// duration(40).pretty()
	d := duration(40)

	fmt.Println((&d).pretty())

}
