package asm

import (
	"brainfOOk-compiler/common"
	"brainfOOk-compiler/parse"
	"fmt"
)

// Body : ProgramInfoListをアセンブリに変換して出力する
func Body(programItemList *[]parse.ProgramItem) {
	fmt.Println("body:")
	pointerPos := 0

	for _, programItem := range *checkLoopPair(programItemList) {
		// >, <
		if programItem.Type == parse.ControlPointer {
			pointerPos += programItem.Value
			checkMinusPointer(pointerPos)
			continue
		}

		// +, -
		if programItem.Type == parse.ControlValue {
			fmt.Println("		mov rdx, rbp")
			fmt.Printf("		sub rdx, %d\n", pointerPos*8+8)
			fmt.Printf("		add byte ptr [rdx], %d\n", programItem.Value)
			continue
		}

		// [
		if programItem.Type == parse.LoopStart {
			fmt.Println("		mov rdx, rbp")
			fmt.Printf("		sub rdx, %d\n", pointerPos*8+8)
			fmt.Println("		cmp byte ptr [rdx], 0")
			fmt.Printf("		je __loop_end_%d\n", programItem.Value)
			fmt.Printf("__loop_start_%d:\n", programItem.Value)
			continue
		}

		// ]
		if programItem.Type == parse.LoopEnd {
			fmt.Println("		mov rdx, rbp")
			fmt.Printf("		sub rdx, %d\n", pointerPos*8+8)
			fmt.Println("		cmp byte ptr [rdx], 0")
			fmt.Printf("		jne __loop_start_%d\n", programItem.Value)
			fmt.Printf("__loop_end_%d:\n", programItem.Value)
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
			continue
		}

		common.Error("不明なエラー")
	}

	// ポインタの値を返り値にする
	fmt.Println("		mov rdx, rbp")
	fmt.Printf("		sub rdx, %d\n", pointerPos*8+8)
	fmt.Println("		movzx rax, byte ptr [rdx]")
}

func checkMinusPointer(pointerPos int) {
	if pointerPos < 0 {
		common.Error("ポインタが負の値を取りました")
	}
}

func checkLoopPair(programItemList *[]parse.ProgramItem) *[]parse.ProgramItem {
	loopCount := 0
	loopStack := make([]int, 0)
	newProgramItemList := make([]parse.ProgramItem, 0)

	// [, ]を対応づけていく
	for _, programItem := range *programItemList {
		// Start
		if programItem.Type == parse.LoopStart {
			loopStack = append(loopStack, loopCount)
			programItem := parse.ProgramItem{Type: parse.LoopStart, Value: loopCount}
			newProgramItemList = append(newProgramItemList, programItem)
			loopCount++
			continue
		}
		// End
		if programItem.Type == parse.LoopEnd {
			if len(loopStack) == 0 {
				common.Error("[, ]の対応が正しくありません")
			}
			loopStart := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]
			programItem := parse.ProgramItem{Type: parse.LoopEnd, Value: loopStart}
			newProgramItemList = append(newProgramItemList, programItem)
			continue
		}
		// Others
		newProgramItemList = append(newProgramItemList, programItem)
	}

	if len(loopStack) != 0 {
		common.Error("[, ]の対応が正しくありません")
	}
	return &newProgramItemList
}
