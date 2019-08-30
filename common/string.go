package common

// GetContinueCharLen : baseのstart番目から連続する文字の数を数えて返す
func GetContinueCharLen(base string, start int) int {
	count := 0
	target := base[start]
	for base[start+count] == target {
		count++
	}
	return count
}
