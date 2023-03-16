package utils

import (
	// "fmt"
	"crypto/md5"
	"fmt"
	"os"
	"path"
)

func IsExist(name string) (string, bool) {
	// filepath := "./accounts/" + name + ".json"
	filepath := path.Join("./accounts", fmt.Sprint(name, ".json"))
	_, err := os.Stat(filepath)
	return filepath, !os.IsNotExist(err)
}

func Md5Salt(str string) string {
	salt := "猜猜我是谁"

	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte(salt))
	ret := fmt.Sprintf("%x", h.Sum(nil))
	return ret
}
