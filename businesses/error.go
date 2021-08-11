package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrRoomsIDResource = errors.New("(RoomsID) not found or empty")

	ErrRoomsTitleResource = errors.New("(RoomsTitle) not found or empty")

	ErrCategoryNotFound = errors.New("category not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")
)
