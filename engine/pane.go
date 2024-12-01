package engine

import (
	"cliabh/engine/render"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

type Pane interface {
	Component
	AddChild(child Component)
	Draw(ctx *render.RenderingContext)
	InitializeAll()
}

type BasePane struct {
	*BaseComponent
	Children []Component
}

func NewBasePane(x, y, width, height float32) *BasePane {
	return &BasePane{
		BaseComponent: NewBaseComponent(x, y, width, height),
		Children:      []Component{},
	}
}

func (bp *BasePane) AddChild(child Component) {
	switch child.(type) {
	case *Window:
		log.Fatalln("Window 는 Pane 에 추가될 수 없습니다.")
		return
	default:
		bp.Children = append(bp.Children, child)
	}
}

func (bp *BasePane) Draw(ctx *render.RenderingContext) {
	for _, child := range bp.Children {
		child.Draw(ctx)
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

func NewGlassPane(x, y, width, height float32, controller EventController) *GlassPane {
	return &GlassPane{
		BasePane:   NewBasePane(x, y, width, height),
		Controller: controller,
	}
}

func (gp *GlassPane) HandleKeyEvent(event glfw.Key) {
	gp.Controller.HandleKeyEvent(event)
}

func (gp *GlassPane) HandleMouseEvent(button glfw.MouseButton, action glfw.Action, xpos, ypos float64) {
	gp.Controller.HandleMouseEvent(button, action, xpos, ypos)
}

func (bp *BasePane) InitializeAll() {
	for _, component := range bp.Children {
		if paneParent, ok := component.(Pane); ok {
			paneParent.InitializeAll()
		} else if initializer, ok := component.(OptionalInitializer); ok {
			err := initializer.Initialize()
			if err != nil {
				log.Fatalf("Initialize Failed. %v", err)
			}
		}
	}
}
