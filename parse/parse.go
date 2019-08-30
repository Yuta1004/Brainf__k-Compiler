package parse

import (
	"brainfOOk-compiler/common"
	"math"
)

// Parse : プログラムを要素ごとに分解する
func Parse(program string) *[]ProgramItem {
	programItemList := make([]ProgramItem, 0)

	for idx := 0; idx < len(program); idx++ {
		cs := program[idx]

		// >, < (ポインタ位置をインクリメント/デクリメントする)
		if cs == '>' || cs == '<' {
			conLen := common.GetContinueCharLen(program, idx)
			if cs == '<' {
				conLen *= -1
			}
			programItem := ProgramItem{ControlPointer, conLen}
			programItemList = append(programItemList, programItem)
			idx += int(math.Abs(float64(conLen))) - 1
			continue
		}

		common.ErrorWithPos(program, "実装されていない文字です", idx)
	}

	return &programItemList
}
