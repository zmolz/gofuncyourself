package iterator

import "fmt"

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type NumRange[T Num] struct {
	end  T
	curr T
	step T

	// Zero value purposefully not intialized for generic comparison against zero of
	// whichever numeric type we are using in this instance.
	zero T
}

var ErrZeroStep = fmt.Errorf("step cannot be zero")

// NewNumRange creates a new numeric range iterator
func NewNumRange[T Num](start, end, step T) (*NumRange[T], error) {
	if step == 0 {
		return nil, ErrZeroStep
	}

	return &NumRange[T]{
		end:  end,
		curr: start - step, // first call to Next() gives start
		step: step,
	}, nil
}

func (n *NumRange[T]) HasNext() bool {
	next := n.curr + n.step
	if n.step > n.zero {
		return next < n.end
	}
	return next > n.end
}

func (n *NumRange[T]) Next() (T, bool) {
	if !n.HasNext() {
		return n.zero, false
	}

	n.curr += n.step
	return n.curr, true
}
