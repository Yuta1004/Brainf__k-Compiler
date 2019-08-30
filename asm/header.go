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
	fmt.Println("header:")
	fmt.Println("		push rbp")
	fmt.Println("		mov rbp, rsp")
	fmt.Println("		mov rdi, 0")
	fmt.Println("__init_stack:")
	fmt.Println("		push 0")
	fmt.Println("		add rdi, 1")
	fmt.Printf("		cmp rdi, %d\n", allocNum)
	fmt.Println("		jb __init_stack")
	fmt.Println("		mov rdi, 0")
	fmt.Println("")
}
