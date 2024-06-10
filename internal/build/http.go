package build

import (
	_ "github.com/3Danger/currency/api/swagger"
	"github.com/3Danger/currency/internal/models"
	"github.com/3Danger/currency/internal/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"golang.org/x/net/context"
)

func (b *Builder) ConfigureAPI(ctx context.Context) func(ctx context.Context) error {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if e, ok := err.(*fiber.Error); ok {
				_ = c.Status(e.Code).JSON(rest.Error{
					Message: e.Error(),
				})

				return nil
			}

			if requestErr := new(models.Error); errors.As(err, &requestErr) {
				_ = c.Status(fiber.StatusBadRequest).JSON(rest.Error{
					Message: requestErr.Error(),
				})

				return nil
			}

			if err != nil {
				zerolog.Ctx(ctx).
					Error().
					Err(err).
					Str("path", c.Path()).
					Msg("handing request")
			}

			_ = c.Status(fiber.StatusInternalServerError).JSON(rest.Error{
				Message: "internal server error",
			})

			return nil
		},
	})

	srvConverter := b.NewServiceConverter()

	handler := rest.NewHandler(srvConverter)

	app.Get("/api/swagger/*", fiberSwagger.WrapHandler)
	app.Post("/api/convert", handler.Convert)

	return func(ctx context.Context) error {
		return app.Listen(b.cnf.Rest.Port)
	}
}
