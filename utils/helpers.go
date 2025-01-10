package utils

func TransformList[T any, R any](input []T, transform func(T) R) []R {
	var result []R
	for _, item := range input {
		result = append(result, transform(item))
	}
	return result
}
