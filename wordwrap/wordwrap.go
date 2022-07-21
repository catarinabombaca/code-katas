package wordwrap

import (
	"strings"
)

func Execute(text string, maxChars int) string {
	var result string
	paragraphs := strings.Split(text, "\n\n")

	for _, paragraph := range paragraphs {
		paragraphWithLimit := ""
		lineWithLimit := ""
		words := strings.Split(paragraph, " ")

		for i, word := range words {
			if len(lineWithLimit)+len(word) <= maxChars {
				lineWithLimit += word + " "
			} else {
				paragraphWithLimit += strings.TrimSpace(lineWithLimit) + "\n"
				lineWithLimit = word + " "
			}

			if i == len(words)-1 {
				paragraphWithLimit += strings.TrimSpace(lineWithLimit) + "\n"
			}
		}

		result += paragraphWithLimit + "\n"
	}

	return strings.TrimRight(result, "\n\n")
}
