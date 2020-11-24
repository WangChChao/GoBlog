package Wlog

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	B  = 1
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
)

type file struct {
	path        string   // 日志文件路径
	name        []string // 所有日志文件的文件名
	maxFileSize int64    // 日志文件最大占用(单位/B)
}

var (
	// lastTime 上次记日志的时间
	lastTime = (time.Now()).Hour()
)

// SetWriter 根据文件路径设置日志输出位置
func (l *logger) setWriter(name string, isCreate bool) {
	name = time.Now().Format("060102150405") + name
	path, err := GetCurrentPath() // 获取当前可执行文件的路径
	if err != nil {
		fmt.Println(err)
	}
	path += name
	fmt.Println(path)
	if isCreate {
		l.writer, err = os.OpenFile(path,
			os.O_CREATE|os.O_WRONLY,
			0644)
		if err != nil {
			fmt.Println(err)
			panic("文件打开错误!!!")
		}
	} else {
		l.writer, err = os.OpenFile(path,
			os.O_APPEND|os.O_WRONLY,
			0466)
		if err != nil {
			fmt.Println(err)
			panic("文件打开错误!!!")
		}
	}
	l.file.path = path
	l.file.name = append(l.file.name, name)
}

// cutLogFileBySize 按文件大小切分日志
func (l *logger) cutLogFileBySize(f *os.File) {
	fileInfo, _ := f.Stat()     //获取文件的基本信息
	fileName := fileInfo.Name() // 获取旧的文件名
	path, _ := GetCurrentPath() // 获取可执行文件的路径(和日志在同一个目录)
	fileName = fileName[12:]
	fileName = time.Now().Format("060102150405") + fileName
	newFile, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY, 0466) // 打开新文件
	if err != nil {
		fmt.Println("打开新文件失败 err:", err)
	}
	l.writer = newFile
	f.Close() // 关闭旧的文件
}

// CloseFile 关闭日志文件
func (l *logger) closeFile() {
	l.writer.Close()
}

// GetCurrentPath 获取可执行文件的路径 来自Go语言中文网
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return path[0 : i+1], nil
}
