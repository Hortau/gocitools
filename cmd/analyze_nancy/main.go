package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"analyze_nancy/model"

	"github.com/sonatype-nexus-community/go-sona-types/ossindex/types"
)

func formatVulnerability(vulnerableDep types.Coordinate) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Package %s is vulnerable:\n", vulnerableDep.Coordinates)

	for _, v := range vulnerableDep.Vulnerabilities {
		fmt.Fprintf(&b, "CvssScore: %s/10\nDescription: %s\n\n", v.CvssScore.String(), v.Description)
	}
	return b.String()
}

func hasHighVulnerabilities(dep types.Coordinate, minCVSS float64) bool {
	for _, v := range dep.Vulnerabilities {
		score, _ := v.CvssScore.Float64()
		if score > minCVSS {
			return true
		}
	}
	return false
}

func parseNancy(file string, minCVSS float64) ([]types.Coordinate, error) {
	jsonFile, err := os.Open(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()

	var data model.ReportInfo
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&data); err != nil {
		panic(err)
	}

	fmt.Printf("Audited %d package(s) by Nancy.\n", data.NumAudited)
	fmt.Printf("Excluded %d package(s) by Nancy.\n", data.NumExclusions)
	fmt.Printf("Found %d potentially vulnerable package(s) by Nancy.\n", data.NumVulnerable)
	fmt.Printf("Vulnerabilities will be filtered out with a minimum threshold of %.1f\n\n", minCVSS)

	var vulnerableDeps []types.Coordinate
	for _, dep := range data.Vulnerable {
		fmt.Printf("Analyzing vulnerabilities for package %s\n", dep.Coordinates)
		if hasHighVulnerabilities(dep, minCVSS) {
			fmt.Printf("Vulnerabilities found for package %s\n", dep.Coordinates)
			vulnerableDeps = append(vulnerableDeps, dep)
		}
	}
	return vulnerableDeps, nil
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("usage: analyze_nancy results.json 0")
		os.Exit(1)
	}
	jsonFile := os.Args[1]

	var minCVSS float64
	var err error
	if len(args) > 2 {
		minCVSS, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			panic(err)
		}
	}

	if minCVSS > 10 || minCVSS < 0 {
		fmt.Println("CVSS threshold must be between 0 and 10")
		os.Exit(1)
	}

	result, err := parseNancy(jsonFile, minCVSS)
	if err != nil {
		panic(err)
	}

	if len(result) > 0 {
		fmt.Println("--------------------- BEGIN VULNERABILITIES FOUND ---------------------")
		for _, v := range result {
			fmt.Print(formatVulnerability(v))
		}
		fmt.Println("---------------------- END VULNERABILITIES FOUND ----------------------")
		os.Exit(1)
	} else {
		fmt.Println("No vulnerabilities found")
	}
}
