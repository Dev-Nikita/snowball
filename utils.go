package snowball

import (
	"strings"
	"unicode/utf8"
)

// Finds the region after the first non-vowel following a vowel,
// or is the null region at the end of the word if there is no
// such non-vowel.
//
func vnvSuffix(word string) string {
	runes := []rune(word)
	// uscular
	for i := 1; i < len(runes); i++ {
		if isLowerVowel(runes[i-1]) && !isLowerVowel(runes[i]) {
			return string(runes[i+1:])
		}
	}
	return ""
}

// R1 is the region after the first non-vowel following a vowel,
// or is the null region at the end of the word if there is no
// such non-vowel.
//
// R2 is the region after the first non-vowel following a vowel
// in R1, or is the null region at the end of the word if there
// is no such non-vowel.
//
// See http://snowball.tartarus.org/texts/r1r2.html
//
func r1r2(word string) (r1, r2 string) {

	specialPrefixes := []string{"gener", "commun", "arsen"}
	hasSpecialPrefix := false
	specialPrefix := ""
	for _, specialPrefix = range specialPrefixes {
		if strings.HasPrefix(word, specialPrefix) {
			hasSpecialPrefix = true
			break
		}
	}

	if hasSpecialPrefix {
		if specialPrefix == "commun" {
			r1 = word[6:]
		} else {
			r1 = word[5:]
		}

	} else {
		r1 = vnvSuffix(word)
	}
	r2 = vnvSuffix(r1)
	return
}

// Test if a string has a rune, skipping parts of the string
// that are less than `leftSkip` of the beginning and `rightSkip`
// of the end.
//
func hasRune(word string, leftSkip int, rightSkip int, testRunes ...rune) bool {
	leftMin := leftSkip
	rightMax := utf8.RuneCountInString(word) - rightSkip
	for i, r := range word {
		if i < leftMin {
			continue
		} else if i >= rightMax {
			break
		}
		for _, tr := range testRunes {
			if r == tr {
				return true
			}
		}
	}
	return false
}

// Test if a string has a vowel, skipping parts of the string
// that are less than `leftSkip` of the beginning and `rightSkip`
// of the end.  (All counts in runes.)
//
func hasVowel(word string, leftSkip int, rightSkip int) bool {
	return hasRune(word, leftSkip, rightSkip, 97, 101, 105, 111, 117, 121)
}

// Checks if a rune is a lowercase English vowel.
//
func isLowerVowel(r rune) bool {
	switch r {
	case 97, 101, 105, 111, 117, 121:
		return true
	}
	return false
}