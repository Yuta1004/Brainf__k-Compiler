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

	for _, programItem := range *programItemList {
		// >, <
		if programItem.Type == parse.ControlPointer {
			fmt.Printf("		add rbx, %d\n", programItem.Value)
			pointerPos += programItem.Value
			checkMinusPointer(pointerPos)
			continue
		}

		// +, -
		if programItem.Type == parse.ControlValue {
			fmt.Println("		mov rdx, rbp")
			fmt.Printf("		sub rdx, %d\n", pointerPos*8+8)
			fmt.Printf("		mov byte ptr [rdx], %d\n", programItem.Value)
			continue
		}

		common.Error("不明なエラー")
	}
}

func checkMinusPointer(pointerPos int) {
	if pointerPos < 0 {
		common.Error("ポインタが負の値を取りました")
	}
}

func checkLoopPair(programItemList *[]parse.ProgramItem) {
	loopCount := 0
	loopStack := make([]int, 0)

	// [, ]を対応づけていく
	for _, programItem := range *programItemList {
		if programItem.Type == parse.LoopStart {
			programItem.Value = loopCount
			loopStack = append(loopStack, loopCount)
			loopCount++
		}
		if programItem.Type == parse.LoopEnd {
			if len(loopStack) == 0 {
				common.Error("[, ]の対応が正しくありません")
			}
			loopStart := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]
			programItem.Value = loopStart
		}
	}

	if len(loopStack) != 0 {
		common.Error("[, ]の対応が正しくありません")
	}
}
