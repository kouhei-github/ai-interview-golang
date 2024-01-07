package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kouhei-github/ai-interview-golang/route"
	"os"
)

var fiberLambda *fiberadapter.FiberLambda

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
	if os.Getenv("ENVIRONMENT") == "local" {
		if err := router.FiberApp.Listen(":8080"); err != nil {
			panic(err)
		}
	} else {
		// fmt.Println("lambda")
		// AWS Lambdaとの連携設定
		fiberLambda = fiberadapter.New(router.FiberApp)
		lambda.Start(Handler)
	}
}

// Handler is a function that handles HTTP requests and returns an HTTP response.
// It uses the FiberLambda proxy with context method to pass control to the Lambda function.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}
