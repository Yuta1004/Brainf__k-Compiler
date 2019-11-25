package asm

import (
	"brainfOOk-compiler/common"
	c "brainfOOk-compiler/common"
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
			c.PrintAsm("add byte ptr [rbp-%d], %d", pointerPos*8+8, programItem.Value)
			programItem = programItem.Next
			continue
		}

		// [
		if programItem.Type == parse.LoopStart {
			loopStack = append(loopStack, loopCount)
			loopCount++
			c.PrintAsm("cmp byte ptr [rbp-%d], 0", pointerPos*8)
			c.PrintAsm("je .L__loop_end_%d", loopCount-1)
			c.PrintLabel("loop_start_%d", loopCount-1)
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
			c.PrintAsm("cmp byte ptr [rbp-%d], 0", pointerPos*8+8)
			c.PrintAsm("jne .L__loop_start_%d", loopID)
			c.PrintLabel("loop_end_%d", loopID)
			programItem = programItem.Next
			continue
		}

		// .
		if programItem.Type == parse.Write {
			c.PrintAsm("mov rax, 0x2000004")          // Write
			c.PrintAsm("mov rdi, 1")                  // 第1引数 : flides
			c.PrintAsm("mov rsi, rbp")                // 第2引数 : *buf
			c.PrintAsm("sub rsi, %d", pointerPos*8+8) // (ポインタ設定)
			c.PrintAsm("mov rdx, 1")                  // 第3引数 : nbyte
			c.PrintAsm("syscall")
			programItem = programItem.Next
			continue
		}

		// ,
		if programItem.Type == parse.Read {
			c.PrintAsm("mov rax, 0x2000003")
			c.PrintAsm("mov rdi, 1")
			c.PrintAsm("mov rsi, rbp")
			c.PrintAsm("sub rsi, %d", pointerPos*8+8)
			c.PrintAsm("mov rdx, 1")
			c.PrintAsm("syscall")
			programItem = programItem.Next
			continue
		}

		common.Error("不明なエラー")
	}

	// ポインタの値を返り値にする
	c.PrintAsm("movzx rax, byte ptr [rbp-%d]", pointerPos*8+8)
}

func checkMinusPointer(pointerPos int) {
	if pointerPos < 0 {
		common.Error("ポインタが負の値を取りました")
	}
}
