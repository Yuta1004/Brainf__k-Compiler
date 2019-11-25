package asm

import (
	c "brainfOOk-compiler/common"
	"fmt"
)

// Header : アセンブリヘッダーを出力する
func Header(allocNum int) {
	// プロローグ
	fmt.Println(".intel_syntax	noprefix")
	fmt.Println(".global		_main")
	fmt.Println()
	fmt.Println("_main:")
	fmt.Println("header:")
	c.PrintAsm("push rbp")
	c.PrintAsm("mov rbp, rsp")
	c.PrintAsm("mov rdi, 0")

	// スタック初期化処理
	c.PrintLabel("init_stack")
	c.PrintAsm("push 0")
	c.PrintAsm("add rdi, 1")
	c.PrintAsm("cmp rdi, %d", allocNum)
	c.PrintAsm("jb .L__init_stack")
	c.PrintAsm("mov rdi, 0")
	fmt.Println("")
}
