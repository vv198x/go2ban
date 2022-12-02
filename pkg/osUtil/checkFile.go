package osUtil

import (
	"os"
)

func CheckFile(filePach string) (check bool) {
	_, err := os.Stat(filePach)
	check = os.IsNotExist(err)
	return !check
}
