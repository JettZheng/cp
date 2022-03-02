package advanced

import "errors"

type Ring struct {
	rp   uint64
	size uint64
	wp   uint64
	mask uint64
	data []int
}

// NewRing new a ring buffer.
func NewRing(size int) *Ring {
	r := new(Ring)
	r.init(uint64(size))
	return r
}

// Init init ring.
func (r *Ring) Init(size int) {
	r.init(uint64(size))
}

func (r *Ring) init(num uint64) {
	// 2^N
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num--
		}
		num <<= 1
	}
	r.data = make([]int, num)
	r.size = num
	r.mask = r.size - 1
}

// Get get a proto from ring.
func (r *Ring) Get() (data int, err error) {
	if r.rp == r.wp {
		return 0, errors.New("empty")
	}

	return r.data[r.rp&r.mask],nil
}

// Set get a proto to write.
func (r *Ring) Set() (data int, err error) {
	if r.wp-r.rp >= r.size {
		return 0, errors.New("full")
	}
	return r.data[r.wp&r.mask],nil
}

// GetAdv incr read index.
func (r *Ring) GetAdv() {
	r.rp++
}

// SetAdv incr write index.
func (r *Ring) SetAdv() {
	r.wp++
}

// Reset reset ring.
func (r *Ring) Reset() {
	r.rp = 0
	r.wp = 0
}

func Ringbuffer() {
	r := NewRing(6)
	r.Set()
	r.Set()
	r.Set()
	r.Set()
	r.Set()
	r.Set()
	r.Set()
	r.Set()
	r.Set()
}
