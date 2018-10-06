package fifo

import (
	"errors"
	"sync"
	"time"
)

var putMutex *sync.Mutex = &sync.Mutex{}
var popMutex *sync.Mutex = &sync.Mutex{}

func GetFifo(sz int32) *Fifo {
	return &Fifo{c_fifo: make(chan interface{}, sz),
		size:      0,
		maxLength: sz,
		id:        time.Now().UnixNano(),
	}
}

type Fifo struct {
	id        int64
	c_fifo    chan interface{}
	size      int32
	maxLength int32
}

func (s *Fifo) Empty() bool { return s.size == 0 }
func (s *Fifo) Peek() interface{} {
	putMutex.Lock()
	defer putMutex.Unlock()

	item := <-s.c_fifo
	s.c_fifo <- item
	return item
}
func (s *Fifo) Len() int32 { return s.size }

func (s *Fifo) Put(i interface{}) error {
	putMutex.Lock()
	defer putMutex.Unlock()

	if s.size == s.maxLength {
		return errors.New("max length reached")
	}

	s.size += 1
	s.c_fifo <- i

	return nil
}

func (s *Fifo) Pop() (interface{}, error) {
	popMutex.Lock()
	defer popMutex.Unlock()

	if s.size == 0 {
		return nil, errors.New("the fifo is empty")
	}

	s.size -= 1
	item := <-s.c_fifo
	return item, nil
}
