package iter

func Repeat(candidate string, iter int) string {
	var repeated string
	for i := 0; i < iter; i++ {
		repeated += candidate
	}
	return repeated
}
