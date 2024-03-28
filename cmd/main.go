package main

import (
	"fmt"
	"os"

	"nbg-currencies/internal/currency"
	"nbg-currencies/internal/multiselect"
)

func hangUntilEnter() {
	fmt.Print("\nPress enter to exit...\n")
	fmt.Scanln()
}

func main() {
	selectedFormatIndexes, submittedOnQuit, err := multiselect.PromptMultiselect(
		"Which file format(s) do you want to generate for currencies?",
		currency.SupportedFormats,
	)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		hangUntilEnter()
		os.Exit(1)
	}

	if selectedFormatIndexes.IsEmpty() {
		fmt.Println("No file formats selected.")

		if submittedOnQuit {
			hangUntilEnter()
		}

		return
	}

	selectedFormats := make([]string, 0, len(selectedFormatIndexes))
	for i := range selectedFormatIndexes {
		selectedFormats = append(selectedFormats, currency.SupportedFormats[i])
	}

	fmt.Print("Selected formats: ")
	for i, format := range selectedFormats {
		fmt.Print(format)
		if i != len(selectedFormats)-1 {
			fmt.Print(", ")
		} else {
			fmt.Print(".\n\n")
		}
	}

	currencies, elapsedMs, err := currency.FetchCurrencies()
	if err != nil {
		fmt.Printf("Failed to fetch currencies: %s\n", err)
		hangUntilEnter()
		os.Exit(1)
	}

	fmt.Printf(
		"Successfully fetched currencies in %d milliseconds.\n", elapsedMs,
	)

	for _, format := range selectedFormats {
		fileName, err := currency.GenerateDefaultFileNameFromFormat(format)
		if err != nil {
			fmt.Printf("Couldn't generate file name: %s\n", err)
			continue
		}

		err = currency.WriteCurrenciesToFile(currencies, format, fileName)
		if err != nil {
			fmt.Printf("Couldn't write currencies to file: %s\n", err)
			continue
		}

		fmt.Printf(
			"Successfully written currencies to file %s in %s format.\n",
			fileName, format,
		)
	}

	hangUntilEnter()
}
