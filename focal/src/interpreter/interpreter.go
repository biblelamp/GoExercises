package interpreter

import (
	"fmt"
	"messages"
	"strings"
)

func Run() {
	var line string
	println(messages.WELCOME)
	for {
		print(messages.PROMPT)
		n, _ := fmt.Scanln(&line)
		if n > 0 {
			if result := processLine(line); result < 0 {
				break
			}
		}
	}
}

func processLine(line string) float32 {
	parts := strings.Split(line, " ")
	switch strings.ToUpper(parts[0]) {
	case messages.Q, messages.QUIT:
		return -1
	default:
		fmt.Printf(messages.COMMAND_NOT_RECOGNIZED, parts[0])
	}
	return 0
}
