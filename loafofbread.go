package piscine

func LoafOfBread(str string) string {
	result := 0
	for _, r := range str {
		if r != ' ' {
			result = result + 1
		}
	}
	if result < 5 {
		return "\n"
	}

	runes := []rune(str)
	output := ""
	word := ""
	i := 0

	for i < len(runes) {
		if runes[i] == ' ' {
			i = i + 1
			continue
		}
		word += string(runes[i])
		if len(word) == 5 {
			output += word + " "
			word = ""
			i = i + 1
		}
		i = i + 1
	}

	if len(word) > 0 {
		output += word
	} else if len(output) > 0 && output[len(output)-1] == ' ' {
		output = output[:len(output)-1]
	}
	return output + "\n"
}
