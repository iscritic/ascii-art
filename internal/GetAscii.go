package internal

import (
	"strings"
)

func GetAscii(text, data string) (string, int) {
	table := CreateMap(string(data))

	text = strings.ReplaceAll(text, "\\n", "\n")

	s := customSplit(text)

	var result, subresult string
	var maxlen int
	var flag bool

	for i := 0; i < len(s); i++ {

		subs := s[i]

		if subs == "\n" {
			if flag {
				flag = false
				continue
			}
			result += "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range subs {
				if art, ok := table[char]; ok {
					subresult += art[i]
				}
			}

			// Get len of longest string
			if maxlen < len(subresult) {
				maxlen = len(subresult)
			}

			result += subresult + "\n"
			subresult = ""

		}

		// checking nextword existing
		if i < len(s)-1 && len(s[i+1]) > 0 {
			flag = true
		}
	}

	return result, maxlen
}

func CreateMap(s string) map[rune][]string {
	s = strings.ReplaceAll(s, "\r", "")
	lines := strings.Split(s, "\n")

	table := make(map[rune][]string)

	var arr []string
	var char rune = 32

	for i := 1; i < len(lines); i++ {
		if len(arr) != 8 {
			arr = append(arr, lines[i])
		} else {

			table[char] = arr

			arr = []string{}
			char++

		}
	}

	return table
}

func customSplit(s string) []string {
	var result []string
	var word string

	for _, r := range s {
		if r == '\n' {
			if len(word) > 0 {
				result = append(result, word)
				word = ""
			}

			result = append(result, "\n")

		} else {
			word += string(r)
		}
	}

	if len(word) > 0 {
		result = append(result, word)
	}

	return result
}
