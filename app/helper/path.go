package helper

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// 根目录路径
func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	i := strings.LastIndex(s, "\\")
	path := s[0 : i+1]

	return path
}

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}
