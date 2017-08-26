package stringutil


func reverseTwo(s string) string  {

	var out string = ""
	for i := len(s) - 1; i >= 0; i-- {
		//fmt.Println(string(s[i]))
		out += string(s[i])
	}

	return out
}
