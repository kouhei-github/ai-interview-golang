package route

import (
	"github.com/kouhei-github/ai-interview-golang/controller"
)

func (router *Router) GetRouter() {
	// 練習
	router.FiberApp.Post("/", controller.VideoS3UploadHandler)
}
