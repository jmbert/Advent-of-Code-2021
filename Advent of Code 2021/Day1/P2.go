package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var numIncrease = 0
	var depths []int
	var depthSums []int

	depthFile, _ := os.Open("depths.txt")

	scanner := bufio.NewScanner(depthFile)
	for scanner.Scan() {
		text := scanner.Text()
		iText, _ := strconv.Atoi(text)
		depths = append(depths, iText)
	}

	for i := 0; i < len(depths); i++ {
		if i+3 <= len(depths) {
			depthSums = append(depthSums, (depths[i] + depths[i+1] + depths[i+2]))
		}
	}

	for i := 0; i < len(depthSums); i++ {
		if i != 0 {
			if depthSums[i] > depthSums[i-1] {
				numIncrease++
			}
		}
	}

	print(numIncrease)
}
