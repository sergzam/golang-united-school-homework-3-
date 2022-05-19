package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range keys {
		if v, ok := input[key]; ok {
			result = append(result, v)
		}
	}

	return
}
