package common

import (
	"fmt"
)

func PrintAsm(asm string, item ...interface{}) {
	fmt.Printf("\t\t")
	fmt.Printf(asm, item...)
	fmt.Println()
}

func PrintLabel(label string, item ...interface{}) {
	fmt.Printf(".L__")
	fmt.Printf(label, item...)
	fmt.Println(":")
}
