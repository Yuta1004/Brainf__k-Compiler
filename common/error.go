package common

import (
	"fmt"
	"os"
	"strings"
)

// ErrorWithPos : エラー箇所をメッセージとともに表示
func ErrorWithPos(base, msg string, pos int) {
	fmt.Fprintf(os.Stderr, base+"\n")
	fmt.Fprintf(os.Stderr, "%s↑ %s\n", strings.Repeat(" ", pos), msg)
	os.Exit(1)
}
