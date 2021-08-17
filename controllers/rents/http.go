package rents

import (
	"net/http"
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
