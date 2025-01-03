package users

import (
	"context"
	"errors"
	"rocket-web/internal/storage"
	"rocket-web/internal/storage/models"
	"rocket-web/internal/storage/postgres"
	"rocket-web/pkg/passwordhash"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Service struct {
	storage *postgres.Storage
	log     *zap.SugaredLogger
}

type UserStorage interface {
	SaveUser(ctx context.Context, username, password string) (int64, error)
	GetUser(ctx context.Context, username string) (models.User, error)
}

func New(storage *postgres.Storage, log *zap.SugaredLogger) *Service {
	return &Service{storage: storage, log: log}
}

func (s *Service) RegisterUser(ctx *fiber.Ctx) error {
	const prefix = "internal.router.services.account.Register"
	username, password := ctx.FormValue("username"), ctx.FormValue("password")

	log := s.log.With(
		zap.String("op", prefix),
		zap.String("username", username),
	)

	if len(password) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password is required",
		})
	}

	log.Info("Attempting to register user", zap.String("username", username))
	passwordHash, err := passwordhash.New(password)
	if err != nil {
		log.Warnw("Failed to hash password", "error", err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	id, err := s.storage.SaveUser(ctx.Context(), username, passwordHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserAlreadyExists) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "user already exists",
			})
		} else {

			log.Warnw("Failed to save user", "error", err)

			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "failed to save user",
			})
		}
	}

	log.Info("User registered", zap.Int64("id", id))

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func (s *Service) LoginUser(ctx *fiber.Ctx) error {
	return nil
}
