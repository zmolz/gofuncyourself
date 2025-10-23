package sliceutil

type Pair[T, U any] struct {
	First  T
	Second U
}

// Zip takes two slices, xs and ys, and returns a slice of Pairs,
// where each Pair contains an element from xs and the element at same index from ys.
// If the slices are of unequal lengths, the longer slice's extra items will be ignored.
func Zip[T, U any](xs []T, ys []U) []Pair[T, U] {
	l := min(len(xs), len(ys))
	pairs := make([]Pair[T, U], l)

	for i := range l {
		pairs[i] = Pair[T, U]{First: xs[i], Second: ys[i]}
	}

	return pairs
}

// ZipLongest takes two slices, xs and ys, and returns a slice of Pairs,
// where each Pair contains an element from xs and the element at same index from ys.
// If the slices are of unequal lengths, the shorter slice will be padded to be equal length.
func ZipLongest[T, U any](xs []T, ys []U) []Pair[T, U] {
	total := max(len(xs), len(ys))
	pairs := make([]Pair[T, U], total)

	for i := range total {
		var x T
		var y U

		if i < len(xs) {
			x = xs[i]
		}
		if i < len(ys) {
			y = ys[i]
		}

		pairs[i] = Pair[T, U]{First: x, Second: y}
	}

	return pairs
}

// Zip takes two slices, xs and ys, and returns a slice of the result of appling function f to each pair,
// If the slices are of unequal lengths, the longer slice's extra items will be ignored.
func ZipWith[T, U, V any](xs []T, ys []U, f func(T, U) V) []V {
	l := min(len(xs), len(ys))
	zs := make([]V, l)

	for i := range l {
		zs[i] = f(xs[i], ys[i])
	}

	return zs
}

// Unzip takes a slice of pairs and returns two slices
// - one corresponding to all the first elements in the pairs, and the other to the second element.
func Unzip[T, U any](pairs []Pair[T, U]) ([]T, []U) {
	xs := make([]T, len(pairs))
	ys := make([]U, len(pairs))

	for i, pair := range pairs {
		xs[i] = pair.First
		ys[i] = pair.Second
	}

	return xs, ys
}
