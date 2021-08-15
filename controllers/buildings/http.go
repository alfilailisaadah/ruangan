package buildings

import (
	"errors"
	"net/http"
	"rentRoom/businesses/buildings"
	controller "rentRoom/controllers"
	"rentRoom/controllers/buildings/request"
	"rentRoom/controllers/buildings/response"
	"strconv"
	"strings"

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

func (ctrl *BuildingsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Buildings{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	resp, err := ctrl.buildingsUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}
