package asm

import (
	"fmt"
)

// Footer : アセンブリフッターを出力する
func Footer() {
	fmt.Println("")
	fmt.Println("footer:")
	fmt.Println("		mov rsp, rbp")
	fmt.Println("		pop rbp")
	fmt.Println("		ret")
}
