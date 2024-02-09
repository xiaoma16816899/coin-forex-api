package utils

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R {
	result := []R{}

	for i, item := range collection {
		if r, ok := callback(item, i); ok {
			result = append(result, r)
		}
	}

	return result
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, callback func(item T, index int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = callback(item, i)
	}

	return result
}

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[V any](collection []V, callback func(item V, index int) bool) []V {
	result := []V{}

	for i, item := range collection {
		if callback(item, i) {
			result = append(result, item)
		}
	}

	return result
}

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func FilterTake[V any](collection []V, take uint, callback func(item V, index int) bool) ([]V, int) {
	result := []V{}
	var found uint = 0

	for i, item := range collection {
		if callback(item, i) {
			result = append(result, item)

			found += 1
			if take != 0 && found == take {
				return result, i
			}
		}
	}

	return result, 0
}

// Split array into smaller chunks of the given size, reminder keep in the the last array
func Chunks[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func Sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}