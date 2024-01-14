package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GetTerminalSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln("Error getting size of terminal:", err)
	}

	var rows, cols int
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)

	return rows, cols
}

func IsValidTerminal(maxlen, cols int) bool {
	return maxlen > cols
}
