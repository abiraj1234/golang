package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

const format = "%v: info: %s\n"

type logger interface {
	info(string)
}

type ScreenLogger struct{}

func (ScreenLogger) info(message string) {
	fmt.Printf(format, time.Now(), message)
}

type FileLogger struct {
	File *os.File
}

func (l *FileLogger) info(message string) {
	fmt.Fprint(l.File, format, time.Now(), message)
}

func NewFileLogger(filename string) *FileLogger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	return &FileLogger{file}
}

func main() {
	var log logger
	log = NewFileLogger("log.txt")
	log.info("Hi Donald!")
	fmt.Println(reflect.TypeOf(log))
}
