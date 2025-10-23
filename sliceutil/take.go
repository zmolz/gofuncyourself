package sliceutil

// Take returns the first count elements from slice xs.
func Take[T any](xs []T, count int) ([]T, error) {
	if count > len(xs) {
		// should we maybe return a copy of xs?
		return nil, ErrIndexOutOfRange // Should this really be an out of range error?
	}

	takenXs := make([]T, count)
	copy(takenXs, xs[:count])
	return takenXs, nil
}

// Drop returns slice xs excluding the first count xs
func Drop[T any](xs []T, count int) ([]T, error) {
	diff := len(xs) - count
	if diff < 0 {
		// should we maybe return an empty list?
		return nil, ErrIndexOutOfRange
	}

	keptXs := make([]T, diff)
	copy(keptXs, xs[count:])
	return keptXs, nil
}
