package parse

import (
	"brainfOOk-compiler/common"
	"fmt"
	"math"
	"os"
)

// Parse : プログラムを要素ごとに分解する
func Parse(program string) (*ProgramItem, int) {
	programItemEnd := &ProgramItem{}
	programItemTop := programItemEnd
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
			programItem := ProgramItem{ControlPointer, conLen, nil}
			programItemEnd.Next = &programItem
			programItemEnd = &programItem

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
			programItem := ProgramItem{ControlValue, conLen, nil}
			programItemEnd.Next = &programItem
			programItemEnd = &programItem
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
			programItem := ProgramItem{itemType, 0, nil}
			programItemEnd.Next = &programItem
			programItemEnd = &programItem
			continue
		}

		// . (write)
		if cs == '.' {
			programItem := ProgramItem{Write, 0, nil}
			programItemEnd.Next = &programItem
			programItemEnd = &programItem
			continue
		}

		// . (Read)
		if cs == ',' {
			programItem := ProgramItem{Read, 0, nil}
			programItemEnd.Next = &programItem
			fmt.Fprintf(os.Stderr, "%v ", programItemTop)
			programItemEnd = &programItem
			fmt.Fprintf(os.Stderr, "%v\n", programItemTop)
			continue
		}

		common.ErrorWithPos(program, "実装されていない文字です", idx)
	}

	programItemEnd.Next = &ProgramItem{}
	return programItemTop.Next, allocPointerNum + 1
}
