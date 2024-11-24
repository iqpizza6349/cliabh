package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

// Window 는 다수의 Container 를 가지는 형태의 컴포넌트이자
// 하나의 창을 만들고 렌더링하는 주체이다.
type Window struct {
	*Container
	Window   *glfw.Window
	RootPane *BasePane
}

func NewWindow(title string, width, height int) *Window {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialized glfw: ", err)
	}

	// OpenGL 버전을 4.1 로 설정
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		log.Fatalln("Failed to create glfw window: ", err)
	}

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialized gl: ", err)
	}

	rootPane := NewBasePane(0, 0, float32(width), float32(height))

	return &Window{
		Container: NewContainer(0, 0, float32(width), float32(height)),
		Window:    win,
		RootPane:  rootPane,
	}
}

func (w *Window) MainLoop() {
	for !w.Window.ShouldClose() {
		w.Update(1.0 / 60.0) // deltaTime 을 임의로 1/60 초로 지정
		w.Draw()

		w.Window.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}

func (w *Window) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	w.RootPane.Draw()
}

// Update 는 윈도우와 그 안의 모든 컴포넌트의 상태를 업데이트합니다.
func (w *Window) Update(deltaTime float64) {
	w.RootPane.Update(deltaTime)
}
