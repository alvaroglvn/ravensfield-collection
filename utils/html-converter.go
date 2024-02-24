package utils

import (
	"github.com/russross/blackfriday/v2"
)

func MarkdownToHTML(markdown []byte) []byte {
	output := blackfriday.Run(markdown)
	return output
}
