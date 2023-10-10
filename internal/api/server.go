package api

import (
	"errors"

	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/api/handler"
	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/baza-trainee/walking-school-backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/swagger"
)

type Server struct {
	HTTPServer *fiber.App
	Service    service.Service
	Log        *slog.Logger
}

func NewServer(cfg config.Server, service service.Service, log *slog.Logger) *Server {
	server := new(Server)

	server.Service = service

	server.Log = log

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

	server.initRoutes(server.HTTPServer, cfg)

	return server
}

func (s Server) initRoutes(app *fiber.App, cfg config.Server) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/project", timeout.NewWithContext(handler.CreateProjectHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/project", timeout.NewWithContext(handler.GetAllProjectHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/project/:id", timeout.NewWithContext(handler.GetProjectByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Put("/project", timeout.NewWithContext(handler.UpdateProjectByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Delete("/project/:id", timeout.NewWithContext(handler.DeleteProjectByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))

	app.Post("/user", timeout.NewWithContext(handler.CreateUserHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/user", timeout.NewWithContext(handler.GetAllUserHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/user/:id", timeout.NewWithContext(handler.GetUserByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Put("/user", timeout.NewWithContext(handler.UpdateUserByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Delete("/user/:id", timeout.NewWithContext(handler.DeleteUserByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))

	app.Post("/hero", timeout.NewWithContext(handler.CreateHeroHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/hero", timeout.NewWithContext(handler.GetAllHeroHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/hero/:id", timeout.NewWithContext(handler.GetHeroByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Put("/hero", timeout.NewWithContext(handler.UpdateHeroByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Delete("/hero/:id", timeout.NewWithContext(handler.DeleteHeroByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))

	app.Post("/projects-section-description", timeout.NewWithContext(handler.CreateProjSectDescHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/projects-section-description/:id", timeout.NewWithContext(handler.GetProjSectDescByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Put("/projects-section-description", timeout.NewWithContext(handler.UpdateProjSectDescByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))

	app.Post("/image-carousel", timeout.NewWithContext(handler.CreateImagesCarouselHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/image-carousel", timeout.NewWithContext(handler.GetAllImagesCarouselHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/image-carousel/:id", timeout.NewWithContext(handler.GetImagesCarouselByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Delete("/image-carousel/:id", timeout.NewWithContext(handler.DeleteImagesCarouselByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))

	app.Post("/partner", timeout.NewWithContext(handler.CreatePartnerHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/partner", timeout.NewWithContext(handler.GetAllPartnerHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/partner/:id", timeout.NewWithContext(handler.GetPartnerByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Put("/partner", timeout.NewWithContext(handler.UpdatePartnerByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Delete("/partner/:id", timeout.NewWithContext(handler.DeletePartnerByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))

	app.Post("/contact", timeout.NewWithContext(handler.CreateContactHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Get("/contact", timeout.NewWithContext(handler.GetAllContactHandler(s.Service, s.Log), cfg.AppWriteTimeout))
	app.Put("/contact", timeout.NewWithContext(handler.UpdateContactByIDHandler(s.Service, s.Log), cfg.AppWriteTimeout))
}

func corsConfig() cors.Config {
	return cors.Config{
		// AllowOrigins:     `localhost:3000`, `https://<<<walking-school-backend.com>>>`,
		AllowOrigins:     `*`,
		AllowHeaders:     "Origin, Content-Type, Accept, Access-Control-Allow-Credentials",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}
}
