package util

import "log"

func ContainsForArrayString(target string, list []string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}

func SplitArray[T any](array []T, step int) [][]T {
	if step < 1 {
		log.Fatalf("split array not support step: %d", step)
	}
	var result [][]T
	for i := 0; i < len(array); i += step {
		end := i + step
		if end > len(array) {
			end = len(array)
		}
		result = append(result, array[i:end])
	}
	return result
}
