package streams

func NewStream() Stream {
	return &stream{
		in: make(chan interface{}),
	}
}

func NewStreamFromSlice(input []interface{}) Stream {
	s := newStream()

	go func() {
		defer close(s.in)
		for _, e := range input {
			s.in <- e
		}
	}()

	return s
}

func NewStreamFromFeeder(fn func() interface{}) Stream {
	s := newStream()

	go func() {
		defer close(s.in)
		for {
			e := fn()
			if e == nil {
				break
			}
			s.in <- e
		}
	}()

	return s
}

