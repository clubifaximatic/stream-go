package streams

type StreamFilterFunction func(interface{}) bool
type StreamForEachFunction func(interface{})
type StreamMapFunction func(interface{}) interface{}
type StreamFlatMapFunction func(interface{}) []interface{}
type StreamPeekFunction func(interface{}) interface{}

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

type stream struct {
	in   chan interface{}
	out  chan interface{}
	next *stream
}

func NewStream() *stream {

	return &stream{
		in: make(chan interface{}),
	}
}

func (s *stream) Filter(fn StreamFilterFunction) Stream {
	nextStream := s.nextStream()

	go func() {
		for e := range s.in {
			if fn(e) {
				s.out <- e
			}
		}
	}()

	return nextStream
}

func (s *stream) Map(fn StreamMapFunction) Stream {
	nextStream := s.nextStream()

	go func() {
		for e := range s.in {
			s.out <- fn(e)
		}
	}()

	return nextStream
}

func (s *stream) FlatMap(fn StreamFlatMapFunction) Stream {
	nextStream := s.nextStream()

	go func() {
		for e := range s.in {
			items := fn(e)
			for item := range items {
				s.out <- item
			}
		}
	}()

	return nextStream
}

func (s *stream) Peek(fn StreamPeekFunction) Stream {
	go func() {
		for e := range s.in {
			fn(e)
		}
	}()

	return s
}

func (s *stream) Distinct() Stream {
	nextStream := s.nextStream()

	go func() {
		elements := make(map[interface{}]struct{})
		for e := range s.in {
			if _, ok := elements[e]; ok {
				continue
			}
			elements[e] = struct{}{}
			s.out <- e
		}
	}()

	return nextStream
}

func (s *stream) Skip(n int) Stream {
	nextStream := s.nextStream()

	go func() {
		remain := n
		for e := range s.in {
			if remain > 0 {
				remain--
				continue
			}
			s.out <- e
		}
	}()

	return nextStream
}

func (s *stream) Limit(n int) Stream {
	nextStream := s.nextStream()

	go func() {
		remain := n
		for e := range s.in {
			if remain <= 0 {
				continue
			}
			remain--
			s.out <- e
		}
	}()

	return nextStream
}

func (s *stream) ForEach(fn StreamForEachFunction) {
	go func() {
		for e := range s.in {
			fn(e)
		}
	}()
}

func (s *stream) Push(e interface{}) {
	s.in <- e
}

func (s *stream) nextStream() *stream {
	nextStream := NewStream()
	s.out = nextStream.in
	return nextStream
}
