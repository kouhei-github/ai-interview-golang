package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/ai-interview-golang/utils/authorization"
	"strings"
)

// CheckJwtToken はリクエストヘッダから JWT トークンを検証し、コンテキストにユーザー ID を設定します。
// トークンが無効または期限切れの場合はエラーを返します。
func CheckJwtToken(c *fiber.Ctx) error {
	// リクエストヘッダーからAuthorizationの取得
	bearerToken := c.Get("Authorization")
	// Bearer の文字を削除
	bearerToken = strings.ReplaceAll(bearerToken, "Bearer ", "")
	// Bearerトークンからemailの取り出し
	userId, err := authorization.IsTokenExpired(bearerToken)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
	}

	// LocalsメソッドでMiddlewareで取得した値をコンテキストに設定できる
	c.Locals("myUserId", userId)

	// LocalsメソッドでMiddlewareで取得した値をコンテキストから取得する
	// token := c.Locals("bearer") // これで取得できる
	return c.Next()
}
