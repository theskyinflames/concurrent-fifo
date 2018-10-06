package fifo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_fifo(t *testing.T) {

	fifo := GetFifo(int32(4))

	fifo.Put(4)
	fifo.Put(3)
	fifo.Put(2)
	fifo.Put(1)

	go extract(t, fifo)
	go extract(t, fifo)
	go extract(t, fifo)
	go extract(t, fifo)

	time.Sleep(1 * time.Second)
}

func extract(t *testing.T, fifo *Fifo) {
	item, err := fifo.Pop()
	assert.NoError(t, err)
	t.Log(item.(int))
}
