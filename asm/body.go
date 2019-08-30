package asm

import (
	"brainfOOk-compiler/common"
	"brainfOOk-compiler/parse"
	"fmt"
)

// Body : ProgramInfoListをアセンブリに変換して出力する
func Body(programItemList *[]parse.ProgramItem) {
	pointerPos := 0

	for _, programItem := range *programItemList {
		// >, <
		if programItem.Type == parse.ControlPointer {
			fmt.Printf("		add rbx, %d\n", programItem.Value)
			pointerPos += programItem.Value
			checkMinusPointer(pointerPos)
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
