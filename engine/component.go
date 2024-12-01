package engine

import (
	"cliabh/engine/render"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Component interface {
	Draw(ctx *render.RenderingContext) // 컴포넌트를 렌더링하는 메소드
	SetPosition(x, y float32)          // 위치 설정
	SetSize(width, height float32)     // 크기 설정
}

// OptionalInitializer 필요 시 구현할 수 있는 선택적 인터페이스
type OptionalInitializer interface {
	Initialize() error
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

func (c *BaseComponent) GetTransformMatrix() mgl32.Mat4 {
	// 위치 변환
	translation := mgl32.Translate3D(c.X, c.Y, 0)

	// 크기 변환
	scaling := mgl32.Scale3D(c.Width, c.Height, 1)

	// 최종 변환 행렬
	return translation.Mul4(scaling)
}

func (c *BaseComponent) Draw(ctx *render.RenderingContext) {
	// 세이더 사용
	GlobalContext.ResourceManager.ShaderProgram.Use()

	ctx.ApplyProjection(GlobalContext.ResourceManager.ShaderProgram.ID)

	transform := c.GetTransformMatrix()

	// 셰이더에서 'transform' 유니폼 설정
	transformLoc := gl.GetUniformLocation(GlobalContext.ResourceManager.ShaderProgram.ID, gl.Str("transform\x00"))
	gl.UniformMatrix4fv(transformLoc, 1, false, &transform[0])

	// 렌더링 수행
	gl.BindVertexArray(GlobalContext.ResourceManager.VertexData.VAO)
}

//
//type Label struct {
//	*BaseComponent
//	Text  string
//	Model *render.Model
//}
//
//// NewLabel 은 텍스트와 위치 및 크기를 가진 라벨을 생성합니다.
//func NewLabel(x, y, width, height float32, text string, model *render.Model) *Label {
//	label := &Label{
//		BaseComponent: NewBaseComponent(x, y, width, height),
//		Text:          model.LabelText,
//		Model:         model,
//	}
//	model.RegisterObserver(label) // 모델을 옵저버로 등록
//	return label
//}
//
//func (l *Label) CreateTextImage() *image.RGBA {
//	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
//	// 배경을 흰색으로 설정
//	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.Black}, image.Point{}, draw.Src)
//
//	face := basicfont.Face7x13
//	d := &font.Drawer{
//		Dst:  img,
//		Src:  image.White,
//		Face: face,
//		Dot:  fixed.P(20, 50),
//	}
//	d.DrawString(l.Text)
//
//	saveImage("output.png", img)
//	return img
//}
//
//// saveImage 는 이미지를 파일로 저장합니다 (디버깅용).
//func saveImage(filename string, img image.Image) {
//	file, err := os.Create(filename)
//	if err != nil {
//		log.Fatalf("failed to create image file: %v", err)
//	}
//	defer file.Close()
//
//	if err := png.Encode(file, img); err != nil {
//		log.Fatalf("failed to encode image: %v", err)
//	}
//}
//
//// Draw 는 Label 을 화면에 그리는 역할을 합니다.
//func (l *Label) Draw() {
//
//}
//
//// Update 는 Model 의 변경 사항을 반영하여 Label 를 업데이트합니다.
//// Observer 인터페이스의 메서드입니다.
//func (l *Label) Update() {
//	// Mode l의 LabelText 값을 가져와서 자신의 텍스트를 업데이트
//	l.Text = l.Model.LabelText
//}
