package currency

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	nbgCurrenciesUrl                  = "https://nbg.gov.ge/en/monetary-policy/currency"
	currenciesCssSelector             = ".mt-3-4.border-b-2.border-grey-400.border-solid > .jsx-182984682"
	currencyIndexCssSelector          = ".jsx-182984682.px-2-2 > span"
	currencyConversionRateCssSelector = ".jsx-182984682.flex.items-center.justify-end > span"
)

// Fetches currencies and their conversion rates relative to GEL from
// https://nbg.gov.ge/en/monetary-policy/currency
func FetchCurrencies() (currencies Currencies, elapsedMs int64, err error) {
	startTime := time.Now()

	resp, err := http.Get(nbgCurrenciesUrl)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, 0, fmt.Errorf(
			"status code error: %d %s", resp.StatusCode, resp.Status,
		)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	currenciesSelection := document.Find(currenciesCssSelector)
	currenciesSelection.Each(func(i int, s *goquery.Selection) {
		index := s.Find(currencyIndexCssSelector).Text()
		conversionRateStr := s.Find(currencyConversionRateCssSelector).Text()

		conversionRate, err := strconv.ParseFloat(conversionRateStr, 64)
		if err != nil {
			return
		}

		currencies = append(currencies, Currency{index, conversionRate})
	})

	elapsedMs = int64(time.Since(startTime) / time.Millisecond)

	return currencies, elapsedMs, nil
}
