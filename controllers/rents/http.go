package rents

import (
	"errors"
	"net/http"
	"rentRoom/app/middleware"
	"rentRoom/businesses"
	"rentRoom/businesses/rents"
	controller "rentRoom/controllers"
	"rentRoom/controllers/rents/request"
	"rentRoom/controllers/rents/response"

	echo "github.com/labstack/echo/v4"
)

type RentsController struct {
	rentsUseCase rents.Usecase
}

func NewRentsController(rentsUC rents.Usecase) *RentsController {
	return &RentsController{
		rentsUseCase: rentsUC,
	}
}

func (ctrl *RentsController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)
	if user.UserType == "1" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("invalid role"))
	}
	req := request.Rents{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.rentsUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, businesses.ErrUserIdorRoomIdNotFound)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *RentsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.rentsUseCase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Rents{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *RentsController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)
	if user.UserType == "2" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("invalid role"))
	}
	req := request.Rents{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	resp, err := ctrl.rentsUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}
