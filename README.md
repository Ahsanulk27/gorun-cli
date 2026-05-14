# goku

A CLI tool for converting configuration files between JSON and YAML formats.

## Features

- Convert JSON → YAML and YAML → JSON
- Saves the converted file in the same directory as the input
- `--dry-run` flag to preview output without saving
- Smart truncation — previews first 20 lines for large files
- Handles edge cases: same format conversion, invalid extensions, missing files

## Installation

### Using go install (recommended)
```bash
go install github.com/Ahsanulk27/goku@latest
```
Then use it anywhere on your machine:
```bash
goku -i config.json -o yaml
```

### From source
```bash
git clone https://github.com/Ahsanulk27/goku.git
cd goku
go build -o goku .
```

## Usage

```bash
goku -i <input file> -o <output format>
```

**Convert JSON to YAML:**
```bash
goku -i config.json -o yaml
```

**Convert YAML to JSON:**
```bash
goku -i config.yaml -o json
```

**Preview without saving:**
```bash
goku -i config.json -o yaml --dry-run
```

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--input` | `-i` | Path to the input file (required) |
| `--output` | `-o` | Output format: `json` or `yaml` (required) |
| `--dry-run` | `-d` | Preview converted output without saving |

## Error Handling

| Scenario | Error Message |
|----------|---------------|
| Same input and output format | `requested format should be different than the input format` |
| Unsupported file extension | `invalid file format .txt` |
| File not found | `file not found: config.json` |
| Missing flags | `required flag(s) "input", "output" not set` |

## Project Structure

```
goku/
├── main.go          # Entry point
├── cmd/
│   └── root.go      # CLI flags and command setup
└── internal/
    └── converter.go  # Core conversion logic
```
