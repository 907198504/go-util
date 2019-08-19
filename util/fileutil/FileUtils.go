package fileutil

import (
	"bufio"
	"io"
	"os"
)

// 换行符
const fileLineSplitString byte = '\n'

/**
读取文件
*/
func ReadFileLine(filePath string) {
	// open
	f, err := os.Open(filePath)
	// err
	if err != nil {
		panic(err)
	}
	// 关闭异常
	CloseFile(f)
	// 读取文件
	rd := bufio.NewReader(f)
	for {
		// 以'\n'为结束符读入一行
		_, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
	}
}

/**
关闭文件
*/
func CloseFile(file *os.File) {
	if file == nil {
		return
	}
	// 尝试关闭
	e := file.Close()
	// 关闭失败
	if e != nil {
		panic(e)
	}
}
