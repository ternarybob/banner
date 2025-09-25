package banner

import (
	"fmt"
	"strings"
)

const (
	ColorReset        = "\033[0m"
	ColorRed          = "\033[31m"
	ColorGreen        = "\033[32m"
	ColorYellow       = "\033[33m"
	ColorBlue         = "\033[34m"
	ColorPurple       = "\033[35m"
	ColorCyan         = "\033[36m"
	ColorWhite        = "\033[37m"
	ColorBold         = "\033[1m"
	ColorBrightGreen  = "\033[92m"
	ColorPrimaryGreen = "\033[38;2;0;255;0m" // #00FF00 - Aktis primary green
)

type Style string

const (
	StyleSimple Style = "simple"
	StyleDouble Style = "double"
	StyleBold   Style = "bold"
	StyleRound  Style = "round"
	StyleASCII  Style = "ascii"
)

type BorderChars struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
	LeftJoin    string
	RightJoin   string
}

var borderStyles = map[Style]BorderChars{
	StyleSimple: {
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
		Horizontal:  "─",
		Vertical:    "│",
		LeftJoin:    "├",
		RightJoin:   "┤",
	},
	StyleDouble: {
		TopLeft:     "╔",
		TopRight:    "╗",
		BottomLeft:  "╚",
		BottomRight: "╝",
		Horizontal:  "═",
		Vertical:    "║",
		LeftJoin:    "╠",
		RightJoin:   "╣",
	},
	StyleBold: {
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┗",
		BottomRight: "┛",
		Horizontal:  "━",
		Vertical:    "┃",
		LeftJoin:    "┣",
		RightJoin:   "┫",
	},
	StyleRound: {
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
		Horizontal:  "─",
		Vertical:    "│",
		LeftJoin:    "├",
		RightJoin:   "┤",
	},
	StyleASCII: {
		TopLeft:     "+",
		TopRight:    "+",
		BottomLeft:  "+",
		BottomRight: "+",
		Horizontal:  "-",
		Vertical:    "|",
		LeftJoin:    "+",
		RightJoin:   "+",
	},
}

type Banner struct {
	Width       int
	Style       Style
	BorderColor string
	TextColor   string
	Bold        bool
	borders     BorderChars
}

func New() *Banner {
	return &Banner{
		Width:       80,
		Style:       StyleSimple,
		BorderColor: ColorPrimaryGreen,
		TextColor:   "",
		Bold:        false,
		borders:     borderStyles[StyleSimple],
	}
}

func (b *Banner) SetStyle(style Style) *Banner {
	if borders, ok := borderStyles[style]; ok {
		b.Style = style
		b.borders = borders
	}
	return b
}

func (b *Banner) SetWidth(width int) *Banner {
	b.Width = width
	return b
}

func (b *Banner) SetBorderColor(color string) *Banner {
	b.BorderColor = color
	return b
}

func (b *Banner) SetTextColor(color string) *Banner {
	b.TextColor = color
	return b
}

func (b *Banner) SetBold(bold bool) *Banner {
	b.Bold = bold
	return b
}

func (b *Banner) PrintTopLine() {
	innerWidth := b.Width - 2
	line := b.borders.TopLeft + strings.Repeat(b.borders.Horizontal, innerWidth) + b.borders.TopRight
	b.printColorized(b.BorderColor, line)
}

func (b *Banner) PrintBottomLine() {
	innerWidth := b.Width - 2
	line := b.borders.BottomLeft + strings.Repeat(b.borders.Horizontal, innerWidth) + b.borders.BottomRight
	b.printColorized(b.BorderColor, line)
}

func (b *Banner) PrintSeparatorLine() {
	innerWidth := b.Width - 2
	line := b.borders.LeftJoin + strings.Repeat(b.borders.Horizontal, innerWidth) + b.borders.RightJoin
	b.printColorized(b.BorderColor, line)
}

func (b *Banner) PrintText(text string) {
	b.PrintTextWithAlignment(text, AlignLeft)
}

func (b *Banner) PrintCenteredText(text string) {
	b.PrintTextWithAlignment(text, AlignCenter)
}

func (b *Banner) PrintRightText(text string) {
	b.PrintTextWithAlignment(text, AlignRight)
}

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)

func (b *Banner) PrintTextWithAlignment(text string, align Alignment) {
	innerWidth := b.Width - 4

	if len(text) > innerWidth {
		text = text[:innerWidth]
	}

	var paddedText string
	switch align {
	case AlignCenter:
		paddedText = b.centerText(text, innerWidth)
	case AlignRight:
		paddedText = b.rightAlignText(text, innerWidth)
	default:
		paddedText = b.leftAlignText(text, innerWidth)
	}

	borderPrefix := b.applyColor(b.BorderColor, b.borders.Vertical)
	borderSuffix := b.applyColor(b.BorderColor, b.borders.Vertical)

	textColor := b.TextColor
	if b.Bold {
		textColor = ColorBold + textColor
	}

	textContent := b.applyColor(textColor, " "+paddedText+" ")

	fmt.Printf("%s%s%s\n", borderPrefix, textContent, borderSuffix)
}

func (b *Banner) PrintEmptyLine() {
	b.PrintText("")
}

func (b *Banner) PrintKeyValue(key, value string, padding int) {
	formatted := fmt.Sprintf("%-*s %s", padding, key+":", value)
	b.PrintText(formatted)
}

func (b *Banner) centerText(text string, width int) string {
	if len(text) >= width {
		return text[:width]
	}

	totalPadding := width - len(text)
	leftPad := totalPadding / 2
	rightPad := totalPadding - leftPad

	return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
}

func (b *Banner) leftAlignText(text string, width int) string {
	if len(text) >= width {
		return text[:width]
	}
	return text + strings.Repeat(" ", width-len(text))
}

func (b *Banner) rightAlignText(text string, width int) string {
	if len(text) >= width {
		return text[:width]
	}
	return strings.Repeat(" ", width-len(text)) + text
}

func (b *Banner) printColorized(color, text string) {
	if color != "" {
		fmt.Printf("%s%s%s\n", color, text, ColorReset)
	} else {
		fmt.Println(text)
	}
}

func (b *Banner) applyColor(color, text string) string {
	if color != "" {
		return color + text + ColorReset
	}
	return text
}

func PrintSimple(title, subtitle string) {
	b := New()
	b.PrintTopLine()
	if title != "" {
		b.PrintCenteredText(title)
	}
	if subtitle != "" {
		b.PrintCenteredText(subtitle)
	}
	b.PrintBottomLine()
}

func PrintColorized(title, subtitle string, borderColor, textColor string) {
	b := New().SetBorderColor(borderColor).SetTextColor(textColor).SetBold(true)
	b.PrintTopLine()
	if title != "" {
		b.PrintCenteredText(title)
	}
	if subtitle != "" {
		b.PrintCenteredText(subtitle)
	}
	b.PrintBottomLine()
}
