package parse

// ProgramItem : プログラムを要素ごとに分解した時の要素を表す
type ProgramItem struct {
	Type  ItemType
	Value int
	Next  *ProgramItem
}
