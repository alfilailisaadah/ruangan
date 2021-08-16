package buildings

import (
	"net/http"
	"rentRoom/businesses/buildings"
	controller "rentRoom/controllers"
	"rentRoom/controllers/buildings/request"
	"rentRoom/controllers/buildings/response"

	echo "github.com/labstack/echo/v4"
)

type BuildingsController struct {
	buildingsUseCase buildings.Usecase
}

func NewBuildingsController(buildingsUC buildings.Usecase) *BuildingsController {
	return &BuildingsController{
		buildingsUseCase: buildingsUC,
	}
}

func (ctrl *BuildingsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	ip := c.QueryParam("ip")

	req := request.Buildings{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.buildingsUseCase.Store(ctx, ip, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *BuildingsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.buildingsUseCase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Buildings{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}
