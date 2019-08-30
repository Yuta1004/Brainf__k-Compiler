package common

import (
	"io/ioutil"
	"os"
)

// ReadFile : ファイルを読み込んでその内容を返す
func ReadFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		Error("ファイル開封エラー")
	}
	defer file.Close()

	result, err := ioutil.ReadAll(file)
	if err != nil {
		Error("ファイル読み込みエラー")
	}
	return string(result)
}
