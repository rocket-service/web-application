package users

import "github.com/gofiber/fiber/v2"

func (s *Service) RenderRegisterPage(ctx *fiber.Ctx) error {
	return ctx.Render("auth/signup", fiber.Map{})
}

func (s *Service) RenderLoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("auth/signin", fiber.Map{})
}
