package asm

import (
	"fmt"
)

// Footer : アセンブリフッターを出力する
func Footer() {
	fmt.Println("		mov rax, 0")
	fmt.Println("		ret")
}
