package justify

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func AsciiJustify() {
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

	output, maxlen := GetAscii(arg, data)

	PrintAsciiArt(output)
	fmt.Println(maxlen)

	// cols, ok := isValidTerminal(maxlen)
	// if ok {
	// 	log.Fatalln("Invalid terminal size")
	// }
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

func GetAscii(text, data string) ([][]string, int) {
	table := CreateMap(data)

	text = strings.ReplaceAll(text, "\\n", "\n")
	s := customSplit(text)

	var superresult [][]string
	var maxlen int

	for _, subs := range s {
		if subs == "\n" {
			superresult = append(superresult, []string{"\n"})
			continue
		}

		words := strings.Fields(subs)
		for _, word := range words {
			var wordArt []string
			for i := 0; i < 8; i++ {
				var line string
				for _, char := range word {
					if art, ok := table[char]; ok {
						line += art[i]
					}
				}
				wordArt = append(wordArt, line)
				if len(line) > maxlen {
					maxlen = len(line)
				}
			}
			superresult = append(superresult, wordArt)
		}
	}
	return superresult, maxlen
}

func isValidTerminal(maxlen int) (int, bool) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln("Error getting size of terminal:", err)
	}

	var rows, cols int
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)

	return cols, maxlen > cols
}

func PrintAsciiArt(superresult [][]string) {
	for _, arr := range superresult {
		if arr[0] == "\n" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			fmt.Println(arr[i])
		}
	}
}
