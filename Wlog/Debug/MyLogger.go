package Wlog

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// LogLevel 日志等级类型
type LogLevel uint8

//日志等级
const (
	UNKOWNLogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// logger 日志结构体
type logger struct {
	level  LogLevel // 日志级别
	writer *os.File // 书写日志位置
	file            // 日志文件
}

var logChan chan []interface{}

// NewLogger 创建一个日志对象(构造函数)
//goland:noinspection GoExportedFuncWithUnexportedType
func newLogger(levelStr string, maxFileSize int64) *logger {
	level, err := parseLogLeve(levelStr) // 将用户输入的字符串转换为我可以认识的level
	if err != nil {
		panic(err)
	}
	logChan = make(chan []interface{}, 128)
	return &logger{
		level:  level,
		writer: os.Stdout,
		file: file{
			path:        "",
			name:        make([]string, 2),
			maxFileSize: maxFileSize,
		},
	}
}

// parseLogLeve 用户输入的日志等级(字符串)转换为日志库认识的LogLevel类型
func parseLogLeve(s string) (LogLevel, error) {
	s = strings.ToLower(s) // 把用户输入的单词全部转换为小写字母
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKOWNLogLevel, err
	}
}

// getInfo 获取代码运行时信息
func getInfo(skip int) (funcName, file string, line int) {
	pc, file /*文件名*/, line /*行号*/, ok := runtime.Caller(skip) // 利用Go GC获取程序运行时代码执行的一些信息
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1] // 获取函数名
	return
	/*
		skip: 函数调用的层级数
		如果是0 则获取当前函数的信息
		如果是1 则获取调用当前函数的那个函数的信息
		...
	*/
}

// writLog 记录日志
func writLog(lv string /*日志等级*/,
	l *logger          /*日志对象*/,
	s ...interface{}   /*日志内容*/) {
	if l.writer != os.Stdout { // 如果日志输出位置不是标准输出(终端) 需要切分日志
		fileInfo, _ := l.writer.Stat() // 获取文件信息
		fileSize := fileInfo.Size()    // 获取当前文件大小
		if fileSize >= l.file.maxFileSize {
			//切分日志
			l.cutLogFileBySize(l.writer)
		}
		fmt.Println("日志文件大小:", fileSize, "B")
	}
	funcName, file, line := getInfo(4)
	_, _ = fmt.Fprintf(l.writer,
		"[time:%s] [LogLevel:%s] [file:%s funcName:%s line:%d]\t",
		time.Now().Format("2006/01/02 03:04:05 PM"), lv, file, funcName, line,
	)
	_, _ = fmt.Fprintln(l.writer, s...)
}
