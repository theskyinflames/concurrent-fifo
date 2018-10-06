# fifo 
It's a simple FIFO thread safe implementation

## Usage
```go
func main() {
	fifo := GetFifo(int32(4))

	fifo.Put(4)
	fifo.Put(3)
	fifo.Put(2)
	fifo.Put(1)
    
    println(fifo.Pop().(int))
	println(fifo.Pop().(int))
	println(fifo.Pop().(int))
    println(fifo.Pop().(int))
}
```