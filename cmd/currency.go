package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	JsonArray string = "json-array"
	JsonMap   string = "json-map"
	Csv       string = "csv"
)

func IsSupportedFormat(format string) bool {
	switch format {
	case JsonArray, JsonMap, Csv:
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

// Returns bytes representation of currencies with given format.
// Supported formats: json-array, json-map, csv
func (currencies Currencies) ToBytes(format string) ([]byte, error) {
	switch format {
	case JsonArray:
		return currencies.ToJsonArray()
	case JsonMap:
		return currencies.ToJsonMap()
	case Csv:
		return currencies.ToCSV()
	default:
		return nil, fmt.Errorf("invalid format: %s", format)
	}
}

// Returns Currencies from json array that has format:
//
//	[{"index":"USD","conversionRate":2.0},{"index":"EUR","conversionRate":3.0}]
func GetCurrenciesFromJsonArray(jsonArray []byte) (Currencies, error) {
	currencies := Currencies{}

	err := json.Unmarshal(jsonArray, &currencies)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}

const (
	NBG_CURRENCIES_URL                = "https://nbg.gov.ge/en/monetary-policy/currency"
	CURRENCIES_SELECTOR               = ".mt-3-4.border-b-2.border-grey-400.border-solid > .jsx-182984682"
	CURRENCY_INDEX_SELECTOR           = ".jsx-182984682.px-2-2 > span"
	CURRENCY_CONVERSION_RATE_SELECTOR = ".jsx-182984682.flex.items-center.justify-end > span"
)

// Fetches currencies and their conversion rates relative to GEL from
// https://nbg.gov.ge/en/monetary-policy/currency
func FetchCurrencies() (Currencies, error) {
	fmt.Println("Fetching currencies...")
	fetchStart := time.Now()

	resp, err := http.Get(NBG_CURRENCIES_URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			"status code error: %d %s", resp.StatusCode, resp.Status,
		)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	currencies := Currencies{}

	currenciesSelector := document.Find(CURRENCIES_SELECTOR)
	currenciesSelector.Each(func(i int, s *goquery.Selection) {
		index := s.Find(CURRENCY_INDEX_SELECTOR).Text()
		conversionRateStr := s.Find(CURRENCY_CONVERSION_RATE_SELECTOR).Text()

		conversionRate, err := strconv.ParseFloat(conversionRateStr, 64)
		if err != nil {
			return
		}

		currencies = append(currencies, Currency{index, conversionRate})
	})

	fmt.Printf(
		"Currencies fetched in %d milliseconds.\n",
		time.Since(fetchStart)/time.Millisecond,
	)

	return currencies, nil
}
