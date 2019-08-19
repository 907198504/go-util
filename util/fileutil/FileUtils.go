package fileutil

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
按行读取文件，返回list
*/
func ReadLineAsSlice(filePath string, spiltChar string) []string {
	file, err := ioutil.ReadFile(filePath)
	// 处理异常
	if err != nil {
		panic(err)
	}
	str := string(file)
	// 分割字符串
	return strings.Split(str, spiltChar)
}

/**
向文件后追加内容
*/
func AppendToFile(filePath string, content []byte) {
	// 追加内容
	err := ioutil.WriteFile(filePath, content, os.ModeAppend)
	//
	if err != nil {
		panic(err)
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
