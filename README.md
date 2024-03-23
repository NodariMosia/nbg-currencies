# NBG Currencies

## About

This project enables user to get GEL (Georgian Lari) conversion rates. The project scrapes [National Bank of Georgia's Website](https://nbg.gov.ge/en/monetary-policy/currency) and parses currency data. Users are able to specify format and filename to save data in different formats.

Format can be **json-array**, **json-map** or **csv**.

## Usage

```bash
go run ./cmd <format> <file name>
```

## Examples

```bash
go run ./cmd json-array currencies-array.json
```

```bash
go run ./cmd json-map currencies-map.json
```

```bash
go run ./cmd csv currencies.csv
```
