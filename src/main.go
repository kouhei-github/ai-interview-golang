package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kouhei-github/ai-interview-golang/route"
)

const MaxBodySize = 500 * 1024 * 1024

// main はアプリケーションのエントリー・ポイントです。
// Fiberアプリを設定し、ルーターをセットアップする。
// また、CORSを設定し、Webサーバーを起動するか、環境に応じてAWS Lambdaに接続します。
// ローカル環境では8080番ポートをリッスンし、そうでない場合はAWS Lambdaインテグレーションを使用する。
// ルーターは、さまざまなルートとそれに対応するハンドラの処理を担当する。
// ルーターは、許可されるヘッダーとオリジンを指定するCORS設定で構成される。
// ルートはLoadRouter関数を使用してロードされ、それに応じてWebサーバーまたはLambdaが起動します。
func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: MaxBodySize,
	})

	// ルーターの設定
	router := &route.Router{FiberApp: app}

	// CORS (Cross Origin Resource Sharing)の設定
	// アクセスを許可するドメイン等を設定します
	router.FiberApp.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Company-Value, Common-Value",
		AllowOrigins: "*",
	}))

	route.LoadRouter(router)

	// Webサーバー起動時のエラーハンドリング => localhostの時コメントイン必要
	if err := router.FiberApp.Listen(":80"); err != nil {
		panic(err)
	}
}
