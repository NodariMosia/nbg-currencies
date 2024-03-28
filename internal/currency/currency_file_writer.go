package currency

import (
	"fmt"
	"os"
	"time"
)

func WriteCurrenciesToFile(
	currencies Currencies,
	format string,
	fileName string,
) error {
	if !IsSupportedFormat(format) {
		return fmt.Errorf("unsupported format: %s", format)
	}

	fileContent, err := currencies.toBytes(format)
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

	return nil
}

func GenerateDefaultFileNameFromFormat(format string) (string, error) {
	timestamp := time.Now().Format("2006-01-02")

	switch format {
	case JsonArrayFormat:
		return fmt.Sprintf("nbg-currencies-array-%s.json", timestamp), nil
	case JsonMapFormat:
		return fmt.Sprintf("nbg-currencies-map-%s.json", timestamp), nil
	case CsvFormat:
		return fmt.Sprintf("nbg-currencies-%s.csv", timestamp), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}
