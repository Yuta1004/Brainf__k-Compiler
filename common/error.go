package common

import (
	"fmt"
	"os"
	"strings"
)

// ErrorWithPos : エラー箇所をメッセージとともに表示
func ErrorWithPos(base, msg string, pos int) {
	fmt.Println(base)
	fmt.Println(strings.Repeat(" ", pos) + " : " + msg)
	os.Exit(1)
}
