package router

import (
	"fmt"
	"rocket-web/internal/config"

	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Router struct {
	app *fiber.App
	log *zap.SugaredLogger
}

func New(cfg *config.Config, log *zap.SugaredLogger) *Router {
	viewEngine := html.New("./web/templates", ".html")
	switch cfg.Env {
	case config.EnvLocal:
		viewEngine.Reload(true)
	case config.EnvProduction:
		viewEngine.Reload(false)
	}

	app := fiber.New(fiber.Config{
		Views:                 viewEngine,
		DisableStartupMessage: true,
	})

	app.Static("/static", "./web/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world")
	})
	return &Router{app: app, log: log}
}

func (r *Router) MustRun(port int32) error {
	// TODO: Setup routes
	r.log.Infow("starting server", zap.Int32("port", port))
	return r.run(port)
}

func (r *Router) run(port int32) error {
	return r.app.Listen(fmt.Sprintf(":%d", port))
}

func (r *Router) Close() error {
	return r.app.Shutdown()
}
