package googleads

import "strconv"

func String(s string) *string {
	return &s
}

func Int64(i int64) *int64 {
	return &i
}

func Int32(i int32) *int32 {
	return &i
}

func Float64(f float64) *float64 {
	return &f
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

type tempIdGenerator func() string

func newTempIdGenerator() tempIdGenerator {
	i := 0
	return func() string {
		i--
		return strconv.Itoa(i)
	}
}
