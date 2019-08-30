package parse

import (
	"brainfOOk-compiler/common"
)

// Parse : プログラムを要素ごとに分解する
func Parse(program string) *[]ProgramItem {
	programItemList := make([]ProgramItem, 0)

	for idx := 0; idx < len(program); idx++ {
		cs := program[idx]

		// +, - (ポインタ位置をインクリメント/デクリメントする)
		if cs == '+' || cs == '-' {
			conLen := common.GetContinueCharLen(program, idx)
			programItem := ProgramItem{ControlPointer, conLen}
			programItemList = append(programItemList, programItem)
			idx += conLen - 1
			continue
		}

		common.ErrorWithPos(program, "実装されていない文字です", idx)
	}

	return &programItemList
}
