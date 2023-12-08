package input

type RingBuffer struct {
	buffer string
	index  int
	length int
}

func (r *RingBuffer) Next() byte {
	result := r.buffer[r.index]
	r.index += 1
	if r.index >= r.length {
		r.index = 0
	}
	return result
}

func NewRingBuffer(s string) RingBuffer {
	return RingBuffer{
		buffer: s,
		length: len(s),
	}
}
