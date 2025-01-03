package router

import (
	"fmt"
	"rocket-web/internal/config"
	"rocket-web/internal/router/users"
	"rocket-web/internal/storage/postgres"

	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Router struct {
	app     *fiber.App
	log     *zap.SugaredLogger
	storage *postgres.Storage

	// Services
	users *users.Service
}

func New(
	cfg *config.Config,
	storage *postgres.Storage,
	log *zap.SugaredLogger,
) *Router {
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

	app.Static("/static", "./web/static")

	users := users.New(storage, log)

	return &Router{
		app:     app,
		log:     log,
		storage: storage,
		users:   users,
	}
}

func (r *Router) MustRun(port int32) error {
	// TODO: Setup routes
	r.log.Infow("starting server", zap.Int32("port", port))

	r.log.Info("setup account routes")
	r.setupAccountRoutes()

	return r.run(port)
}

func (r *Router) run(port int32) error {
	return r.app.Listen(fmt.Sprintf(":%d", port))
}

func (r *Router) Close() error {
	return r.app.Shutdown()
}

func (r *Router) setupAccountRoutes() {
	r.app.Get("/register", r.users.RenderRegisterPage)
	r.app.Post("/register", r.users.RegisterUser)
	r.app.Get("/login", r.users.RenderLoginPage)
	r.app.Post("/login", r.users.LoginUser)
}
