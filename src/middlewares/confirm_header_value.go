package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func CustomHeaderIncludeHandler(c *fiber.Ctx) error {
	companyHeader := c.Get("Company-Value")
	commonHeader := c.Get("Common-Value")

	if companyHeader != os.Getenv("COMPANY_HEADER") || commonHeader != os.Getenv("COMMON_HEADER") {
		return c.Status(fiber.StatusForbidden).JSON(map[string]string{"message": "Incorrect certification."})
	}
	return c.Next()
}
