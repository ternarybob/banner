# Banner

A flexible Go library for creating styled console banners with customizable borders, colors, and text alignment.

## Installation

```bash
go get github.com/ternarybob/banner
```

## Features

- **Multiple Border Styles**: Simple, Double, Bold, Round, and ASCII
- **Color Support**: Full ANSI color support for borders and text
- **Text Alignment**: Left, center, and right alignment options
- **Flexible Layout**: Top/bottom lines, separators, empty lines, and key-value pairs
- **Fluent API**: Method chaining for easy configuration

## Quick Start

```go
package main

import "github.com/ternarybob/banner"

func main() {
    // Simple banner
    banner.PrintSimple("MY APPLICATION", "Version 1.0.0")

    // Colorized banner
    banner.PrintColorized("SERVICE", "Running", banner.ColorPurple, banner.ColorYellow)
}
```

## Border Styles

The library supports 5 different border styles:

| Style | Example |
|-------|---------|
| `StyleSimple` | `┌─────┐` |
| `StyleDouble` | `╔═════╗` |
| `StyleBold` | `┏━━━━━┓` |
| `StyleRound` | `╭─────╮` |
| `StyleASCII` | `+-----+` |

## Colors

Available color constants:
- `ColorReset`, `ColorRed`, `ColorGreen`, `ColorYellow`
- `ColorBlue`, `ColorPurple`, `ColorCyan`, `ColorWhite`
- `ColorBrightGreen` - Standard 80's console green
- `ColorPrimaryGreen` - Aktis primary green (#00FF00 - default border color)
- `ColorBold` - Bold text modifier

## Basic Usage

### Simple Banners

```go
// Basic banner with title and subtitle
banner.PrintSimple("MY APP", "Version 1.0")

// Colored banner
banner.PrintColorized("STATUS", "Online", banner.ColorGreen, banner.ColorWhite)
```

### Custom Banners

```go
// Create a custom banner
b := banner.New().
    SetStyle(banner.StyleDouble).
    SetBorderColor(banner.ColorCyan).
    SetTextColor(banner.ColorWhite).
    SetBold(true).
    SetWidth(60)

b.PrintTopLine()
b.PrintCenteredText("CUSTOM APPLICATION")
b.PrintSeparatorLine()
b.PrintKeyValue("Version", "2.5.0", 12)
b.PrintKeyValue("Environment", "Production", 12)
b.PrintBottomLine()
```

### Text Alignment

```go
b := banner.New()
b.PrintTopLine()
b.PrintText("Left aligned text")           // Default left alignment
b.PrintCenteredText("Center aligned text") // Center alignment
b.PrintRightText("Right aligned text")     // Right alignment
b.PrintBottomLine()
```

## Advanced Example

```go
package main

import (
    "fmt"
    "github.com/ternarybob/banner"
)

func main() {
    // Service status banner
    status := banner.New().
        SetStyle(banner.StyleBold).
        SetBorderColor(banner.ColorPurple).
        SetTextColor(banner.ColorCyan).
        SetWidth(70)

    status.PrintTopLine()
    status.PrintEmptyLine()
    status.PrintCenteredText("SERVICE MONITOR")
    status.PrintEmptyLine()
    status.PrintSeparatorLine()
    status.PrintText("Status: Online")
    status.PrintText("Requests: 1,234,567")
    status.PrintText("Uptime: 99.99%")
    status.PrintSeparatorLine()
    status.PrintKeyValue("CPU", "45%", 10)
    status.PrintKeyValue("Memory", "2.3GB", 10)
    status.PrintKeyValue("Disk", "120GB", 10)
    status.PrintBottomLine()

    // Different styles demonstration
    styles := []struct {
        name  string
        style banner.Style
        color string
    }{
        {"Simple", banner.StyleSimple, banner.ColorGreen},
        {"Double", banner.StyleDouble, banner.ColorBlue},
        {"Bold", banner.StyleBold, banner.ColorRed},
        {"Round", banner.StyleRound, banner.ColorPurple},
        {"ASCII", banner.StyleASCII, ""},
    }

    for _, s := range styles {
        fmt.Printf("\n%s Style:\n", s.name)
        b := banner.New().
            SetStyle(s.style).
            SetBorderColor(s.color).
            SetWidth(50)

        b.PrintTopLine()
        b.PrintCenteredText("Demo Application")
        b.PrintSeparatorLine()
        b.PrintText("Left aligned")
        b.PrintCenteredText("Center aligned")
        b.PrintRightText("Right aligned")
        b.PrintBottomLine()
    }
}
```

## API Reference

### Banner Methods

#### Configuration
- `SetStyle(style Style) *Banner` - Set border style
- `SetWidth(width int) *Banner` - Set banner width
- `SetBorderColor(color string) *Banner` - Set border color
- `SetTextColor(color string) *Banner` - Set text color
- `SetBold(bold bool) *Banner` - Enable/disable bold text

#### Printing
- `PrintTopLine()` - Print top border
- `PrintBottomLine()` - Print bottom border
- `PrintSeparatorLine()` - Print separator line
- `PrintText(text string)` - Print left-aligned text
- `PrintCenteredText(text string)` - Print centered text
- `PrintRightText(text string)` - Print right-aligned text
- `PrintEmptyLine()` - Print empty line with borders
- `PrintKeyValue(key, value string, padding int)` - Print formatted key-value pair

### Convenience Functions
- `PrintSimple(title, subtitle string)` - Quick simple banner
- `PrintColorized(title, subtitle string, borderColor, textColor string)` - Quick colored banner

## Output Examples

### Simple Banner
```
┌──────────────────────────────────────────────────────────────────────────────┐
│                                MY APPLICATION                                │
│                                Version 1.0.0                                 │
└──────────────────────────────────────────────────────────────────────────────┘
```

### Service Monitor
```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                    ┃
┃                          SERVICE MONITOR                           ┃
┃                                                                    ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ Status: Online                                                     ┃
┃ Requests: 1,234,567                                                ┃
┃ Uptime: 99.99%                                                     ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ CPU:       45%                                                     ┃
┃ Memory:    2.3GB                                                   ┃
┃ Disk:      120GB                                                   ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

## Release Process

This project uses automatic semantic versioning with GitHub releases. Every commit to the `main` branch automatically:

1. **Increments** the patch version (e.g., v1.0.0 → v1.0.1)
2. **Creates** a git tag with the new version
3. **Publishes** a GitHub release with auto-generated release notes
4. **Updates** the go module version for `go get`

### Version Management

For local development, use the provided scripts:

```bash
# Check current version
./scripts/version.sh current

# Increment versions locally (optional)
./scripts/version.sh patch   # v1.0.0 → v1.0.1
./scripts/version.sh minor   # v1.0.0 → v1.1.0
./scripts/version.sh major   # v1.0.0 → v2.0.0
```

Or on Windows:
```powershell
# Check current version
.\scripts\version.ps1 current

# Increment versions locally (optional)
.\scripts\version.ps1 patch   # v1.0.0 → v1.0.1
.\scripts\version.ps1 minor   # v1.0.0 → v1.1.0
.\scripts\version.ps1 major   # v1.0.0 → v2.0.0
```

### Manual Releases

You can also create manual releases by pushing tags:

```bash
git tag v1.5.0
git push origin v1.5.0
```

This will create a release without auto-incrementing the version.

## License

MIT License