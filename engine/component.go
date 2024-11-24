package engine

type Component interface {
	Draw()                         // 컴포넌트를 렌더링하는 메소드
	Update(deltaTime float64)      // 컴포넌트 상태를 업데이트하는 메소드
	SetPosition(x, y float32)      // 위치 설정
	SetSize(width, height float32) // 크기 설정
}

// BaseComponent 기본적인 컴포넌트
type BaseComponent struct {
	X, Y          float32
	Width, Height float32
}

func NewBaseComponent(x, y, width, height float32) *BaseComponent {
	return &BaseComponent{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (c *BaseComponent) SetPosition(x, y float32) {
	c.X = x
	c.Y = y
}

func (c *BaseComponent) SetSize(width, height float32) {
	c.Width = width
	c.Height = height
}
