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
	// store our gamma value in binary
	var gammaBin = make([]rune, 12)

	// store epsilon in binary
	var epsilonBin = make([]rune, 12)

	// diagnostics
	var diagnostics []string

	// most common bit for each bit place
	var Counts [12]int

	// store diagnostics file in diagFile
	diagFile, _ := os.Open("diagnostics.txt")

	// store the diagnostics in diagnostics
	scanner := bufio.NewScanner(diagFile)
	for scanner.Scan() {
		text := scanner.Text()
		diagnostics = append(diagnostics, text)
	}

	// loop over diagnostics
	for i := 0; i < len(diagnostics); i++ {
		// loop over each bit
		for j := 0; j < len(diagnostics[i]); j++ {
			// check if the bit is a 1 or not
			if []byte(diagnostics[i])[j] == 49 {
				// if it is, increment Counts at that bit place
				Counts[j]++
			}
		}
	}

	// loop over the bits in count
	for i := 0; i < len(Counts); i++ {
		// check if 1 is the most common bit
		if Counts[i] > len(diagnostics)/2 {
			// if it is, epsilon is a 0 at that bit
			epsilonBin[i] = '0'
			// and gamma is a 1
			gammaBin[i] = '1'
		} else {
			// otherwise, epsilon is a 1
			epsilonBin[i] = '1'
			// and gamma a 0
			gammaBin[i] = '0'
		}
	}

	// convert the binary rune slice into a string
	gammaString := string(gammaBin)
	epsilonString := string(epsilonBin)

	// debug prints
	println(gammaString)
	println(epsilonString)

	// Parse the binary string into an integer
	gamma, _ := strconv.ParseInt(string(gammaBin), 2, 64)
	epsilon, _ := strconv.ParseInt(string(epsilonBin), 2, 64)

	// Debug prints
	println(gamma)
	println(epsilon)
	println(Counts[1])

	// print gamma times epsilon
	println(gamma * epsilon)

}
