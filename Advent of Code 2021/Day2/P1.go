// package main
package main

// import block
import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// main function
func main() {
	// our sub's depth
	var depth = 0
	// our sub's distance
	var distance = 0
	// our list of commands
	var cmds []string

	// store the command file in cmdsFile
	cmdsFile, _ := os.Open("commands.txt")

	// store the commands in cmds
	scanner := bufio.NewScanner(cmdsFile)
	for scanner.Scan() {
		text := scanner.Text()
		cmds = append(cmds, text)
	}

	// loop over cmds
	for i := 0; i < len(cmds); i++ {
		// set the command to the first text part
		cmd := strings.Fields(cmds[i])[0]
		// set the value to the second int part
		val, _ := strconv.Atoi(strings.Fields(cmds[i])[1])

		// switch case for which command
		switch cmd {
		case "forward":
			distance += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}

	}

	// print depth * distance
	println(depth * distance)
}
