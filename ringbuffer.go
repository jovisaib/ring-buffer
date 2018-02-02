package water

type RingBuffer struct {
	items      []int
	inPointer  int
	outPointer int
	size       int
}

func (cb *RingBuffer) New(size int) *RingBuffer {
	cb.items = make([]int, size)
	cb.inPointer = 0
	cb.outPointer = 0
	cb.size = size
	return cb
}

func (cb *RingBuffer) Push(val int) bool {
	cb.items[cb.inPointer] = val
	cb.inPointer++
	cb.inPointer %= cb.size
	return cb.IsFull()
}

func (cb *RingBuffer) Pop(val *int) bool {
	isEmpty := cb.IsEmpty()
	if !isEmpty {
		val = &cb.items[cb.outPointer]
		cb.outPointer++
		cb.outPointer %= cb.size
	}
	return isEmpty
}

func (cb *RingBuffer) GetSlice() []int {
	slice := []int{}
	slice = append(slice, cb.items[cb.inPointer:]...)
	slice = append(slice, cb.items[:cb.inPointer]...)
	return slice
}

func (cb *RingBuffer) Clear() {
	cb.inPointer = 0
	cb.outPointer = 0
}

func (cb *RingBuffer) Size() int {
	return cb.size
}

func (cb *RingBuffer) IsFull() bool {
	return ((cb.inPointer + 1) % cb.size) == cb.outPointer
}

func (cb *RingBuffer) IsEmpty() bool {
	return cb.inPointer == cb.outPointer
}
