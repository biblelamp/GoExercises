package interpreter

import (
	"bufio"
	"calculations/calculate"
	"fmt"
	"messages"
	"os"
	"sort"
	"strconv"
	"strings"
	"util"
)

var formatNumber string
var variables = make(map[string]float64)
var programLines = make(map[float64]string)

func Run() {
	var line string
	formatNumber = messages.DEFAULT_FORMAT_NUMBER
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
	// TODO implement ASK command
}

func commandSet(line string) {
	parameters := strings.Split(line[strings.Index(line, " ")+1:], "=")
	result := calculate.Calculate(parameters[1], variables)
	variables[parameters[0]] = result
}

func commandType(line string) {
	parameters := util.SplitString(line[strings.Index(line, " ")+1:])
	for _, item := range parameters {
		if item[0:1] == "\"" {
			fmt.Print(item[1 : len(item)-1])
		} else if item[0:1] == "%" {
			// TODO set format
		} else if item[0:1] == "$" {
			for key, value := range variables {
				fmt.Printf("%s()="+formatNumber+"\n", key, value)
			}
		} else if item == "!" {
			fmt.Print("\n")
		} else if item == "#" {
			fmt.Print("\r")
		} else if item == ":" {
			fmt.Print("\t")
		} else {
			result := calculate.Calculate(item, variables)
			fmt.Printf(formatNumber, result)
		}
	}
}

func commandWrite(tokens []string) {
	var keys []float64
	for k := range programLines {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	for _, key := range keys {
		fmt.Printf("%05.2f %s\n", key, programLines[key])
	}
}

func processLine(line string) float32 {
	tokens := strings.Split(line, " ")
	if util.IsValidLineNumber(tokens[0]) {
		if value, err := strconv.ParseFloat(tokens[0], 32); err == nil {
			if len(tokens) > 1 {
				// to add line
				programLines[value] = line[strings.Index(line, " ")+1:]
			} else {
				// erase line
				delete(programLines, value)
			}
		}
	} else {
		switch strings.ToUpper(tokens[0]) {
		case messages.S, messages.SET:
			commandSet(line)
		case messages.T, messages.TYPE:
			commandType(line)
		case messages.W, messages.WRITE:
			commandWrite(tokens)
		case messages.Q, messages.QUIT:
			return -1
		default:
			fmt.Printf(messages.COMMAND_NOT_RECOGNIZED, tokens[0])
		}
	}
	return 0
}
