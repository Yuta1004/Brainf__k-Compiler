package common

import (
	"fmt"
	"os"
	"strings"
)

// Error : エラーメッセージを出力する
func Error(msg string) {
	fmt.Fprintf(os.Stderr, msg+"\n")
	os.Exit(1)
}

// ErrorWithPos : エラー箇所をメッセージとともに表示
func ErrorWithPos(base, msg string, pos int) {
	fmt.Fprintf(os.Stderr, base+"\n")
	fmt.Fprintf(os.Stderr, "%s↑ %s\n", strings.Repeat(" ", pos), msg)
	os.Exit(1)
}
