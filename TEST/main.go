package main

import (
	"fmt"

	"github.com/russross/blackfriday"
)

func main() {

	input := []byte("STRING")
	output := blackfriday.MarkdownBasic(input)
	fmt.Println(string(output))
}
