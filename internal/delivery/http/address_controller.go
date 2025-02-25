package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AddressController struct {
	UseCase *usecase.AddressUseCase
	Log     *logrus.Logger
}

// Structs for swagger documentation
type WebAddressResponse model.WebResponse[*model.AddressResponse]
type WebAddressListResponse model.WebResponse[[]model.AddressResponse]
type WebAddressDeleteResponse model.WebResponse[bool]

func NewAddressController(useCase *usecase.AddressUseCase, log *logrus.Logger) *AddressController {
	return &AddressController{
		Log:     log,
		UseCase: useCase,
	}
}

// Create method to create new address.
// @Description Create new address.
// @Summary create new address
// @Tags Address
// @Accept json
// @Produce json
// @Param contactId path string true "Contact ID"
// @Param body body model.CreateAddressRequest true "Create Address Request"
// @Success 200 {object} WebAddressResponse
// @Router /contacts/{contactId}/addresses [post]
func (c *AddressController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ContactId = ctx.Params("contactId")

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

// List method to list addresses.
// @Description List addresses.
// @Summary list addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param contactId path string true "Contact ID"
// @Success 200 {object} WebAddressListResponse
// @Router /contacts/{contactId}/addresses [get]
func (c *AddressController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	request := &model.ListAddressRequest{
		UserId:    auth.ID,
		ContactId: contactId,
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list addresses")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.AddressResponse]{Data: responses})
}

// Get method to get address.
// @Description Get address.
// @Summary get address
// @Tags Address
// @Accept json
// @Produce json
// @Param contactId path string true "Contact ID"
// @Param addressId path string true "Address ID"
// @Success 200 {object} WebAddressResponse
// @Router /contacts/{contactId}/addresses/{addressId} [get]
func (c *AddressController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	addressId := ctx.Params("addressId")

	request := &model.GetAddressRequest{
		UserId:    auth.ID,
		ContactId: contactId,
		ID:        addressId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

// Update method to update address.
// @Description Update address.
// @Summary update address
// @Tags Address
// @Accept json
// @Produce json
// @Param contactId path string true "Contact ID"
// @Param addressId path string true "Address ID"
// @Param body body model.UpdateAddressRequest true "Update Address Request"
// @Success 200 {object} WebAddressResponse
// @Router /contacts/{contactId}/addresses/{addressId} [put]
func (c *AddressController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ContactId = ctx.Params("contactId")
	request.ID = ctx.Params("addressId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

// Delete method to delete address.
// @Description Delete address.
// @Summary delete address
// @Tags Address
// @Accept json
// @Produce json
// @Param contactId path string true "Contact ID"
// @Param addressId path string true "Address ID"
// @Success 200 {object} WebAddressDeleteResponse
// @Router /contacts/{contactId}/addresses/{addressId} [delete]
func (c *AddressController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	addressId := ctx.Params("addressId")

	request := &model.DeleteAddressRequest{
		UserId:    auth.ID,
		ContactId: contactId,
		ID:        addressId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete address")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
