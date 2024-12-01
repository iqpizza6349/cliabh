package engine

import (
	"cliabh/engine/render"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// EventController 는 MVC 패턴에서 컨트롤러 역할을 수행합니다.
type EventController struct {
	Model *render.Model
}

// NewEventController 는 새로운 EventController 를 생성합니다.
func NewEventController(model *render.Model) EventController {
	return EventController{
		Model: model,
	}
}

// HandleMouseEvent 는 컨트롤러가 마우스 이벤트를 처리하는 방법을 정의합니다.
func (ec *EventController) HandleMouseEvent(button glfw.MouseButton, action glfw.Action, xpos, ypos float64) {
	if action == glfw.Press {
		ec.Model.SetLabelText("Mouse Clicked")
	}
}

// HandleKeyEvent 는 컨트롤러가 키보드 이벤트를 처리하는 방법을 정의합니다.
func (ec *EventController) HandleKeyEvent(event glfw.Key) {
	if event == glfw.KeySpace {
		ec.Model.SetLabelText("Space Key Pressed")
	}
}
