package helper

func Reduce[SliceType any, ResultType any](slice []SliceType, reduceFunc func(acc ResultType, curr SliceType) ResultType, initialValue ResultType) (result ResultType) {
	result = initialValue
	for _, value := range slice {
		result = reduceFunc(result, value)
	}
	return result
}

func Clone[T any](slice []T) (clone []T) {
	clone = make([]T, len(slice))
	copy(clone, slice)
	return clone
}
