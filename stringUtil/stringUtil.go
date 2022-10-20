package stringutil

func SubString(source string, startIndex int, length int) string {
	if source == "" || len(source) < startIndex+length {
		return source
	} else {
		arry := []rune(source)
		s := arry[startIndex:length]
		return string(s)
	}
}
