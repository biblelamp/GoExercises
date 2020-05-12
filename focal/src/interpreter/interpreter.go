package interpreter

import (
	"bufio"
	"calculations/calculate"
	"fmt"
	"messages"
	"os"
	"strings"
	"util"
)

func Run() {
	var line string
	println(messages.WELCOME)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print(messages.PROMPT)
		scanner.Scan()
		line = scanner.Text()
		if len(line) > 0 {
			if result := processLine(line); result < 0 {
				break
			}
		}
	}
}

func commandAsk() {

}

func commandSet() {

}

func commandType(line string) {
	parameters := util.SplitString(line[strings.Index(line, " ")+1:])
	for _, item := range parameters {
		if item[0:1] == "\"" {
			fmt.Print(item[1 : len(item)-1])
		} else if item[0:1] == "%" {
			// TODO set format
		} else if item[0:1] == "$" {
			// TODO print all vbariables
		} else if item == "!" {
			fmt.Print("\n")
		} else if item == "#" {
			fmt.Print("\r")
		} else if item == ":" {
			fmt.Print("\t")
		} else {
			result := calculate.Calculate(item)
			fmt.Print(result)
		}
	}
}

func processLine(line string) float32 {
	tokens := strings.Split(line, " ")
	switch strings.ToUpper(tokens[0]) {
	case messages.T, messages.TYPE:
		commandType(line)
	case messages.Q, messages.QUIT:
		return -1
	default:
		fmt.Printf(messages.COMMAND_NOT_RECOGNIZED, tokens[0])
	}
	return 0
}
