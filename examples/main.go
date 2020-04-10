package main

import (
	"fmt"

	streams "github.com/clubifaximatic/stream-go"
)

func main() {
	// Create the stream
	stream := streams.NewStream()

	stream.
		// filter only even numbers
		Filter(func(i interface{}) bool {
			n, ok := i.(int)
			if !ok {
				return false
			}

			return n%2 == 0
		}).

		// map number to text
		Map(func(i interface{}) interface{} {
			n, ok := i.(int)
			if !ok {
				n = 0
			}
			return fmt.Sprintf("number: %d", n)
		}).

		// fetch only 10 results
		Limit(10).

		// skip 5 items
		Skip(5).

		// remove duplicates
		Distinct().

		// terminal operator
		ForEach(func(e interface{}) {
			fmt.Println(e)
		})

	// start feeding the stream
	for i := 1; i < 1000; i++ {
		stream.Push("string")
		stream.Push(i * 3)
		stream.Push(i * 3)
	}
}
