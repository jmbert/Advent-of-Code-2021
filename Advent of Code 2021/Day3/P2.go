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
	diagFile, _ := os.Open("t_diagnostics.txt")

	// store the diagnostics in diagnostics
	scanner := bufio.NewScanner(diagFile)
	for scanner.Scan() {
		text := scanner.Text()
		diagnostics = append(diagnostics, text)
	}

	// set the filtered diagnostics to diagnostics
	filteredCODiagnostics = diagnostics
	filteredOxDiagnostics = diagnostics

	var tmpFilteredOx []string
	var tmpFilteredCO []string

	for bit := 0; bit < len(diagnostics[0]); bit++ {
		if len(filteredOxDiagnostics) > 1 {
			tmpFilteredOx = nil
			Counts := CalculateCommons(filteredOxDiagnostics)
			println()
			for k := 0; k < len(Counts); k++ {
				print(Counts[k])
			}
			println()
			println("The most common bit is " + strconv.Itoa(Counts[bit]))
			if Counts[bit] == 1 {
				for bin := 0; bin < len(filteredOxDiagnostics); bin++ {
					if filteredOxDiagnostics[bin][bit] == '1' {
						println("Keep " + filteredOxDiagnostics[bin])
						tmpFilteredOx = append(tmpFilteredOx, filteredOxDiagnostics[bin])
					}
				}
			} else if Counts[bit] == 0 {
				for bin := 0; bin < len(filteredOxDiagnostics); bin++ {
					if filteredOxDiagnostics[bin][bit] == '0' {
						println("Keep " + filteredOxDiagnostics[bin])
						tmpFilteredOx = append(tmpFilteredOx, filteredOxDiagnostics[bin])
					}
				}
			}

			println()

			for j := 0; j < len(tmpFilteredOx); j++ {
				println("Kept " + tmpFilteredOx[j])
			}

			filteredOxDiagnostics = tmpFilteredOx
		}
	}

	for bit := 0; bit < len(diagnostics[0]); bit++ {
		if len(filteredCODiagnostics) > 1 {
			tmpFilteredCO = nil
			Counts := CalculateCommons(filteredCODiagnostics)
			println()
			for k := 0; k < len(Counts); k++ {
				print(Counts[k])
			}
			println()
			println("The most common bit is " + strconv.Itoa(Counts[bit]))
			if Counts[bit] == 0 {
				for bin := 0; bin < len(filteredCODiagnostics); bin++ {
					if filteredCODiagnostics[bin][bit] == '1' {
						println("Keep " + filteredCODiagnostics[bin])
						tmpFilteredCO = append(tmpFilteredCO, filteredCODiagnostics[bin])
					}
				}
			} else if Counts[bit] == 1 {
				for bin := 0; bin < len(filteredCODiagnostics); bin++ {
					if filteredCODiagnostics[bin][bit] == '0' {
						println("Keep " + filteredCODiagnostics[bin])
						tmpFilteredCO = append(tmpFilteredCO, filteredCODiagnostics[bin])
					}
				}
			}

			println()

			for j := 0; j < len(tmpFilteredCO); j++ {
				println("Kept " + tmpFilteredCO[j])
			}

			filteredCODiagnostics = tmpFilteredCO
		}
	}

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
func CalculateCommons(s []string) [12]int {
	var Counts [12]int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == 49 {
				Counts[j]++
			}
		}
	}
	for i := 0; i < len(Counts); i++ {
		if Counts[i] >= (len(s) - Counts[i]) {
			Counts[i] = 1
		} else {
			Counts[i] = 0
		}
	}

	return Counts
}
