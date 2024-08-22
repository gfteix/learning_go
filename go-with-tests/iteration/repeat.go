package iteration

func Repeat(character string, number int) string {
	var repeated string
	for i := 0; i < number; i++ {
		repeated += character
	}
	return repeated
}
