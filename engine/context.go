package engine

import "cliabh/engine/render"

type Context struct {
	ResourceManager *render.ResourceManager
	// 기타 전역 상태
}

var GlobalContext = &Context{}

func InitializeContext() {
	GlobalContext.ResourceManager = render.NewResourceManager()
	// 기타 초기화 코드
}
