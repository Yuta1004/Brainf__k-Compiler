package parse

// ItemType : ProgramItemの種類を表す
type ItemType int

const (
	// ControlPointer : >, <
	ControlPointer ItemType = iota
	// ControlValue : +, -
	ControlValue
	// LoopStart : [
	LoopStart
	// LoopEnd : ]
	LoopEnd
)
