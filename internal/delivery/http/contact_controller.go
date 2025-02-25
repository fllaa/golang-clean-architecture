package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ContactController struct {
	UseCase *usecase.ContactUseCase
	Log     *logrus.Logger
}

// Structs for swagger documentation
type WebContactResponse model.WebResponse[*model.ContactResponse]
type WebContactListResponse model.WebResponse[[]model.ContactResponse]
type WebContactDeleteResponse model.WebResponse[bool]

func NewContactController(useCase *usecase.ContactUseCase, log *logrus.Logger) *ContactController {
	return &ContactController{
		UseCase: useCase,
		Log:     log,
	}
}

// Create method to create new contact.
// @Description Create new contact.
// @Summary create new contact
// @Tags Contact
// @Accept json
// @Produce json
// @Param body body model.CreateContactRequest true "Create Contact Request"
// @Success 200 {object} WebContactResponse
// @Router /contacts [post]
func (c *ContactController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.UserId = auth.ID

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

// List method to list contacts.
// @Description List contacts.
// @Summary list contacts
// @Tags Contact
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param email query string false "Email"
// @Param phone query string false "Phone"
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Success 200 {object} WebContactListResponse
// @Router /contacts [get]
func (c *ContactController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.SearchContactRequest{
		UserId: auth.ID,
		Name:   ctx.Query("name", ""),
		Email:  ctx.Query("email", ""),
		Phone:  ctx.Query("phone", ""),
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching contact")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.ContactResponse]{
		Data:   responses,
		Paging: paging,
	})
}

// Get method to get contact.
// @Description Get contact.
// @Summary get contact
// @Tags Contact
// @Accept json
// @Produce json
// @Param contactId path string true "Contact ID"
// @Success 200 {object} WebContactResponse
// @Router /contacts/{contactId} [get]
func (c *ContactController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetContactRequest{
		UserId: auth.ID,
		ID:     ctx.Params("contactId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

// Update method to update contact.
// @Description Update contact.
// @Summary update contact
// @Tags Contact
// @Accept json
// @Produce json
// @Param contactId path string true "Contact ID"
// @Param body body model.UpdateContactRequest true "Update Contact Request"
// @Success 200 {object} WebContactResponse
// @Router /contacts/{contactId} [put]
func (c *ContactController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ID = ctx.Params("contactId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

// Delete method to delete contact.
// @Description Delete contact.
// @Summary delete contact
// @Tags Contact
// @Accept json
// @Produce json
// @Param contactId path string true "Contact ID"
// @Success 200 {object} WebContactDeleteResponse
// @Router /contacts/{contactId} [delete]
func (c *ContactController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	request := &model.DeleteContactRequest{
		UserId: auth.ID,
		ID:     contactId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
