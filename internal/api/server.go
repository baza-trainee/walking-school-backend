package api

import (
	"errors"

	"github.com/baza-trainee/walking-school-backend/internal/api/handler"
	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type Server struct {
	HTTPServer *fiber.App
}

func NewServer(cfg config.Server, h handler.Handler) *Server {
	server := new(Server)

	fconfig := fiber.Config{
		ReadTimeout:  cfg.AppReadTimeout,
		WriteTimeout: cfg.AppWriteTimeout,
		IdleTimeout:  cfg.AppIdleTimeout,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := fiber.ErrInternalServerError.Message

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}

			ctx.Status(code)

			if err := ctx.JSON(model.SetResponse(code, message)); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}

			return nil
		},
	}

	server.HTTPServer = fiber.New(fconfig)

	server.HTTPServer.Use(cors.New(corsConfig()))

	server.HTTPServer.Use(recover.New())

	server.initRoutes(server.HTTPServer, cfg, h)

	return server
}

func (s Server) initRoutes(app *fiber.App, cfg config.Server, h handler.Handler) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello my WORLD!!!")
	})
}

func corsConfig() cors.Config {
	return cors.Config{
		// AllowOrigins:     `https://walking-school-backend.com`,
		AllowOrigins:     `*`,
		AllowHeaders:     "Origin, Content-Type, Accept, Access-Control-Allow-Credentials",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}
}
