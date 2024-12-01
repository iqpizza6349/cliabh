package render

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type ShaderProgram struct {
	ID uint32
}

// NewShaderProgram 함수는 버텍스 셰이더와 프래그먼트 셰이더 소스를 받아 컴파일하고 프로그램을 생성합니다.
func NewShaderProgram(vertexSrc, fragmentSrc string) (*ShaderProgram, error) {
	// 버텍스 셰이더 컴파일
	vertexShader, err := compileShader(vertexSrc, gl.VERTEX_SHADER)
	if err != nil {
		return nil, fmt.Errorf("vertex shader compilation failed: %v", err)
	}

	// 프래그먼트 셰이더 컴파일
	fragmentShader, err := compileShader(fragmentSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, fmt.Errorf("fragment shader compilation failed: %v", err)
	}

	// 셰이더 프로그램 생성 및 링크
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	// 링크 성공 여부 확인
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	// 컴파일된 셰이더 삭제 (프로그램에 링크되었으므로)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &ShaderProgram{ID: program}, nil
}

// compileShader 함수는 셰이더 소스를 받아 지정된 타입의 셰이더를 컴파일합니다.
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	compiledShader, free := gl.Strs(source + "\x00")
	defer free()
	gl.ShaderSource(shader, 1, compiledShader, nil)
	gl.CompileShader(shader)

	// 컴파일 성공 여부 확인
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		shaderTypeStr := "UNKNOWN"
		if shaderType == gl.VERTEX_SHADER {
			shaderTypeStr = "VERTEX"
		} else if shaderType == gl.FRAGMENT_SHADER {
			shaderTypeStr = "FRAGMENT"
		}

		return 0, fmt.Errorf("%s shader compilation failed: %v", shaderTypeStr, log)
	}

	return shader, nil
}

func (sp *ShaderProgram) Use() {
	gl.UseProgram(sp.ID)
}

type VertexAttribute struct {
	Index      uint32  // 셰이더에서의 위치 (layout location)
	Size       int32   // 요소의 개수 (예: vec3이면 3)
	Type       uint32  // 데이터 타입 (예: gl.FLOAT)
	Normalized bool    // 정규화 여부
	Offset     uintptr // 버텍스 시작점으로부터의 바이트 오프셋
}

type VertexData struct {
	VAO        uint32
	VBO        uint32
	EBO        uint32
	IndexCount int32
}

func NewVertexData(vertices []float32, indices []uint32, attributes []VertexAttribute, stride int32) (*VertexData, error) {
	if len(attributes) == 0 {
		return nil, fmt.Errorf("attributes cannot be empty")
	}
	if stride <= 0 {
		return nil, fmt.Errorf("stride must be positive")
	}

	// 속성들의 유효성 검사
	for _, attr := range attributes {
		if attr.Size <= 0 {
			return nil, fmt.Errorf("attribute size must be positive")
		}
		if attr.Type != gl.FLOAT && attr.Type != gl.INT && attr.Type != gl.UNSIGNED_INT {
			return nil, fmt.Errorf("unsupported attribute type: %v", attr.Type)
		}
	}

	var vao, vbo, ebo uint32

	// VAO 생성 및 바인딩
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// VBO 생성 및 바인딩
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// 인덱스가 있는 경우 EBO 생성 및 바인딩
	indexCount := int32(0)
	if len(indices) > 0 {
		gl.GenBuffers(1, &ebo)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
		indexCount = int32(len(indices))
	}

	// 버텍스 속성 설정
	for _, attr := range attributes {
		gl.EnableVertexAttribArray(attr.Index)
		gl.VertexAttribPointer(attr.Index, attr.Size, attr.Type, attr.Normalized, stride, gl.PtrOffset(int(attr.Offset)))
	}

	// VAO 와 VBO 바인딩 해제 (EBO 는 VAO 에 바인딩된 상태로 남음)
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return &VertexData{
		VAO:        vao,
		VBO:        vbo,
		EBO:        ebo,
		IndexCount: indexCount,
	}, nil
}
