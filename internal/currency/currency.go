package currency

import (
	"encoding/json"
	"fmt"
)

const (
	JsonArrayFormat string = "json-array"
	JsonMapFormat   string = "json-map"
	CsvFormat       string = "csv"
)

var SupportedFormats = []string{JsonArrayFormat, JsonMapFormat, CsvFormat}

func IsSupportedFormat(format string) bool {
	switch format {
	case JsonArrayFormat, JsonMapFormat, CsvFormat:
		return true
	default:
		return false
	}
}

type Currency struct {
	Index          string  `json:"index"`
	ConversionRate float64 `json:"conversionRate"`
}

type Currencies []Currency

// Returns json array representation of currencies with format:
//
//	[{"index":"USD","conversionRate":2.0},{"index":"EUR","conversionRate":3.0}]
func (currencies Currencies) ToJsonArray() ([]byte, error) {
	return json.Marshal(currencies)
}

// Returns json map representation of currencies with format:
//
//	{"USD":2.0,"EUR":3.0}
func (currencies Currencies) ToJsonMap() ([]byte, error) {
	jsonMap := make(map[string]float64)

	for _, currency := range currencies {
		jsonMap[currency.Index] = currency.ConversionRate
	}

	return json.Marshal(jsonMap)
}

// Returns csv representation of currencies with format:
//
//	index,conversionRate
//	USD,2.0
//	EUR,3.0
func (currencies Currencies) ToCSV() ([]byte, error) {
	csv := "index,conversionRate\n"

	for _, currency := range currencies {
		csv += fmt.Sprintf("%s,%f\n", currency.Index, currency.ConversionRate)
	}

	return []byte(csv), nil
}

// Returns byte slice representation of currencies with given format.
// Supported formats: json-array, json-map, csv
func (currencies Currencies) toBytes(format string) ([]byte, error) {
	switch format {
	case JsonArrayFormat:
		return currencies.ToJsonArray()
	case JsonMapFormat:
		return currencies.ToJsonMap()
	case CsvFormat:
		return currencies.ToCSV()
	default:
		return nil, fmt.Errorf("invalid format: %s", format)
	}
}
