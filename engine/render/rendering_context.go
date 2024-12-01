package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type RenderingContext struct {
	ProjectionMatrix mgl32.Mat4
	ScreenWidth      float32
	ScreenHeight     float32
}

// NewRenderingContext 렌더링 컨텍스트 생성
func NewRenderingContext(screenWidth, screenHeight float32) *RenderingContext {
	return &RenderingContext{
		ProjectionMatrix: mgl32.Ortho(0, screenWidth, screenHeight, 0, -1.0, 1.0),
		ScreenWidth:      screenWidth,
		ScreenHeight:     screenHeight,
	}
}

// ApplyProjection 투영 행렬을 셰이더로 전달
func (ctx *RenderingContext) ApplyProjection(shaderProgram uint32) {
	projectionLoc := gl.GetUniformLocation(shaderProgram, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionLoc, 1, false, &ctx.ProjectionMatrix[0])
}
