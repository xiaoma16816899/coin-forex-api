package utils

import "github.com/gofiber/fiber/v2"

const (
	X_Real_IP       = "X-Real-IP"
	X_Forwarded_For = "X-Forwarded-For"
)

func GetHeaderPeekByKey(ctx *fiber.Ctx, key string) string {
	if key == "" {
		return ""
	}
	peek := ctx.Context().Request.Header.Peek(key)
	convertPeekToStr := string(peek)
	return convertPeekToStr
}
