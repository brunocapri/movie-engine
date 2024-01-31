package utils

func Prepend[T any](slice []T, elems ...T) []T {
	return append(elems, slice...)
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
