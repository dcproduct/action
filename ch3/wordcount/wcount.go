/*package wordcount provides count function to calculate
how many words within a file. Usage as below:
$> wordcount filename

Then program will print the result.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dcproduct/action/ch3/words"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your '%s'. \n", count, filename)
}
