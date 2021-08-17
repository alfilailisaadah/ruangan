package rents

import (
	"net/http"
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

	ip := c.QueryParam("ip")

	req := request.Rents{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.rentsUseCase.Store(ctx, ip, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}
