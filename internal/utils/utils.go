package utils

func Prepend[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}
