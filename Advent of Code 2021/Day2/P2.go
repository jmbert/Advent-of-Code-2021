package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var aim = 0
	var depth = 0
	var distance = 0
	var cmds []string

	cmdsFile, _ := os.Open("commands.txt")

	scanner := bufio.NewScanner(cmdsFile)
	for scanner.Scan() {
		text := scanner.Text()
		cmds = append(cmds, text)
	}

	for i := 0; i < len(cmds); i++ {
		cmd := strings.Fields(cmds[i])[0]
		val, _ := strconv.Atoi(strings.Fields(cmds[i])[1])

		switch cmd {
		case "forward":
			distance += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}

	}
	println(depth * distance)
}
