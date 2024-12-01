package main

import (
	"cliabh/engine"
	"cliabh/engine/render"
	"log"
	"runtime"
)

func init() {
	// 메인 스레드에서 실행해야 GLFW 및 OpenGL이 정상 작동합니다.
	runtime.LockOSThread()
}

func main() {
	window := engine.NewWindow("My Game Engine", 800, 600)

	model := render.NewModel()

	contentPane := engine.NewBasePane(0, 0, 800, 600)
	//label := engine.NewLabel(100, 100, 200, 50, "Initial Text", model)
	//TODO: open-gl 은 기본적으로 한 가운데가 좌표점이기 때문에 원래 (0, 0)을 표기하려면 크기 / 2만큼 해줘야함
	imageComponent := engine.NewImageComponent(50, 50, 100, 100, "output.png")

	contentPane.AddChild(imageComponent)
	window.AddChild(contentPane)

	// GlassPane 을 Window 에 추가
	glassPane := engine.NewGlassPane(0, 0, 800, 600, engine.NewEventController(model))
	window.AddChild(glassPane)

	// 메인 루프 시작
	log.Println("Starting main loop...")
	window.MainLoop()
}
