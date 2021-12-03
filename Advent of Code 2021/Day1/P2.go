// package main
package main

// import block
import (
	"bufio"
	"os"
	"strconv"
)

// main function
func main() {
	// number of increases
	var numIncrease = 0

	// the depths
	var depths []int

	// the 3-window sums of the depths
	var depthSums []int

	// assign depthfile the file containing our depths
	depthFile, _ := os.Open("depths.txt")

	// read through the file line by line, and assign the depths into our slice
	scanner := bufio.NewScanner(depthFile)
	for scanner.Scan() {
		text := scanner.Text()
		iText, _ := strconv.Atoi(text)
		depths = append(depths, iText)
	}

	// loop over every depth
	for i := 0; i < len(depths); i++ {
		// check if we can add 3 more numbers
		if i+3 <= len(depths) {
			// add 3 numbers, then append that sum to depthSums
			depthSums = append(depthSums, (depths[i] + depths[i+1] + depths[i+2]))
		}
	}

	// loop over our sums, and check if they increase
	for i := 0; i < len(depthSums); i++ {
		if i != 0 {
			if depthSums[i] > depthSums[i-1] {
				numIncrease++
			}
		}
	}

	// print the number of times it increases
	print(numIncrease)
}
