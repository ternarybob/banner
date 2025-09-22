package banner

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestNew(t *testing.T) {
	b := New()

	if b.Width != 80 {
		t.Errorf("Expected width 80, got %d", b.Width)
	}

	if b.Style != StyleSimple {
		t.Errorf("Expected StyleSimple, got %v", b.Style)
	}

	if b.BorderColor != "" {
		t.Errorf("Expected empty border color, got %s", b.BorderColor)
	}

	if b.TextColor != "" {
		t.Errorf("Expected empty text color, got %s", b.TextColor)
	}

	if b.Bold != false {
		t.Errorf("Expected Bold false, got %v", b.Bold)
	}
}

func TestSetters(t *testing.T) {
	b := New()

	// Test method chaining
	result := b.SetWidth(60).
		SetStyle(StyleDouble).
		SetBorderColor(ColorRed).
		SetTextColor(ColorBlue).
		SetBold(true)

	if result != b {
		t.Error("Setters should return the same banner instance for chaining")
	}

	if b.Width != 60 {
		t.Errorf("Expected width 60, got %d", b.Width)
	}

	if b.Style != StyleDouble {
		t.Errorf("Expected StyleDouble, got %v", b.Style)
	}

	if b.BorderColor != ColorRed {
		t.Errorf("Expected ColorRed, got %s", b.BorderColor)
	}

	if b.TextColor != ColorBlue {
		t.Errorf("Expected ColorBlue, got %s", b.TextColor)
	}

	if b.Bold != true {
		t.Errorf("Expected Bold true, got %v", b.Bold)
	}
}

func TestSetStyleInvalid(t *testing.T) {
	b := New()
	original := b.Style

	// Try to set invalid style
	b.SetStyle(Style("invalid"))

	if b.Style != original {
		t.Error("Invalid style should not change the banner style")
	}
}

func TestBorderStyles(t *testing.T) {
	tests := []struct {
		style    Style
		expected string
	}{
		{StyleSimple, "┌"},
		{StyleDouble, "╔"},
		{StyleBold, "┏"},
		{StyleRound, "╭"},
		{StyleASCII, "+"},
	}

	for _, test := range tests {
		t.Run(string(test.style), func(t *testing.T) {
			b := New().SetStyle(test.style).SetWidth(10)
			output := captureOutput(func() {
				b.PrintTopLine()
			})

			if !strings.Contains(output, test.expected) {
				t.Errorf("Expected output to contain %s, got %s", test.expected, output)
			}
		})
	}
}

func TestPrintTopLine(t *testing.T) {
	b := New().SetWidth(10)
	output := captureOutput(func() {
		b.PrintTopLine()
	})

	expected := "┌────────┐\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintBottomLine(t *testing.T) {
	b := New().SetWidth(10)
	output := captureOutput(func() {
		b.PrintBottomLine()
	})

	expected := "└────────┘\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintSeparatorLine(t *testing.T) {
	b := New().SetWidth(10)
	output := captureOutput(func() {
		b.PrintSeparatorLine()
	})

	expected := "├────────┤\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintText(t *testing.T) {
	b := New().SetWidth(20)
	output := captureOutput(func() {
		b.PrintText("Hello")
	})

	expected := "│ Hello            │\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintCenteredText(t *testing.T) {
	b := New().SetWidth(20)
	output := captureOutput(func() {
		b.PrintCenteredText("Hello")
	})

	expected := "│      Hello       │\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintRightText(t *testing.T) {
	b := New().SetWidth(20)
	output := captureOutput(func() {
		b.PrintRightText("Hello")
	})

	expected := "│            Hello │\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintEmptyLine(t *testing.T) {
	b := New().SetWidth(10)
	output := captureOutput(func() {
		b.PrintEmptyLine()
	})

	expected := "│        │\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintKeyValue(t *testing.T) {
	b := New().SetWidth(30)
	output := captureOutput(func() {
		b.PrintKeyValue("Version", "1.0.0", 8)
	})

	expected := "│ Version: 1.0.0             │\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestTextTruncation(t *testing.T) {
	b := New().SetWidth(10)
	output := captureOutput(func() {
		b.PrintText("This is a very long text that should be truncated")
	})

	// Width 10 - 4 (borders and spaces) = 6 characters for text
	if !strings.Contains(output, "This i") {
		t.Errorf("Text should be truncated to fit width, got %q", output)
	}
}

func TestCenterTextFunction(t *testing.T) {
	b := New()

	tests := []struct {
		text     string
		width    int
		expected string
	}{
		{"Hello", 10, "  Hello   "},
		{"Hi", 6, "  Hi  "},
		{"Test", 4, "Test"},
		{"Toolong", 4, "Tool"},
	}

	for _, test := range tests {
		result := b.centerText(test.text, test.width)
		if result != test.expected {
			t.Errorf("centerText(%q, %d) = %q, expected %q",
				test.text, test.width, result, test.expected)
		}
	}
}

