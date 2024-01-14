package main

import (
	"fmt"
	"log"
	"os"

	. "ascii/internal"
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

	output, maxlen := GetAscii(arg, data)

	_, cols := GetTerminalSize()

	if IsValidTerminal(maxlen, cols) {
		log.Fatalln("Invalid terminal size")
	}

	fmt.Print(output)
}
