package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		description := fmt.Sprintf("%d is converted to %s", test.Arabic, test.Roman)
		t.Run(description, func(t *testing.T) {
			want := test.Roman
			got, _ := ConvertToRoman(test.Arabic)

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		description := fmt.Sprintf("%s is converted to %d", test.Roman, test.Arabic)
		t.Run(description, func(t *testing.T) {
			want := test.Arabic
			got := ConvertToArabic(test.Roman)

			if got != want {
				t.Errorf("got %d, want %d", got, want)
				fmt.Printf("%s\n", (windowedRoman(test.Roman).Symbols()))
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}

		roman, _ := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 100}); err != nil {
		t.Error("failed checks", err)
	}
}

var romanCharacters = []byte{'I', 'X', 'V', 'C', 'D', 'L'}

func TestPropertyOfRepetition(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman, _ := ConvertToRoman(arabic)
		for _, symbol := range romanCharacters {
			repeatedChars := strings.Repeat(string([]byte{symbol}), 4)
			if res := strings.Contains(roman, repeatedChars); res {
				return false
			}
		}
		return true
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 100}); err != nil {
		t.Error("failed checks", err)
	}
}
