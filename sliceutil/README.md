# Slice Util

```go
func Filter[T any](xs []T, pred func(T) bool) []T
func Partition[T any](xs []T, pred func(T) bool) ([]T, []T)
func GroupBy[T any, U comparable](xs []T, keyFunc func(T) U) map[U][]T

func FoldLeft[T, U any](xs []T, a U, f func(T, U) U) U
func Reduce[T any](xs []T, f func(T, T) T) (T, error)
func FoldRight[T, U any](xs []T, a U, f func(T, U) U) U

func Map[T, U any](xs []T, f func(T) U) []U

func Flatten[T any](xss [][]T) []T

func FilterMap[T, U any](xs []T, f func(T) (U, bool) ) []U
func FlatMap[T, U any](xs []T, f func(T) []U) []U 

func Compose[T, U, V any](f func(T) U, g func(U) V) func(T) V

func Find[T any](xs []T, f func(T) bool) (T, bool)
func FindAll[T any](xs []T, f func(T) bool) []T
func FindIndex[T any](xs []T, f func(T) bool) (int, bool) 
func Contains[T cmp.Ordered](xs []T, toFind T) bool
func Any[T any](xs []T, f func(T) bool) bool
func All[T any](xs []T, f func(T) bool) bool

func Remove[T any](xs []T, i int) []T
func RemoveUnordered[T any](xs []T, i int) ([]T, error)
func Pop[T any](xs []T) ([]T, T, error)
func RemoveLast[T any](xs []T) ([]T, error)
func RemoveIndices[T any](xs []T, indices []int) ([]T, error) 

func Take[T any](xs []T, count int) ([]T, error)
func Drop[T any](xs []T, count int) ([]T, error)

func Zip[T, U any](xs []T, ys []U) []Pair[T, U] 
func ZipLongest[T, U any](xs []T, ys []U) []Pair[T, U]
func ZipWith[T, U, V any](xs []T, ys []U, f func(T, U) V) []V 
func Unzip[T, U any](pairs []Pair[T, U]) ([]T, []U)
```
