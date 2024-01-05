package route

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	FiberApp *fiber.App
}

// RouterImp は、GetRouter() および GetAuthRouter() メソッドを定義するインターフェイスです。
//
// 使用例：
// func LoadRouter(router RouterImp) { // router.GetRouter()
//
//		router.GetRouter()
//		router.GetAuthRouter()
//	}
type RouterImp interface {
	GetRouter()
	GetAuthRouter()
}

// LoadRouter 与えられたRouterImpからルーターと認証ルーターをロードする。
func LoadRouter(router RouterImp) {
	router.GetRouter()
	router.GetAuthRouter()
}
