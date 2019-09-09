package parse

import (
	"brainfOOk-compiler/common"
	"math"
)

// Parse : プログラムを要素ごとに分解する
func Parse(program string) (*ProgramItem, int) {
	programItemTop := ProgramItem{-1, -1, nil}
	pointerPos := 0
	allocPointerNum := 0

	for idx := 0; idx < len(program); idx++ {
		cs := program[idx]

		// >, < (ポインタ位置をインクリメント/デクリメントする)
		if cs == '>' || cs == '<' {
			// 値取り出し
			conLen := common.GetContinueCharLen(program, idx)
			if cs == '<' {
				conLen *= -1
			}
			programItem := ProgramItem{ControlPointer, conLen, &programItemTop}
			programItemTop = programItem

			// ポインタ, メモリチェック
			idx += int(math.Abs(float64(conLen))) - 1
			pointerPos += conLen
			allocPointerNum = int(math.Max(float64(pointerPos), float64(allocPointerNum)))
			continue
		}

		// +, - (ポインタが指す値をインクリメント/デクリメントする)
		if cs == '+' || cs == '-' {
			// 値取り出し
			conLen := common.GetContinueCharLen(program, idx)
			if cs == '-' {
				conLen *= -1
			}
			programItem := ProgramItem{ControlValue, conLen, &programItemTop}
			programItemTop = programItem
			idx += int(math.Abs(float64(conLen))) - 1
			continue
		}

		// [, ] (ループ)
		if cs == '[' || cs == ']' {
			var itemType ItemType
			if cs == '[' {
				itemType = LoopStart
			} else {
				itemType = LoopEnd
			}
			programItem := ProgramItem{itemType, 0, &programItemTop}
			programItemTop = programItem
			continue
		}

		// . (write)
		if cs == '.' {
			programItem := ProgramItem{Write, 0, &programItemTop}
			programItemTop = programItem
			continue
		}

		// . (Read)
		if cs == ',' {
			programItem := ProgramItem{Read, 0, &programItemTop}
			programItemTop = programItem
			continue
		}

		common.ErrorWithPos(program, "実装されていない文字です", idx)
	}

	return &programItemTop, allocPointerNum + 1
}
