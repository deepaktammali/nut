package helpers

func FirstNonEmpty(val1, val2 string) string {
	if val1 != "" {
		return val1
	}

	return val2
}
