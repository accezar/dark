package converter

func stringInSlice(search string, haystack []string) bool {
	for _, item := range haystack {
		if item == search {
			return true
		}
	}

	return false
}

func intPtr(input int) *int {
	return &input
}

func float64Ptr(input float64) *float64 {
	return &input
}
