package asm

import (
	c "brainfOOk-compiler/common"
	"fmt"
)

// Footer : アセンブリフッターを出力する
func Footer() {
	fmt.Println("")
	fmt.Println("footer:")
	c.PrintAsm("mov rsp, rbp")
	c.PrintAsm("pop rbp")
	c.PrintAsm("ret")
}