func TestLeftAlignTextFunction(t *testing.T) {
	b := New()

	tests := []struct {
		text     string
		width    int
		expected string
	}{
		{"Hello", 10, "Hello     "},
		{"Hi", 6, "Hi    "},
		{"Test", 4, "Test"},
		{"Toolong", 4, "Tool"},
	}

	for _, test := range tests {
		result := b.leftAlignText(test.text, test.width)
		if result != test.expected {
			t.Errorf("leftAlignText(%q, %d) = %q, expected %q",
				test.text, test.width, result, test.expected)
		}
	}
}

func TestRightAlignTextFunction(t *testing.T) {
	b := New()

	tests := []struct {
		text     string
		width    int
		expected string
	}{
		{"Hello", 10, "     Hello"},
		{"Hi", 6, "    Hi"},
		{"Test", 4, "Test"},
		{"Toolong", 4, "Tool"},
	}

	for _, test := range tests {
		result := b.rightAlignText(test.text, test.width)
		if result != test.expected {
			t.Errorf("rightAlignText(%q, %d) = %q, expected %q",
				test.text, test.width, result, test.expected)
		}
	}
}

func TestApplyColor(t *testing.T) {
	b := New()

	// Test with color
	result := b.applyColor(ColorRed, "Hello")
	expected := ColorRed + "Hello" + ColorReset
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test without color
	result = b.applyColor("", "Hello")
	if result != "Hello" {
		t.Errorf("Expected %q, got %q", "Hello", result)
	}
}

func TestPrintColorized(t *testing.T) {
	b := New()

	// Test with color
	output := captureOutput(func() {
		b.printColorized(ColorRed, "Hello")
	})
	expected := ColorRed + "Hello" + ColorReset + "\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

	// Test without color
	output = captureOutput(func() {
		b.printColorized("", "Hello")
	})
	expected = "Hello\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintSimple(t *testing.T) {
	output := captureOutput(func() {
		PrintSimple("Title", "Subtitle")
	})

	// Should contain both title and subtitle
	if !strings.Contains(output, "Title") {
		t.Error("Output should contain title")
	}
	if !strings.Contains(output, "Subtitle") {
		t.Error("Output should contain subtitle")
	}
	// Should have top and bottom borders
	if !strings.Contains(output, "┌") || !strings.Contains(output, "└") {
		t.Error("Output should contain top and bottom borders")
	}
}

func TestPrintColorizedFunction(t *testing.T) {
	output := captureOutput(func() {
		PrintColorized("Title", "Subtitle", ColorRed, ColorBlue)
	})

	// Should contain both title and subtitle
	if !strings.Contains(output, "Title") {
		t.Error("Output should contain title")
	}
	if !strings.Contains(output, "Subtitle") {
		t.Error("Output should contain subtitle")
	}
	// Should contain color codes
	if !strings.Contains(output, ColorRed) {
		t.Error("Output should contain border color")
	}
	if !strings.Contains(output, ColorBlue) {
		t.Error("Output should contain text color")
	}
}

func TestBoldTextFormatting(t *testing.T) {
	b := New().SetWidth(20).SetBold(true).SetTextColor(ColorBlue)
	output := captureOutput(func() {
		b.PrintText("Hello")
	})

	// Should contain bold formatting when bold is enabled
	if !strings.Contains(output, ColorBold) {
		t.Error("Output should contain bold formatting")
	}
}

func TestColorConstants(t *testing.T) {
	// Test that color constants are properly defined
	colors := map[string]string{
		"ColorReset":  ColorReset,
		"ColorRed":    ColorRed,
		"ColorGreen":  ColorGreen,
		"ColorYellow": ColorYellow,
		"ColorBlue":   ColorBlue,
		"ColorPurple": ColorPurple,
		"ColorCyan":   ColorCyan,
		"ColorWhite":  ColorWhite,
		"ColorBold":   ColorBold,
	}

	for name, color := range colors {
		if color == "" {
			t.Errorf("%s should not be empty", name)
		}
		if !strings.HasPrefix(color, "\033[") {
			t.Errorf("%s should be an ANSI escape sequence, got %q", name, color)
		}
	}
}

func BenchmarkPrintText(b *testing.B) {
	banner := New().SetWidth(80)

	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		banner.PrintText("Hello World")
	}
}

func BenchmarkPrintSimple(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintSimple("Title", "Subtitle")
	}
}

func ExampleNew() {
	b := New().SetWidth(40).SetStyle(StyleDouble)
	b.PrintTopLine()
	b.PrintCenteredText("Example Banner")
	b.PrintBottomLine()
	// Output will show a banner with double-line borders
}

func ExamplePrintSimple() {
	PrintSimple("My App", "Version 1.0")
	// Output:
	// ┌──────────────────────────────────────────────────────────────────────────────┐
	// │                                    My App                                    │
	// │                                 Version 1.0                                  │
	// └──────────────────────────────────────────────────────────────────────────────┘
}

func ExampleBanner_PrintKeyValue() {
	b := New().SetWidth(30)
	b.PrintTopLine()
	b.PrintKeyValue("Version", "1.0.0", 8)
	b.PrintKeyValue("Status", "Running", 8)
	b.PrintBottomLine()
	// Output shows formatted key-value pairs
}