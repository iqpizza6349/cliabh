package render

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TextureManager struct {
	textures map[string]uint32
}

func NewTextureManager() *TextureManager {
	return &TextureManager{
		textures: make(map[string]uint32),
	}
}

func (tm *TextureManager) GetTexture(path string) (uint32, error) {
	if tex, ok := tm.textures[path]; ok {
		return tex, nil // 이미 로드된 텍스처 반환
	}

	// 텍스처 로드 및 생성
	tex, err := loadTexture(path)
	if err != nil {
		return 0, err
	}

	tm.textures[path] = tex
	return tex, nil
}

func loadTexture(path string) (uint32, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer func(imgFile *os.File) {
		err := imgFile.Close()
		if err != nil {
			log.Fatalln("cannot close image file", err)
		}
	}(imgFile)

	img, _, err := image.Decode(imgFile)
	if err != nil {
		if strings.Contains(err.Error(), "format") {
			ext := strings.TrimPrefix(filepath.Ext(path), ".")
			if ext == "png" {
				img, err = png.Decode(imgFile)
				if err != nil {
					return 0, err
				}
			} else if ext == "gif" {
				img, err = gif.Decode(imgFile)
				if err != nil {
					return 0, err
				}
			}

		} else {
			return 0, err
		}
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		log.Println("Warning: Non-standard stride detected, but continuing")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	// 텍스처 파라미터 설정
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	gl.GenerateMipmap(gl.TEXTURE_2D)
	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatalf("OpenGL Error during mipmap generation: %v", err)
	}

	// OpenGL 에러 확인
	if err := gl.GetError(); err != gl.NO_ERROR {
		return 0, fmt.Errorf("OpenGL error during texture upload: %v", err)
	}

	// 텍스처 바인딩 해제
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return texture, nil
}
