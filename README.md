# NBG Currencies

## About

This project is a simple golang TUI (Terminal User Interface) tool that enables
user to get GEL (Georgian Lari) conversion rates. The project scrapes
[National Bank of Georgia](https://nbg.gov.ge/en/monetary-policy/currency)'s
website and parses currency data. Users are able to specify file formats from
TUI to generate data in different formats.

Project uses [goquery](https://github.com/PuerkitoBio/goquery) for parsing HTML
and [bubbletea](https://github.com/charmbracelet/bubbletea) to build the
terminal user interface.

Supported formats are: **json-array**, **json-map**, **csv**.

## Setup

```bash
make setup

# or

go mod tidy
go mod download
```

## Usage

### From Command Line

```bash
make run

# or

go run ./cmd
```

### From Binary Executable

```bash
make build
make launch

# or

go build -o bin/ ./cmd
./bin/cmd
```
