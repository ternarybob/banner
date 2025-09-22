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
- `ColorBrightGreen` (80's console green - default border color)
- `ColorBold` (modifier)

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

## License

MIT License