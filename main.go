package main

import (
	"brainfOOk-compiler/asm"
	"brainfOOk-compiler/common"
	"brainfOOk-compiler/parse"
	"flag"
	"strings"
)

func main() {
	flag.Parse()
	programFile := flag.Arg(0)
	if programFile == "" {
		common.Error("引数が少なすぎます!")
	}
	program := strings.Split(common.ReadFile(programFile), "\n")[0]

	programItemTop, allocMemoryNum := parse.Parse(program)

	asm.Header(allocMemoryNum)
	asm.Body(programItemTop)
	asm.Footer()
}
