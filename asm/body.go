package asm

import (
	"brainfOOk-compiler/common"
	"brainfOOk-compiler/parse"
	"fmt"
)

// Body : ProgramInfoListをアセンブリに変換して出力する
func Body(programItemTop *parse.ProgramItem) {
	fmt.Println("body:")
	pointerPos := 0
	loopCount := 0
	loopStack := make([]int, 0)

	programItem := programItemTop
	for programItem.Next != nil {
		// >, <
		if programItem.Type == parse.ControlPointer {
			pointerPos += programItem.Value
			checkMinusPointer(pointerPos)
			programItem = programItem.Next
			continue
		}

		// +, -
		if programItem.Type == parse.ControlValue {
			fmt.Printf("		add byte ptr [rbp-%d], %d\n", pointerPos*8+8, programItem.Value)
			programItem = programItem.Next
			continue
		}

		// [
		if programItem.Type == parse.LoopStart {
			loopStack = append(loopStack, loopCount)
			loopCount++
			fmt.Printf("		cmp byte ptr [rbp-%d], 0\n", pointerPos*8)
			fmt.Printf("		je __loop_end_%d\n", loopCount-1)
			fmt.Printf("__loop_start_%d:\n", loopCount-1)
			programItem = programItem.Next
			continue
		}

		// ]
		if programItem.Type == parse.LoopEnd {
			loopID := -1
			if len(loopStack) > 0 {
				loopID = loopStack[len(loopStack)-1]
				loopStack = loopStack[:len(loopStack)-1]
			} else {
				common.Error("[, ]の対応が正しくありません")
			}
			fmt.Printf("		cmp byte ptr [rbp-%d], 0\n", pointerPos*8+8)
			fmt.Printf("		jne __loop_start_%d\n", loopID)
			fmt.Printf("__loop_end_%d:\n", loopID)
			programItem = programItem.Next
			continue
		}

		// .
		if programItem.Type == parse.Write {
			fmt.Println("		mov rax, 0x2000004")           // Write
			fmt.Println("		mov rdi, 1")                   // 第1引数 : flides
			fmt.Println("		mov rsi, rbp")                 // 第2引数 : *buf
			fmt.Printf("		sub rsi, %d\n", pointerPos*8+8) // (ポインタ設定)
			fmt.Println("		mov rdx, 1")                   // 第3引数 : nbyte
			fmt.Println("		syscall")
			programItem = programItem.Next
			continue
		}

		// ,
		if programItem.Type == parse.Read {
			fmt.Println("		mov rax, 0x2000003")
			fmt.Println("		mov rdi, 1")
			fmt.Println("		mov rsi, rbp")
			fmt.Printf("		sub rsi, %d\n", pointerPos*8+8)
			fmt.Println("		mov rdx, 1")
			fmt.Println("		syscall")
			programItem = programItem.Next
			continue
		}

		common.Error("不明なエラー")
	}

	// ポインタの値を返り値にする
	fmt.Printf("		movzx rax, byte ptr [rbp-%d]\n", pointerPos*8+8)
}

func checkMinusPointer(pointerPos int) {
	if pointerPos < 0 {
		common.Error("ポインタが負の値を取りました")
	}
}
