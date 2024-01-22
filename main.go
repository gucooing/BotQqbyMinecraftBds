package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gucooing/BotQqbyMinecraftBds/bot"
	"github.com/gucooing/BotQqbyMinecraftBds/config"
	"github.com/gucooing/BotQqbyMinecraftBds/logger"
	"github.com/gucooing/BotQqbyMinecraftBds/server"
)

func main() {
	// 启动读取配置
	confName := "BotQqbyMinecraftBds.json"
	err := config.LoadConfig(confName)
	if err != nil {
		if err == config.FileNotExist {
			p, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
			cf, _ := os.Create("./" + confName)
			cf.Write(p)
			cf.Close()
			fmt.Printf("找不到配置文件\n已生成默认配置文件 %s \n", confName)
			main()
		} else {
			panic(err)
		}
	}
	// 初始化日志
	logger.InitLogger()
	logger.SetLogLevel(strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("BotQqbyMinecraftBds")

	conf := config.GetConfig()
	// 初始化
	newbot := bot.NewBot(conf)
	if newbot == nil {
		logger.Error("服务器初始化失败")
		return
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 启动游戏
	go func() {
		for {
			// 启动服务器
			// 创建一个带有取消机制的上下文（context）
			_, cancel := context.WithCancel(context.Background())
			programPath := config.GetConfig().ServerPath
			bdsServer := server.NewServer(programPath)
			if bdsServer == nil {
				logger.Error("服务器启动失败: %v\n", err)
				continue
			}
			defer func(cmd *exec.Cmd) {
				// 停止服务器进程
				err := cmd.Process.Kill()
				if err != nil {
					fmt.Println("无法终止服务器进程:", err)
				}
			}(bdsServer.Cmd)
			logger.Info("\nServer 启动！")
			// 启动控制台监控
			go bdsServer.Command()
			go bdsServer.ServerStdout()
			err = bdsServer.Cmd.Wait()
			// 函数退出，触发进程守护重启服务器
			logger.Info("发送停止信息...")
			cancel()
			logger.Info("2秒后重启服务器...")
			time.Sleep(2 * time.Second)
		}
	}()

	// 启动BOT
	go func() {
		if err = newbot.Run(); err != nil {
			logger.Error("无法启动BotQqbyMinecraftBds服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}
}
