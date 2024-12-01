## Clibah

스코틀랜드 게일어로 "나무 상자"라는 의미를 가지고 있으며,
간단한 2D 게임 엔진을 만드는 것을 목표로 두고 있다.

## Vertex 사용 예시
```go
// 버텍스 데이터 (예: 위치와 텍스처 좌표)
vertices := []float32{
    // Positions       // Texture Coords
     0.5,  0.5, 0.0,    1.0, 1.0,
     0.5, -0.5, 0.0,    1.0, 0.0,
    -0.5, -0.5, 0.0,    0.0, 0.0,
    -0.5,  0.5, 0.0,    0.0, 1.0,
}

indices := []uint32{
    0, 1, 3,
    1, 2, 3,
}

// 버텍스 속성 정의
attributes := []VertexAttribute{
    {
        Index:      0,
        Size:       3,
        Type:       gl.FLOAT,
        Normalized: false,
        Offset:     0,
    },
    {
        Index:      1,
        Size:       2,
        Type:       gl.FLOAT,
        Normalized: false,
        Offset:     uintptr(3 * 4), // 위치 데이터 이후에 텍스처 좌표가 위치함
    },
}

stride := int32((3 + 2) * 4) // 버텍스 하나당 바이트 수 (float32는 4바이트)

vertexData, err := NewVertexData(vertices, indices, attributes, stride)
if err != nil {
    log.Fatalf("Failed to create VertexData: %v", err)
}
```

### Rendering example
```go
// 셰이더 프로그램 사용
shaderProgram.Use()

// 필요한 유니폼 설정
// 예: shaderProgram.SetInt("texture1", 0)

// VAO 바인딩
gl.BindVertexArray(vertexData.VAO)

// 드로우 콜
if vertexData.IndexCount > 0 {
    gl.DrawElements(gl.TRIANGLES, vertexData.IndexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
} else {
    vertexCount := int32(len(vertices)) / (stride / 4)
    gl.DrawArrays(gl.TRIANGLES, 0, vertexCount)
}

// VAO 바인딩 해제 (선택 사항)
gl.BindVertexArray(0)
```

### Resource Free
```go
func (v *VertexData) Delete() {
    if v.VAO != 0 {
        gl.DeleteVertexArrays(1, &v.VAO)
    }
    if v.VBO != 0 {
        gl.DeleteBuffers(1, &v.VBO)
    }
    if v.EBO != 0 {
        gl.DeleteBuffers(1, &v.EBO)
    }
}
```
