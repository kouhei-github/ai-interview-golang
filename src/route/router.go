package route

import (
	"github.com/kouhei-github/ai-interview-golang/controller"
	"github.com/kouhei-github/ai-interview-golang/middlewares"
)

func (router *Router) GetRouter() {
	// 練習
	router.FiberApp.Post("/", middlewares.CustomHeaderIncludeHandler, controller.VideoS3UploadHandler)
}
