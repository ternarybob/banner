package main

import (
	"fmt"
	"github.com/ternarybob/banner"
)

func main() {
	fmt.Println("\n=== Simple Banner ===")
	banner.PrintSimple("MY APPLICATION", "Version 1.0.0")

	fmt.Println("\n=== Colorized Banner ===")
	banner.PrintColorized("AWESOME SERVICE", "High Performance Server", banner.ColorPurple, banner.ColorYellow)

	fmt.Println("\n=== Custom Banner with Details ===")
	b := banner.New().
		SetStyle(banner.StyleDouble).
		SetBorderColor(banner.ColorCyan).
		SetTextColor(banner.ColorWhite).
		SetBold(true)

	b.PrintTopLine()
	b.PrintCenteredText("CUSTOM APPLICATION")
	b.PrintCenteredText("Enterprise Edition")
	b.PrintSeparatorLine()
	b.PrintKeyValue("Version", "2.5.0", 12)
	b.PrintKeyValue("Environment", "Production", 12)
	b.PrintKeyValue("Port", "8080", 12)
	b.PrintBottomLine()

	fmt.Println("\n=== Different Styles ===")

	styles := []struct {
		name  string
		style banner.Style
		color string
	}{
		{"Simple Style", banner.StyleSimple, banner.ColorGreen},
		{"Double Style", banner.StyleDouble, banner.ColorBlue},
		{"Bold Style", banner.StyleBold, banner.ColorRed},
		{"Round Style", banner.StyleRound, banner.ColorPurple},
		{"ASCII Style", banner.StyleASCII, ""},
	}

	for _, s := range styles {
		fmt.Printf("\n%s:\n", s.name)
		b := banner.New().
			SetStyle(s.style).
			SetBorderColor(s.color).
			SetWidth(60)

		b.PrintTopLine()
		b.PrintCenteredText("Demo Application")
		b.PrintSeparatorLine()
		b.PrintText("Left aligned text")
		b.PrintCenteredText("Center aligned text")
		b.PrintRightText("Right aligned text")
		b.PrintBottomLine()
	}

	fmt.Println("\n=== Complex Layout ===")
	complex := banner.New().
		SetStyle(banner.StyleBold).
		SetBorderColor(banner.ColorPurple).
		SetTextColor(banner.ColorCyan).
		SetWidth(70)

	complex.PrintTopLine()
	complex.PrintEmptyLine()
	complex.PrintCenteredText("SERVICE MONITOR")
	complex.PrintEmptyLine()
	complex.PrintSeparatorLine()
	complex.PrintText("Status: Online")
	complex.PrintText("Requests: 1,234,567")
	complex.PrintText("Uptime: 99.99%")
	complex.PrintSeparatorLine()
	complex.PrintKeyValue("CPU", "45%", 10)
	complex.PrintKeyValue("Memory", "2.3GB", 10)
	complex.PrintKeyValue("Disk", "120GB", 10)
	complex.PrintBottomLine()
}
