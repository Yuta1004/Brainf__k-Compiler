package asm

import (
	"fmt"
)

// Header : アセンブリヘッダーを出力する
func Header(allocNum int) {
	fmt.Println(".intel_syntax	noprefix")
	fmt.Println(".global		_main")
	fmt.Println()
	fmt.Println("_main:")
	fmt.Println("		push rbp")
	fmt.Println("		mov rbp, rsp")
	fmt.Printf("		add rsp, %d\n", allocNum*8)
	fmt.Println("		mov rbx, 0")
}
