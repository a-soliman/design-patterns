package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1. built-in builder
	hello := "Hello"
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("<p>")
	stringBuilder.WriteString(hello)
	stringBuilder.WriteString("</p>")
	fmt.Println(stringBuilder.String())
	stringBuilder.Reset()

	// 2. un-ordered list using built-in builder
	words := []string{"hello", "world"}

	stringBuilder.WriteString("<ul>")
	for _, word := range words {
		stringBuilder.WriteString("<li>")
		stringBuilder.WriteString(word)
		stringBuilder.WriteString("</li>")
	}
	stringBuilder.WriteString("</ul>")
	fmt.Print(stringBuilder.String())
	stringBuilder.Reset()
}
