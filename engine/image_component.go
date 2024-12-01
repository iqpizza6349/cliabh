package engine

import (
	"cliabh/engine/render"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type ImageComponent struct {
	*BaseComponent
	textureID uint32
	imagePath string
	// 기타 필드
}

// NewImageComponent 은 텍스트와 위치 및 크기를 가진 라벨을 생성합니다.
func NewImageComponent(x, y, width, height float32, imagePath string) *ImageComponent {
	imageComponent := &ImageComponent{
		BaseComponent: NewBaseComponent(x, y, width, height),
		imagePath:     imagePath,
	}
	//model.RegisterObserver(label) // 모델을 옵저버로 등록
	return imageComponent
}

func (c *ImageComponent) Initialize() error {
	texID, err := GlobalContext.ResourceManager.TextureManager.GetTexture(c.imagePath)
	if err != nil {
		return err
	}
	c.textureID = texID
	// 기타 초기화 코드
	return nil
}

func (c *ImageComponent) Draw(ctx *render.RenderingContext) {
	c.BaseComponent.Draw(ctx)

	gl.ActiveTexture(gl.TEXTURE0)
	// 텍스처 바인딩
	gl.BindTexture(gl.TEXTURE_2D, c.textureID)

	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
}
