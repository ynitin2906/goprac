package solidprinciples

import "fmt"

type Reader interface {
	Read()
}
type Writer interface {
	Write(content string)
}
type Printer interface {
	Print()
}

type TextDocument struct {
	content string
}
type ReadDocument struct {
}

func (text *TextDocument) Read() {
	fmt.Println("Reading " + text.content)
}
func (text *TextDocument) Write(content string) {
	text.content = content
}
func (text *TextDocument) Print() {
	fmt.Println("Printing " + text.content)
}
func RunInterfaceSegregation() {
	text := TextDocument{content: "Hello"}
	text.Read()
	text.Write("Updated Hello")
	text.Print()
}
