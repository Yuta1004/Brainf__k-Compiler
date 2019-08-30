package main

import (
	"brainfOOk-compiler/asm"
	"brainfOOk-compiler/parse"
)

func main() {
	asm.Header()
	parse.Parse("++++")
	asm.Footer()
}
