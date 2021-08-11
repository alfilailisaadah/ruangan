package rooms

import (
	"net/http"
	"rentRoom/businesses/rooms"
	controller "rentRoom/controllers"
	"rentRoom/controllers/rooms/request"

	echo "github.com/labstack/echo/v4"
)

type RoomsController struct {
	roomsUseCase rooms.Usecase
}

func NewRoomsController(roomsUC rooms.Usecase) *RoomsController {
	return &RoomsController{
		roomsUseCase: roomsUC,
	}
}

func (ctrl *RoomsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	ip := c.QueryParam("ip")

	req := request.Rooms{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.roomsUseCase.Store(ctx, ip, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}
