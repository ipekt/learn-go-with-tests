package iteration

func Repeat(character string, num int) string {
	// Use `var` to declare a variable without initializing
	var repeated string
	for i := 0; i < num; i++ {
		repeated = repeated + character
	}
	return repeated
}
