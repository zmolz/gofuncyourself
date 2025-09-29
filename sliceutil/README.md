# Slice Util 

## Functional ...functions..?
```go
func Filter[T any](xs []T, pred func(T) bool) []T
func Partition[T any](xs []T, pred func(T) bool) ([]T, []T)

func FoldLeft[T any, U any](xs []T, a U, f func(T, U) U) U 
func Reduce[T any](xs []T, f func(T, T) T) (T, error) 
func FoldRight[T any, U any](xs []T, a U, f func(T, U) U) U

func Map[T any, U any](xs []T, f func(T) U) []U
```

## Basic Utility
```go
func Remove[T any](xs []T, i int) []T
func RemoveUnordered[T any](xs []T, i int) ([]T, error)
func Pop[T any](xs []T) ([]T, T, error)
func RemoveLast[T any](xs []T) ([]T, error) 
```