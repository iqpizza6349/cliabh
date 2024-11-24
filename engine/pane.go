package engine

import "github.com/go-gl/glfw/v3.3/glfw"

type Pane interface {
	Component
	AddChild(child Pane)
	Draw()
	Update(deltaTime float64)
}

type BasePane struct {
	*BaseComponent
	Children []Pane
}

func NewBasePane(x, y, width, height float32) *BasePane {
	return &BasePane{
		BaseComponent: NewBaseComponent(x, y, width, height),
		Children:      []Pane{},
	}
}

func (bp *BasePane) AddChild(child Pane) {
	bp.Children = append(bp.Children, child)
}

func (bp *BasePane) Draw() {
	for _, child := range bp.Children {
		child.Draw()
	}
}

func (bp *BasePane) Update(deltaTime float64) {
	for _, child := range bp.Children {
		child.Update(deltaTime)
	}
}

// LayerPane 은 여러 컴포넌트를 오버랩하는 기능을 위한 Pane 입니다.
type LayerPane struct {
	*BasePane
}

// NewLayerPane 은 LayerPane 을 생성합니다.
func NewLayerPane(x, y, width, height float32) *LayerPane {
	return &LayerPane{
		BasePane: NewBasePane(x, y, width, height),
	}
}

// ContentPane 은 Layer 에 위치하는 컴포넌트를 포함하는 Pane 입니다.
type ContentPane struct {
	*BasePane
}

// NewContentPane 은 ContentPane 을 생성합니다.
func NewContentPane(x, y, width, height float32) *ContentPane {
	return &ContentPane{
		BasePane: NewBasePane(x, y, width, height),
	}
}

// GlassPane 은 이벤트를 받는 최상단 Pane 입니다.
type GlassPane struct {
	*BasePane
	Controller EventController
}

func NewGlassPane(x, y, width, height float32) *GlassPane {
	return &GlassPane{
		BasePane:   NewBasePane(x, y, width, height),
		Controller: NewEventController(),
	}
}

func (gp *GlassPane) HandleKeyEvent(event glfw.Key) {
	gp.Controller.HandleKeyEvent(event)
}

func (gp *GlassPane) HandleMouseEvent(button glfw.MouseButton, action glfw.Action, xpos, ypos float64) {
	gp.Controller.HandleMouseEvent(button, action, xpos, ypos)
}
