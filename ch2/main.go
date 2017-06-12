package main

import (
	"log"
	"os"

	_ "github.com/dcproduct/action/ch2/matchers"
	"github.com/dcproduct/action/ch2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

// Add comment to test.
func main() {
	search.Run("president")
}
