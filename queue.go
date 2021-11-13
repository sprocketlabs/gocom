package gocom

import (
	"bytes"
)

type CgQueue struct {
	data []interface{}
	length int
	front  int
	back   int
}

func NewCgQueue() *CgQueue {
	q := new(CgQueue)
	q.data = make([]interface{}, 1)

	q.front = 0
	q.back = 0
	q.length = 0

	return q
}

// Len returns the number of elements of queue q.
func (q *CgQueue) CgQueueLen() int {
	return q.length
}

func (q *CgQueue) CgQueueIsEmpty() bool {
	return q.length == 0
}

func (q *CgQueue) resize() {

}

func (q *CgQueue) resize(size int) {
	new_size := q.length
	if (q.length == len(q.data)) {
		newSize = newSize * 2
	} else if (1 < q.length && q.length < len(q.rep)/4) {
		newSize = newSize / 2
	}

	if new_size != q.length {
		adjusted := make([]interface{}, size)
		if q.front < q.back {
			// rep not "wrapped" around, one copy suffices
			copy(adjusted, q.rep[q.front:q.back])
		} else {
			// rep is "wrapped" around, need two copies
			n := copy(adjusted, q.rep[q.front:])
			copy(adjusted[n:], q.rep[:q.back])
		}
		q.data = adjusted
		q.front = 0
		q.back = q.length
	}
}

func (q *Queue) PushFront(v interface{}) {
	q.lazyInit()
	q.lazyGrow()
	q.front = q.dec(q.front)
	q.rep[q.front] = v
	q.length++
}

func (q *Queue) ToString() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.length; i++ {
		result.WriteString(fmt.Sprintf("%v", q.rep[j]))
		if i < q.length-1 {
			result.WriteByte(' ')
		}
		j = q.inc(j)
	}
	result.WriteByte(']')
	return result.String()
}
