package readfile

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

/*
使用pkg/errors，可以将原始error封装，并携带额外的信息。
为ReadFile函数中打开文件的错误加上"open failed"的提示。
为ReadFile函数中ReadAll调用可能产生的错误加上"read failed"的提示。
为ReadConfig函数中的错误加上"could not read config"的提示。
*/
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		//TODO: wrap err, add "open failed"
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		//TODO: wrap err, add "read failed"
	}
	return buf, nil
}

func ReadConfig(filename string) ([]byte, error) {
	config, err := ReadFile(filename)
	//TODO: add message for err, "could not read config"
}
