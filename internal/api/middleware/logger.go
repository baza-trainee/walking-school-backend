package middleware

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logging(log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now().UTC()

		log.Info(fmt.Sprintf("Time spent: %v, Method %v: %v", time.Since(start), c.Method(), c.OriginalURL()))

		return c.Next()
	}
}
