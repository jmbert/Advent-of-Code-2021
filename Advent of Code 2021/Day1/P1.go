package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var numIncrease = 0
	var depths []int

	depthFile, _ := os.Open("depths.txt")

	scanner := bufio.NewScanner(depthFile)
	for scanner.Scan() {
		text := scanner.Text()
		iText, _ := strconv.Atoi(text)
		depths = append(depths, iText)
	}

	for i := 0; i < len(depths); i++ {
		if i != 0 {
			if depths[i] > depths[i-1] {
				numIncrease++
			}
		}
	}

	print(numIncrease)
}
