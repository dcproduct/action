package search

import (
	"fmt"
	"log"
)

// Result is search struct
type Result struct {
	Field   string
	Content string
}

// Matcher is an interface which includes a method 'Search'.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match perform match logic base on feed and search term.
func Match(matcher Matcher, feed *Feed, searchTerm string,
	results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

// Display all search results to stdout.
func Display(results chan *Result) {
	for result := range results {
		fmt.Println("%s:\n%s\n\n", result.Field, result.Content)
	}
}
