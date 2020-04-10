 #  stream-go

A Stream library fo Go

## What to do now?

Create a Stream
```go
import (
	streams "github.com/clubifaximatic/stream-go"
)

stream := streams.NewStream()
```

There is an interface:
```go
type Stream interface {
	Filter(fn StreamFilterFunction) Stream
	Map(fn StreamMapFunction) Stream
	FlatMap(fn StreamFlatMapFunction) Stream
	Peek(fn StreamPeekFunction) Stream
	Distinct() Stream
	Skip(n int) Stream
	Limit(n int) Stream
	ForEach(fn StreamForEachFunction)
}
```

Use it:
```go
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
```


# What is missing?

* finish gracefully
* test
* stream connectors for slices, etc
* merge/split streams
* parallel execution
