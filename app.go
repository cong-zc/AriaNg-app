package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	开启aria2()
}

func (a *App) shutdown(ctx context.Context) {
	a.ctx = ctx
	关闭aria2()
}

// Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }

func 开启aria2() {
	// 获取当前目录
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败")
		return
	}
	fmt.Println("当前目录:", cwd)
	// cmd命令
	var cmd = exec.Command("cmd", "/C", "aria2c --conf-path="+cwd+"/aria2/aria2.conf")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
}

func 关闭aria2() {
	// 结束aria2进程
	cmd := exec.Command("cmd", "/C", "taskkill /f /t /im aria2c.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
}
