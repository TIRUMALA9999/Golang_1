package even_odd

func IsEven(a int) string {
	if a%2 == 0 { // ✅ boolean expression
		return "even"
	}
	return "odd"
}
