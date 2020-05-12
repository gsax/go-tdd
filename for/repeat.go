package iteration

const repeateCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeateCount; i++ {
		repeated += character
	}
	return repeated
}
