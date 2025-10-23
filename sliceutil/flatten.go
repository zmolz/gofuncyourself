package sliceutil

// Flatten takes a slice of dimension >= 2, and flattens it by 1 dimension
func Flatten[T any](xss [][]T) []T {
	xs := []T{}

	for _, inner := range xss {
		xs = append(xs, inner...)
	}

	return xs
}

// // ArbitraryFlatten takes a slice of dimension >= 2, and flattens it to 1 dimension
// func ArbitraryFlatten(input any) (out []any) {
// 	return
// }