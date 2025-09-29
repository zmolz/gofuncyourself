package sliceutil

const (
	PartitionSliceInitFactor = 2
)

// Filter returns all elements of slice xs for which predicate function pred is true
func Filter[T any](xs []T, pred func(T) bool) []T {
	ret := []T{}

	for _, x := range xs {
		if pred(x) {
			ret = append(ret, x)
		}
	}

	return ret
}

// Partition split slice xs into two slices,
// the first being all elements of xs for which predicate function pred is true,
// the second being all elements for which pred is false.
func Partition[T any](xs []T, pred func(T) bool) ([]T, []T) {
	trueElements := []T{}
	falseElements := []T{}

	for _, x := range xs {
		if pred(x) {
			trueElements = append(trueElements, x)
		} else {
			falseElements = append(falseElements, x)
		}
	}
	
	return trueElements, falseElements
}
