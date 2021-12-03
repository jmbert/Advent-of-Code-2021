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
	// our oxygen Generator status
	var oxGen int64
	// CO2 scrubber status
	var CO2Scrubber int64

	// diagnostics
	var diagnostics []string
	// our oxygen diagnostics (filtered)
	var filteredOxDiagnostics []string
	// our CO2 scrubber diagnostics (filtered)
	var filteredCODiagnostics []string

	// store the diagnostic file in diagFile
	diagFile, _ := os.Open("diagnostics.txt")

	// store the diagnostics in diagnostics
	scanner := bufio.NewScanner(diagFile)
	for scanner.Scan() {
		text := scanner.Text()
		diagnostics = append(diagnostics, text)
	}

	// set the filtered diagnostics to diagnostics
	filteredOxDiagnostics = diagnostics
	filteredCODiagnostics = diagnostics

	// loop over each bit
	for i := 0; i < len(diagnostics[0]); i++ {
		// check if there is only one left
		if len(filteredOxDiagnostics) > 1 {
			// ReCalculate the most common values
			Counts := CalculateCommons(filteredOxDiagnostics)
			// Debug print
			println("next bit")
			// the indexes of the diagnostics to be filtered
			var toBeFiltered []int
			// loop over all the diagnostics left
			for j := 0; j < len(filteredOxDiagnostics); j++ {
				// check if the most common bit here is a 1 or 0
				if Counts[i] == 1 {
					// set all the 0s to be filtered out
					if filteredOxDiagnostics[j][i] == 48 {
						toBeFiltered = append(toBeFiltered, j)
					}
				} else if Counts[i] == 0 {
					// otherwise, set all 1s to be filtered out
					if filteredOxDiagnostics[j][i] == 49 {
						toBeFiltered = append(toBeFiltered, j)
					}
				}
			}
			// debug print
			println(len(toBeFiltered))
			// loop over the indexes to be filtered
			for j := 0; j < len(toBeFiltered); j++ {
				// debug print
				println(toBeFiltered[j])
				// remove them from the filtered diagnostics (backwards, so indexes are preserved)
				filteredOxDiagnostics = RemoveIndex(filteredOxDiagnostics, toBeFiltered[len(toBeFiltered)-j-1])
			}

		} else {
			// otherwise, break out of the loop
			break
		}

	}

	// !!CO2 section is broken, suspected bit filter not working, as it filters every number out !! loop over each bit
	for i := 0; i < len(diagnostics[0]); i++ {
		// debug print
		//println(filteredCODiagnostics[0])
		// check if this is the last number
		if len(filteredCODiagnostics) > 1 {
			// recalculate the most common values
			Counts := CalculateCommons(filteredCODiagnostics)
			// debug print
			println("next bit")
			// indexes to be filtered
			var toBeCOFiltered []int
			// loop over the numbers
			for j := 0; j < len(filteredCODiagnostics); j++ {
				// check if the most common bit is a 1
				if Counts[i] == 1 {
					// if it is, filter out the 1s
					if filteredCODiagnostics[j][i] == 49 {
						toBeCOFiltered = append(toBeCOFiltered, j)
					}
				} else if Counts[i] == 0 {
					// otherwise, filter out the 0s
					if filteredCODiagnostics[j][i] == 48 {
						toBeCOFiltered = append(toBeCOFiltered, j)
					}
				}
			}
			// debug print
			println(len(toBeCOFiltered))
			// loop of the indexes to be filtered
			for j := 0; j < len(toBeCOFiltered); j++ {
				// debug print
				println(toBeCOFiltered[j])
				// remove those indexes, backwards to preserve indexes
				filteredCODiagnostics = RemoveIndex(filteredCODiagnostics, toBeCOFiltered[len(toBeCOFiltered)-j-1])
			}
		} else {
			// otherwise, break out
			break
		}
	}

	// debug prints
	println(len(filteredOxDiagnostics))
	println(len(filteredCODiagnostics))
	println(filteredOxDiagnostics[0])
	println(strconv.ParseInt(filteredOxDiagnostics[0], 2, 64))
	println(filteredCODiagnostics[0])
	println(strconv.ParseInt(filteredCODiagnostics[0], 2, 64))

	// Parse the binary to ints
	oxGen, _ = strconv.ParseInt(filteredOxDiagnostics[0], 2, 64)
	CO2Scrubber, _ = strconv.ParseInt(filteredCODiagnostics[0], 2, 64)

	// print oxygen Generator * the CO2 scrubber
	println(oxGen * CO2Scrubber)

}

// Remove Index function
func RemoveIndex(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// Calculate Commons function
func CalculateCommons(diagnostics []string) [12]int {
	var Counts [12]int
	for i := 0; i < len(diagnostics); i++ {
		for j := 0; j < len(diagnostics[i]); j++ {
			if diagnostics[i][j] == 49 {
				Counts[j]++
			}
		}
	}
	for i := 0; i < len(Counts); i++ {
		if Counts[i] >= len(diagnostics)/2 {
			Counts[i] = 1
		} else {
			Counts[i] = 0
		}
	}

	return Counts
}
