package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/ai-interview-golang/utils/amazon"
	"os"
)

// VideoS3UploadHandler 関数は、アップロードされたファイルをディスクに保存するPOSTリクエストを処理する関数である。
//
// この関数は、リクエストのコンテキストを表す `fiber.Ctx` パラメータを受け取る。
// S3への保存に失敗した場合はエラーを返し、そうでない場合は `nil` を返す。
//
// リクエストにマルチパートフォームが含まれている場合、この関数はフォームからファイルを抽出してS3に保存する。
// その後、成功メッセージとともに JSON レスポンスを返す。
//
// フォームがマルチパートフォームでない場合、またはエラーが発生した場合、 // この関数はエラーメッセージとともに JSON レスポンスを返す。
//
// 使用例：
// ```
// router.FiberApp.Post("/", controller.HelloWorld)
// ```
// ここで `router` は `Router` 構造体のインスタンスである。
func VideoS3UploadHandler(c *fiber.Ctx) error {
	// s3にアップロード
	sessionAws, err := amazon.NewSessionAWS()
	if err != nil {
		return c.Status(200).JSON(map[string]string{"message": err.Error()})
	}
	// S3クライアントの作成
	s3Service := amazon.NewS3Service(sessionAws, os.Getenv("AWS_BUCKET"))

	// fiberのコンテキストからMultipartFormを取得しています。
	// MultipartFormは "enctype"属性が "multipart/form-data"に設定されているフォームを処理する際に利用されます。
	form, err := c.MultipartForm()

	// MultipartFormの取得時にエラーが発生した場合、サーバーはHTTPステータスコード 200とエラーメッセージを含むJSONをレスポンスします。
	if err != nil {
		return c.Status(200).JSON(map[string]string{"message": err.Error()})
	}

	// MultipartFormから全てのファイルを取得し、ローカル変数 "files"に保存しています。
	files := form.File["file"]

	// files配列内の全てのファイルに対して、
	for _, file := range files {
		// io.Readerを実装したものを取得
		fileReader, err := file.Open()
		if err != nil {
			return c.Status(200).JSON(map[string]string{"message": err.Error()})
		}
		// S3にアップロードする
		if err := s3Service.Upload(file.Filename, fileReader); err != nil {
			return c.Status(200).JSON(map[string]string{"message": err.Error()})
		}
	}

	// 全て正常に実行されたので、関数は HTTP 200ステータスコードとJSONメッセージを返します。
	return c.Status(200).JSON(map[string]string{"message": "success"})
}
