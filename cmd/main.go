package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func printFormatHelp() {
	fmt.Println("Format can be one of: json-array, json-map, csv")
}

func printUsageHelp() {
	fmt.Println("Usage: go run ./cmd <format> <file name>")
	printFormatHelp()
	fmt.Println("Example: go run ./cmd json-array currencies.json")
}

func main() {
	args := os.Args

	if (len(args) == 1) || (args[1] == "-h") || (args[1] == "--help") {
		printUsageHelp()
		os.Exit(0)
	}

	if len(args) != 3 {
		printUsageHelp()
		os.Exit(1)
	}

	format, fileName := args[1], args[2]

	if !IsSupportedFormat(format) {
		fmt.Printf("Invalid format: %s\n", format)
		printFormatHelp()
		os.Exit(1)
	}

	if fileName == "" {
		fmt.Println("Invalid file name: empty")
		os.Exit(1)
	}

	fileStat, err := os.Stat(fileName)

	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Failed to check if file %s exists: %s\n", fileName, err)
			os.Exit(1)
		}

		fmt.Printf("File %s does not exist. Creating...\n", fileName)

		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Failed to create file %s: %s\n", fileName, err)
			os.Exit(1)
		}

		file.Close()

		fmt.Printf("File %s created.\n", fileName)
	} else if fileStat.IsDir() {
		fmt.Printf("Invalid file name: %s is a directory\n", fileName)
		os.Exit(1)
	} else {
		fmt.Printf(
			"File %s already exists. Do you want to overwrite it? (y/n): ",
			fileName,
		)

		var overwrite string
		fmt.Scanln(&overwrite)

		if overwrite != "y" && overwrite != "Y" {
			fmt.Printf("Exiting. No changes were made.\n")
			os.Exit(0)
		}

		fmt.Printf("File %s will be overwritten.\n", fileName)
	}

	fmt.Println("Fetching currencies...")
	fetchStart := time.Now()

	currencies, err := FetchCurrencies()
	if err != nil {
		fmt.Printf("Failed to fetch currencies: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf(
		"Currencies fetched in %d milliseconds.\n",
		time.Since(fetchStart)/time.Millisecond,
	)

	fmt.Printf(
		"Writing currencies to file %s in %s format...\n",
		fileName, format,
	)

	fileContent, err := currencies.ToBytes(format)
	if err != nil {
		fmt.Printf("Failed to convert currencies to %s: %s\n", format, err)
		os.Exit(1)
	}

	err = os.WriteFile(fileName, fileContent, 0644)
	if err != nil {
		fmt.Printf("Failed to write currencies to file %s: %s\n", fileName, err)
		os.Exit(1)
	}

	fmt.Printf("Currencies written to file %s.\n", fileName)
}
