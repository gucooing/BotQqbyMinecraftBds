package server

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gucooing/BotQqbyMinecraftBds/logger"
)

// Command 监听控制台输入，并将结果传入服务器
func (s *Server) Command() {
	scanner := bufio.NewScanner(os.Stdin)
	// 开启一个goroutine监听控制台输入
	for scanner.Scan() {
		command := scanner.Text()
		s.SendCommand(command)
	}
}

// Sender 第三方命令传入通道
func Sender(msg string) {
	SERVER.SendCommand(msg)
}

// SendCommand 发送命令到服务器
func (s *Server) SendCommand(command string) {
	_, err := fmt.Fprintln(s.Stdin, command)
	if err != nil {
		logger.Debug("无法输入: %s", err.Error())
	}
}

func (s *Server) ServerStdout() {
	scanner := bufio.NewScanner(s.Stdout)
	for scanner.Scan() {
		logger.Info(scanner.Text())
	}
}
