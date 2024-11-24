package main

import (
	"cliabh/engine"
	"log"
	"runtime"
)

func init() {
	// 메인 스레드에서 실행해야 GLFW 및 OpenGL이 정상 작동합니다.
	runtime.LockOSThread()
}

func main() {
	window := engine.NewWindow("My Game Engine", 800, 600)

	// 메인 루프 시작
	log.Println("Starting main loop...")
	window.MainLoop()
}
