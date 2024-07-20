package normalize

import (
	"strings"
	"unicode"
	"golang.org/x/text/unicode/norm"
)

func RemoveAccents(text string) string {

	t := norm.NFD.String(text)

	var b strings.Builder

	for _, r := range t {

		if unicode.Is(unicode.Mn, r) {
			continue

		}
		b.WriteRune(r)
	}

	return b.String()

}

func RemoveSpecialCharacters(text string) string {

	return strings.Map(func(r rune) rune {

		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			return r
		}

		return -1
	}, text)
}

func ToLowerCase(text string) string {

	return strings.ToLower(text)

}

func Text(text string) string {

	text = ToLowerCase(text)
	text = RemoveSpecialCharacters(text)
	text = RemoveAccents(text)
	text = strings.TrimSpace(text)

	return text
}
