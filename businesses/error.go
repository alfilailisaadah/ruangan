package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrBuildingsIDResource = errors.New("(BuildingsID) not found or empty")

	ErrBuildingsTitleResource = errors.New("(BuildingsTitle) not found or empty")

	ErrRoomsNotFound = errors.New("rooms not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")
	
	ErrUserIdorRoomIdNotFound = errors.New("(User Id) or (Room Id) Not Found")
)
