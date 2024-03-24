package main

import (
	"fmt"
	"os"
)

func printUsageHelp() {
	fmt.Println("Usage: go run ./cmd <format> <filename>")
	fmt.Println("Arguments <format> and <filename> are optional.")
	fmt.Println(
		"When both arguments are omitted, program will generate",
		"files for all supported formats with default generated names.",
	)
	fmt.Println(
		"When <filename> is omitted, program will generate",
		"a file for provided <format> with a default generated name.",
	)
	fmt.Println("Supported formats are: json-array, json-map, csv")
	fmt.Println("Examples: \n\tgo run ./cmd")
	fmt.Println("\tgo run ./cmd json-array")
	fmt.Println("\tgo run ./cmd json-map currencies-map.json")
	fmt.Println("\tgo run ./cmd csv currencies.csv")
}

func main() {
	args := os.Args

	if len(args) <= 1 {
		currencies, err := FetchCurrencies()
		if err != nil {
			fmt.Printf("Failed to fetch currencies: %s\n", err)
			os.Exit(1)
		}

		err = WriteCurrenciesToDefaultFilesWithAllFormats(currencies)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		return
	}

	if len(args) > 3 || (args[1] == "-h") || (args[1] == "--help") {
		printUsageHelp()
		os.Exit(0)
	}

	format := args[1]
	var fileName string

	if len(args) == 3 {
		fileName = args[2]
	} else {
		fileName = GetDefaultFileNameFromFormat(format)
	}

	if !IsSupportedFormat(format) {
		fmt.Printf("Invalid format: %s\n", format)
		printUsageHelp()
		os.Exit(1)
	}

	err := ValidateFileName(fileName, true)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	currencies, err := FetchCurrencies()
	if err != nil {
		fmt.Printf("Failed to fetch currencies: %s\n", err)
		os.Exit(1)
	}

	err = WriteCurrenciesToFileWithFormat(currencies, format, fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
