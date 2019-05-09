package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xiaozefeng/gbdc/command"
)

func main() {
	path := os.Args[1]
	if path == "" {
		log.Printf("请输入你要运行的java文件\n")
		return
	}
	if !strings.HasSuffix(path, ".java") {
		log.Printf("不是Java源文件，请选择正确的文件\n")
		return
	}
	index := strings.LastIndex(path, `.`)
	className := (path)[:index]
	// command.Run("javac " + dir + "/" + *path)
	command.Run("javac " + path)
	result, err := command.Run("java " + className)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	command.Run("rm " + className + ".class")
}
