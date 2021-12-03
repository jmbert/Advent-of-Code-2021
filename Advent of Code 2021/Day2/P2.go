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
	// our sub's aim
	var aim = 0
	// our sub's depth
	var depth = 0
	// our sub's distance
	var distance = 0
	var cmds []string

	// store command file in cmdsFile
	cmdsFile, _ := os.Open("commands.txt")

	// store commands in cmds
	scanner := bufio.NewScanner(cmdsFile)
	for scanner.Scan() {
		text := scanner.Text()
		cmds = append(cmds, text)
	}

	// loop over cmds
	for i := 0; i < len(cmds); i++ {
		// store the command in cmd
		cmd := strings.Fields(cmds[i])[0]
		// store the value in val
		val, _ := strconv.Atoi(strings.Fields(cmds[i])[1])

		// switch over cmd
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
	// print depth * distance
	println(depth * distance)
}
