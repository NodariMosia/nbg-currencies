# NBG Currencies

## About

This project is a CLI tool that enables user to get GEL (Georgian Lari)
conversion rates. The project scrapes
[National Bank of Georgia](https://nbg.gov.ge/en/monetary-policy/currency)'s
website and parses currency data. Users are able to specify format and filename
from CLI to generate data in different formats, or build and run the binary and
generate all formats at once.

## Setup

```bash
go mod tidy
go mod download
go build -o bin/ ./cmd
```

## Usage

```bash
go run ./cmd <format> <file name>
```

Arguments **format** and **filename** are optional.

When both arguments are omitted, program will generate files for all supported
formats with default generated names.

When **filename** is omitted, program will generate a file for provided
**format** with a default generated name.

Supported formats are: **json-array**, **json-map**, **csv**.

## Examples

```bash
go run ./cmd
```

```bash
go run ./cmd json-array
```

```bash
go run ./cmd json-map currencies-map.json
```

```bash
go run ./cmd csv currencies.csv
```
