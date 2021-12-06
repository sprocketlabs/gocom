package gocom

import (
	"bytes"
	"fmt"
)

// Cqueue dataresents a double-ended Cqueue.
// The zero value is an empty Cqueue ready to use.
type Cqueue struct {
	// PushBack writes to data[back] then increments back; PushFront
	// decrements front then writes to data[front]; len(data) is a power
	// of two; unused slots are nil and not garbage.
	isStack bool
	data    []interface{}
	front   int
	back    int
	length  int
}

// New returns an initialized empty CCqueue.
func New() *Cqueue {
	return new(Cqueue).Init()
}

// Init initializes or clears Cqueue q.
func (q *Cqueue) Init() *Cqueue {
	q.data = make([]interface{}, 1)
	q.front, q.back, q.length = 0, 0, 0
	return q
}

// lazyInit lazily initializes a zero Cqueue value.
//
// I am mostly doing this because container/list does the same thing.
// Personally I think it's a little wasteful because every single
// PushFront/PushBack is going to pay the overhead of calling this.
// But that's the price for making zero values useful immediately.
func (q *Cqueue) lazyInit() {
	if q.data == nil {
		q.Init()
	}
}

// Len returns the number of elements of Cqueue q.
func (q *Cqueue) Len() int {
	return q.length
}

// empty returns true if the Cqueue q has no elements.
func (q *Cqueue) empty() bool {
	return q.length == 0
}

// full returns true if the Cqueue q is at capacity.
func (q *Cqueue) full() bool {
	return q.length == len(q.data)
}

// sparse returns true if the Cqueue q has excess capacity.
func (q *Cqueue) sparse() bool {
	return 1 < q.length && q.length < len(q.data)/4
}

// resize adjusts the size of Cqueue q's underlying slice.
func (q *Cqueue) resize(size int) {
	adjusted := make([]interface{}, size)
	if q.front < q.back {
		// data not "wrapped" around, one copy suffices
		copy(adjusted, q.data[q.front:q.back])
	} else {
		// data is "wrapped" around, need two copies
		n := copy(adjusted, q.data[q.front:])
		copy(adjusted[n:], q.data[:q.back])
	}
	q.data = adjusted
	q.front = 0
	q.back = q.length
}

// lazyGrow grows the underlying slice if necessary.
func (q *Cqueue) lazyGrow() {
	if q.full() {
		q.resize(len(q.data) * 2)
	}
}

// lazyShrink shrinks the underlying slice if advisable.
func (q *Cqueue) lazyShrink() {
	if q.sparse() {
		q.resize(len(q.data) / 2)
	}
}

// String returns a string dataresentation of Cqueue q formatted
// from front to back.
func (q *Cqueue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.length; i++ {
		result.WriteString(fmt.Sprintf("%v", q.data[j]))
		if i < q.length-1 {
			result.WriteByte(' ')
		}
		j = q.inc(j)
	}
	result.WriteByte(']')
	return result.String()
}

// inc returns the next integer position wrapping around Cqueue q.
func (q *Cqueue) inc(i int) int {
	return (i + 1) & (len(q.data) - 1) // requires l = 2^n
}

// dec returns the previous integer position wrapping around Cqueue q.
func (q *Cqueue) dec(i int) int {
	return (i - 1) & (len(q.data) - 1) // requires l = 2^n
}

// Front returns the first element of Cqueue q or nil.
func (q *Cqueue) Front() interface{} {
	// no need to check q.empty(), unused slots are nil
	return q.data[q.front]
}

// Back returns the last element of Cqueue q or nil.
func (q *Cqueue) Back() interface{} {
	// no need to check q.empty(), unused slots are nil
	return q.data[q.dec(q.back)]
}

func (q *Cqueue) Push(v interface{}) {
	if q.isStack {
		q.pushBack(v)
	} else {
		panic("Not a stack")
	}
}

func (q *Cqueue) Pop(v interface{}) {
	if q.isStack {
		q.popBack()
	} else {
		panic("Not a stack")
	}
}

func (q *Cqueue) Enqueue(v interface{}) {
	if !q.isStack {
		q.pushFront(v)
	}
	panic("Not a stack")
}

func (q *Cqueue) Dequeue(v interface{}) {
	if !q.isStack {
		q.popBack()
	}
	panic("Not a stack")
}

// PushFront inserts a new value v at the front of Cqueue q.
func (q *Cqueue) pushFront(v interface{}) {
	q.lazyInit()
	q.lazyGrow()
	q.front = q.dec(q.front)
	q.data[q.front] = v
	q.length++
}

// PushBack inserts a new value v at the back of Cqueue q.
func (q *Cqueue) pushBack(v interface{}) {
	q.lazyInit()
	q.lazyGrow()
	q.data[q.back] = v
	q.back = q.inc(q.back)
	q.length++
}

// PopFront removes and returns the first element of Cqueue q or nil.
func (q *Cqueue) popFront() interface{} {
	if q.empty() {
		return nil
	}
	v := q.data[q.front]
	q.data[q.front] = nil // unused slots must be nil
	q.front = q.inc(q.front)
	q.length--
	q.lazyShrink()
	return v
}

// PopBack removes and returns the last element of Cqueue q or nil.
func (q *Cqueue) popBack() interface{} {
	if q.empty() {
		return nil
	}
	q.back = q.dec(q.back)
	v := q.data[q.back]
	q.data[q.back] = nil // unused slots must be nil
	q.length--
	q.lazyShrink()
	return v
}
