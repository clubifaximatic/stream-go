package main

import (
	"fmt"
	"time"

	streams "github.com/clubifaximatic/stream-go"
)

func main() {

	current := 0

	next := func() interface{} {
		current++
		if current > 50 {
			return nil
		}
		return current
	}

	// Create the stream
	stream := streams.NewStreamFromFeeder(next)

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

		// remove duplicates
		Distinct().

		// fetch only 10 results
		Limit(10).

		// peek into the value
		Peek(func(i interface{}) { fmt.Println("* ", i) }).

		// skip 5 items
		Skip(5).

		// terminal operator
		ForEach(func(e interface{}) {
			fmt.Println(e)
		})

	time.Sleep(5 * time.Second)
	fmt.Println("done!")
}
