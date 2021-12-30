package File

import (
	"os"
)

func CreateIfNotExistDir(fileAddr string) bool {

	_, err := os.Stat(fileAddr)
	success := true
	if err == nil {
		fileErr := os.Mkdir(fileAddr, 755)

		if fileErr != nil {
			success = false
			//todo::panic 错误是不是不能抛需要用error去接收？
			panic(fileErr)
		}
	}

	return success
}
