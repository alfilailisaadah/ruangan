package rooms

import (
	"errors"
	"net/http"
	"rentRoom/app/middleware"
	"rentRoom/businesses/rooms"
	"rentRoom/controllers/rooms/request"
	"rentRoom/controllers/rooms/response"

	controller "rentRoom/controllers"

	echo "github.com/labstack/echo/v4"
)

type RoomsController struct {
	roomsUsecase rooms.Usecase
}

func NewRoomsController(cu rooms.Usecase) *RoomsController {
	return &RoomsController{
		roomsUsecase: cu,
	}
}

func (ctrl *RoomsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.roomsUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Rooms{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *RoomsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Rooms{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.roomsUsecase.Store(ctx,req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *RoomsController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)
	if user.UserType == "2" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("invalid role"))
	}
	req := request.Rooms{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	resp, err := ctrl.roomsUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}