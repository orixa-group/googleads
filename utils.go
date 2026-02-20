package googleads

func String(s string) *string {
	return &s
}

func Int64(i int64) *int64 {
	return &i
}

func Bool(b bool) *bool {
	return &b
}

func Map[I, O any](slice []I, fn func(item I) O) []O {
	ss := make([]O, 0, len(slice))
	for _, item := range slice {
		ss = append(ss, fn(item))
	}

	return ss
}
