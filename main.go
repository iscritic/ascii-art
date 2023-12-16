package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("At least 2 arguments are required")
	}
	arg := os.Args[1]

	if arg == "" {
		return
	}

	data, err := FontPicker()
	if err != nil {
		log.Fatalln("Invalid font")
	}

	output := GetAscii(arg, data)

	if isValidTerminal(output) {
		log.Fatalln("Invalid terminal size")
	}

	fmt.Print(output)
}

func CreateMap(s string) map[rune][]string {
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

const (
	StandardHash   = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	ShadowHash     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	ThinkertoyHash = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
)

func FontPicker() (string, error) {
	font := "standard"
	errf := errors.New("invalid font")

	if len(os.Args) == 3 {
		font = os.Args[2]
	}

	file, err := os.Open(font + ".txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := os.ReadFile(font + ".txt")
	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(data)
	generatedHash := fmt.Sprintf("%x", hasher.Sum(nil))

	switch font {
	case "standard":
		if generatedHash != StandardHash {
			return "", errf
		}
	case "shadow":
		if generatedHash != ShadowHash {
			return "", errf
		}
	case "thinkertoy":
		if generatedHash != ThinkertoyHash {
			return "", errf
		}
	default:
		return "", errf
	}

	return string(data), nil
}

func GetAscii(text, data string) string {
	table := CreateMap(string(data))

	text = strings.ReplaceAll(text, "\\n", "\n")

	s := customSplit(text)

	var result string

	for _, subs := range s {
		if subs == "\n" {
			result += "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range subs {
				if art, ok := table[char]; ok {
					result += art[i]
				}
			}
			result += "\n"
		}
	}

	return result
}

func isValidTerminal(s string) bool {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln("Error getting size of terminal:", err)
	}

	var rows, cols int
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)

	return len(s)/8 >= cols
}
