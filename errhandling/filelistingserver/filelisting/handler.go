package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

// panic
/*
1. 停止当前函数执行
2. 一直向上返回，执行每一层的defer
3. 如果没有遇见recover，程序退出
*/

// recover
/*
1. 仅在defer调用中使用
2. 获取panic的值
3. 如果无法处理，可重新panic
*/

// error vs panic
/*
1. 意料之中的：使用error。如：文件打不开
2. 意料之外的：使用panic。如：数组越界
*/

// 错误处理综合示例
/*
1. defer + panic + recover
2. Type Assertion
3. 函数式编程的应用
*/
