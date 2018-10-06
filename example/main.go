package main

import (
	"time"

	"github.com/theskyinflames/fifo/fifo"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	fifo := fifo.GetFifo(int32(4))

	fifo.Put(4)
	fifo.Put(3)
	fifo.Put(2)
	fifo.Put(1)

	go extract(logger, fifo)
	go extract(logger, fifo)
	go extract(logger, fifo)
	go extract(logger, fifo)

	time.Sleep(1 * time.Second)
}

func extract(logger *zap.Logger, fifo *fifo.Fifo) {
	item, err := fifo.Pop()
	if err != nil {
		panic(err)
	}
	logger.Info("retrieved", zap.Int("item", item.(int)))
}
