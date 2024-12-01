package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"log"
)

type ResourceManager struct {
	TextureManager *TextureManager
	ShaderProgram  *ShaderProgram
	VertexData     *VertexData
}

var vertexShaderSource = `
#version 410 core
layout(location = 0) in vec3 aPos;
layout(location = 1) in vec2 aTexCoord;

out vec2 TexCoord;

uniform mat4 projection;
uniform mat4 transform;

void main() {
    gl_Position = projection * transform * vec4(aPos, 1.0);
    TexCoord = vec2(aTexCoord.x, 1.0 - aTexCoord.y);
}
`
var fragmentShaderSource = `
#version 410 core
out vec4 FragColor;

in vec2 TexCoord;

uniform sampler2D texture1;

void main() {
	// TexCoord 의 y축을 뒤집음
    vec2 flippedTexCoord = vec2(TexCoord.x, 1.0 - TexCoord.y);
    FragColor = texture(texture1, flippedTexCoord);
}
`

var vertices = []float32{
	// Positions        // Texture Coords
	0.5, 0.5, 0.0, 1.0, 1.0, // 우상단
	0.5, -0.5, 0.0, 1.0, 0.0, // 우하단
	-0.5, -0.5, 0.0, 0.0, 0.0, // 좌하단
	-0.5, 0.5, 0.0, 0.0, 1.0, // 좌상단
}
var indices = []uint32{
	0, 1, 3, // 첫 번째 삼각형
	1, 2, 3, // 두 번째 삼각형
}
var attributes = []VertexAttribute{
	{
		Index:      0, // 버텍스 셰이더의 'aPos' 위치
		Size:       3, // x, y, z
		Type:       gl.FLOAT,
		Normalized: false,
		Offset:     0, // 버텍스 시작점에서 오프셋
	},
	{
		Index:      1, // 버텍스 셰이더의 'aTexCoord' 위치
		Size:       2, // u, v
		Type:       gl.FLOAT,
		Normalized: false,
		Offset:     uintptr(3 * 4), // 위치 데이터 이후 (3개의 float * 4바이트)
	},
}

var stride = int32((3 + 2) * 4) // (위치 요소 + 텍스처 좌표 요소) * float32의 바이트 수

var globalShader *ShaderProgram
var globalVertex *VertexData

func NewResourceManager() *ResourceManager {
	// ShaderProgram 초기화
	if globalShader == nil {
		var err error
		globalShader, err = NewShaderProgram(vertexShaderSource, fragmentShaderSource)
		if err != nil {
			log.Fatalf("Failed to create shader program: %v", err)
		}
	}

	// VertexData 초기화
	if globalVertex == nil {
		var err error
		globalVertex, err = NewVertexData(vertices, indices, attributes, stride)
		if err != nil {
			log.Fatalf("Failed to create global vertex data: %v", err)
		}
	}

	return &ResourceManager{
		TextureManager: NewTextureManager(),
		ShaderProgram:  globalShader, // 미리 생성된 셰이더 프로그램
		VertexData:     globalVertex, // 미리 생성된 버텍스 데이터
	}
}
