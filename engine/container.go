package engine

import (
	"cliabh/engine/render"
	"log"
)

// Container 는 여러 개의 Component 들을 가질 수 있는 컴포넌트입니다.
type Container struct {
	*BaseComponent
	Children []Component
}

func NewContainer(x, y, width, height float32) *Container {
	return &Container{
		BaseComponent: NewBaseComponent(x, y, width, height),
		Children:      []Component{},
	}
}

// AddChild 해당 컨테이너에 자식 컴포넌트를 추가한다.
func (c *Container) AddChild(child Component) {
	c.Children = append(c.Children, child)
}

// Draw 는 컨테이너와 그 안의 모든 자식 컴포넌트를 화면에 그립니다.
func (c *Container) Draw(ctx *render.RenderingContext) {
	for _, child := range c.Children {
		child.Draw(ctx)
	}
}

func (c *Container) InitializeAll() {
	for _, component := range c.Children {
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
