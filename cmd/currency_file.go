package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func ValidateFileName(fileName string, promptForOverwrite bool) error {
	if fileName == "" {
		return fmt.Errorf("invalid file name: empty")
	}

	fileStat, err := os.Stat(fileName)

	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf(
				"failed to check if file %s exists: %s", fileName, err,
			)
		}

		fmt.Printf("File %s does not exist. Creating...\n", fileName)

		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %s", fileName, err)
		}

		file.Close()

		fmt.Printf("File %s created.\n", fileName)

		return nil
	}

	if fileStat.IsDir() {
		return fmt.Errorf("invalid file name: %s is a directory", fileName)
	}

	if promptForOverwrite {
		fmt.Printf(
			"File %s already exists. Do you want to overwrite it? (y/n): ",
			fileName,
		)

		var overwrite string
		fmt.Scanln(&overwrite)

		if overwrite != "y" && overwrite != "Y" {
			return fmt.Errorf("exiting. No changes were made")
		}

		fmt.Printf("File %s will be overwritten.\n", fileName)

		return nil
	}

	return fmt.Errorf("file %s already exists", fileName)
}

func WriteCurrenciesToFileWithFormat(
	currencies Currencies,
	format string,
	fileName string,
) error {
	fmt.Printf(
		"Writing currencies to file %s in %s format...\n",
		fileName, format,
	)

	fileContent, err := currencies.ToBytes(format)
	if err != nil {
		return fmt.Errorf(
			"failed to convert currencies to %s: %s", format, err,
		)
	}

	err = os.WriteFile(fileName, fileContent, 0644)
	if err != nil {
		return fmt.Errorf(
			"failed to write currencies to file %s: %s", fileName, err,
		)
	}

	fmt.Printf("Currencies written to file %s.\n", fileName)

	return nil
}

func WriteCurrenciesToDefaultFilesWithAllFormats(currencies Currencies) error {
	err := WriteCurrenciesToFileWithFormat(
		currencies, JsonArray, GetDefaultFileNameFromFormat(JsonArray),
	)
	if err != nil {
		return err
	}

	err = WriteCurrenciesToFileWithFormat(
		currencies, JsonMap, GetDefaultFileNameFromFormat(JsonMap),
	)
	if err != nil {
		return err
	}

	err = WriteCurrenciesToFileWithFormat(
		currencies, Csv, GetDefaultFileNameFromFormat(Csv),
	)
	if err != nil {
		return err
	}

	return nil
}

func GetDefaultFileNameFromFormat(format string) string {
	timestamp := time.Now().Format("2006-01-02")

	switch format {
	case JsonArray:
		return fmt.Sprintf("nbg-currencies-array-%s.json", timestamp)
	case JsonMap:
		return fmt.Sprintf("nbg-currencies-map-%s.json", timestamp)
	case Csv:
		return fmt.Sprintf("nbg-currencies-%s.csv", timestamp)
	default:
		return ""
	}
}
