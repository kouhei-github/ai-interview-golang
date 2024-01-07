package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

// CustomHeaderIncludeHandler はリクエストのcompanyヘッダとcommonヘッダを検証します。
// ヘッダが正しくない場合、JSON メッセージとともに 403 Forbidden レスポンスを返します。
// ヘッダが正しい場合、次のハンドラを呼び出します。
func CustomHeaderIncludeHandler(c *fiber.Ctx) error {
	companyHeader := c.Get("Company-Value")
	commonHeader := c.Get("Common-Value")

	if companyHeader != os.Getenv("COMPANY_HEADER") || commonHeader != os.Getenv("COMMON_HEADER") {
		return c.Status(fiber.StatusForbidden).JSON(map[string]string{"message": "Incorrect certification."})
	}
	return c.Next()
}
