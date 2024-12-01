package engine

import (
	"cliabh/engine/render"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

// Window 는 다수의 Container 를 가지는 형태의 컴포넌트이자
// 하나의 창을 만들고 렌더링하는 주체이다.
type Window struct {
	*Container
	Window *glfw.Window
}

func NewWindow(title string, width, height int) *Window {
	if err := glfw.Init(); err != nil {
		glfw.Terminate()
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

	fmt.Println("OpenGL Version:", gl.GoStr(gl.GetString(gl.VERSION)))
	InitializeContext()

	viewPortWidth, viewPortHeight := win.GetFramebufferSize()
	gl.Viewport(0, 0, int32(viewPortWidth), int32(viewPortHeight))

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	return &Window{
		Container: NewContainer(0, 0, float32(width), float32(height)),
		Window:    win,
	}
}

func (w *Window) MainLoop() {
	w.Container.InitializeAll()
	ctx := render.NewRenderingContext(w.Width, w.Height)

	for !w.Window.ShouldClose() {
		w.Draw(ctx)

		w.Window.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}

func (w *Window) Draw(ctx *render.RenderingContext) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	w.Container.Draw(ctx)
}
