package asm

import (
	"fmt"
)

// Header : アセンブリヘッダーを出力する
func Header() {
	fmt.Println(".intel_syntax	noprefix")
	fmt.Println(".global		_main")
	fmt.Println()
	fmt.Println("_main:")
	fmt.Println("		mov rbx, 0")
}
