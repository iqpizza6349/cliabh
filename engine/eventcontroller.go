package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

// EventController 는 MVC 패턴에서 컨트롤러 역할을 수행합니다.
type EventController struct {
	// 필요한 필드를 추가합니다 (예: 모델 참조, 현재 상태 등)
}

// NewEventController 는 새로운 EventController 를 생성합니다.
func NewEventController() EventController {
	return EventController{}
}

// HandleMouseEvent 는 컨트롤러가 마우스 이벤트를 처리하는 방법을 정의합니다.
func (ec *EventController) HandleMouseEvent(button glfw.MouseButton, action glfw.Action, xpos, ypos float64) {
	// MVC 모델의 동작에 따라 적절히 처리 (예: 모델 업데이트, 뷰 갱신 요청 등)
	log.Printf("Mouse Event: button=%d, action=%d, xpos=%f, ypos=%f", button, action, xpos, ypos)
}

// HandleKeyEvent 는 컨트롤러가 키보드 이벤트를 처리하는 방법을 정의합니다.
func (ec *EventController) HandleKeyEvent(event glfw.Key) {
	// MVC 모델의 동작에 따라 적절히 처리 (예: 모델 업데이트, 뷰 갱신 요청 등)
	log.Printf("Key Event: key=%d", event)
}
