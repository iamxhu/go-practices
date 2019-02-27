package util

func Contains(s []int, v int) bool {
	for _, value := range s {
		if value == v {
			return true
		}
	}

	return false
}
