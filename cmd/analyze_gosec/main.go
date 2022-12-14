package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"analyze_gosec/model"
)

func parseGosec(file string) ([]*model.Issue, error) {
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

	issues := data.Issues
	fmt.Printf("Found a total %v issues\n", len(issues))
	return issues, nil
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("usage: analyze_gosec results.json")
		os.Exit(1)
	}

	issues, err := parseGosec(os.Args[1])
	if err != nil {
		panic(err)
	}

	if len(issues) > 0 {
		fmt.Println("--------------------- BEGIN VULNERABILITIES FOUND ---------------------")
		b, _ := json.MarshalIndent(issues, "", "  ")
		fmt.Println(string(b))
		fmt.Println("---------------------- END VULNERABILITIES FOUND ----------------------")
		os.Exit(1)
	}
}
