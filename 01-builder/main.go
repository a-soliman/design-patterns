package main

import (
	"fmt"
	"strings"
)

const indentSize = 2

// HTMLElement type represents an element
type HTMLElement struct {
	name, text string
	elements   []HTMLElement
}

// String returns the indented string representation of the element
func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

// HTMLBuilder a type of the HTML builder
type HTMLBuilder struct {
	rootName string
	root     HTMLElement
}

// NewHTMLBuilder constructs a new HTMLBuilder and returns a pointer to it
func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{
		rootName,
		HTMLElement{
			rootName, "", []HTMLElement{},
		},
	}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

// AddChild adds a child to a given builder
func (b *HTMLBuilder) AddChild(childName, childText string) *HTMLBuilder {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

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

	// 3. using a new builder
	b := NewHTMLBuilder("ul")
	b.AddChild("li", "This is the first list item").
		AddChild("li", "This is the second list item").
		AddChild("li", "Another one")
	fmt.Println(b.String())
}
