package rest

import (
	"bytes"
	"encoding/json"

	"github.com/3Danger/currency/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
)

type service interface {
	Convert(ctx context.Context, pair models.Pair, value decimal.Decimal) (
		result decimal.Decimal, mediatorCode *models.Code, err error,
	)
}

type Handler struct {
	svc service
}

func NewHandler(svc service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Convert(c *fiber.Ctx) error {
	params := BodyParams{}

	if err := json.NewDecoder(bytes.NewBuffer(c.Body())).Decode(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := params.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, mediator, err := h.svc.Convert(c.Context(), models.JoinCodes(*params.From, *params.To), *params.Value)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	respResult, err := json.Marshal(Result{Result: result, MediatorCode: mediator})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Send(respResult)
}
