package utils

// FilterMapBy filters map of strings by predicate.
func FilterMapBy(data map[string]string, predicate func(string, string) bool) map[string]string {
	result := make(map[string]string)
	for key, value := range data {
		if predicate(key, value) {
			result[key] = value
		}
	}

	return result
}
