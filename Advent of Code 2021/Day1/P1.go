//package main
package main

// import block
import (
	"bufio"
	"os"
	"strconv"
)

// main function
func main() {
	// number of increases in depths
	var numIncrease = 0

	// the depths
	var depths []int

	// assign depthfile the file containing our depths
	depthFile, _ := os.Open("depths.txt")

	// read through the file line by line, and assign the depths into our slice
	scanner := bufio.NewScanner(depthFile)
	for scanner.Scan() {
		text := scanner.Text()
		iText, _ := strconv.Atoi(text)
		depths = append(depths, iText)
	}

	// check if it increases, and if it does increment numIncrease
	for i := 0; i < len(depths); i++ {
		if i != 0 {
			if depths[i] > depths[i-1] {
				numIncrease++
			}
		}
	}

	// print the number of times it increased
	print(numIncrease)
}
