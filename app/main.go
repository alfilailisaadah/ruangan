package main

import (
	_buildingsUsecase "rentRoom/businesses/buildings"
	_buildingsController "rentRoom/controllers/buildings"
	_buildingsRepo "rentRoom/drivers/databases/buildings"

	_roomsUsecase "rentRoom/businesses/rooms"
	_roomsController "rentRoom/controllers/rooms"
	_roomsRepo "rentRoom/drivers/databases/rooms"

	_rentsUsecase "rentRoom/businesses/rents"
	_rentsController "rentRoom/controllers/rents"
	_rentsRepo "rentRoom/drivers/databases/rents"

	_userUsecase "rentRoom/businesses/users"
	_userController "rentRoom/controllers/users"
	_userRepo "rentRoom/drivers/databases/users"

	_dbDriver "rentRoom/drivers/mysql"

	_middleware "rentRoom/app/middleware"
	_routes "rentRoom/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_buildingsRepo.Buildings{},
		&_roomsRepo.Rooms{},
		&_userRepo.Users{},
		&_rentsRepo.Rents{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	roomsRepo := _roomsRepo.NewRoomsRepository(db)
	roomsUsecase := _roomsUsecase.NewRoomsUsecase(timeoutContext, roomsRepo)
	roomsCtrl := _roomsController.NewRoomsController(roomsUsecase)

	rentsRepo := _rentsRepo.NewRentsRepository(db)
	rentsUsecase := _rentsUsecase.NewRentsUsecase(rentsRepo,timeoutContext)
	rentsCtrl := _rentsController.NewRentsController(rentsUsecase)

	buildingsRepo := _buildingsRepo.NewMySQLBuidingsRepository(db)
	buildingsUsecase := _buildingsUsecase.NewBuildingsUsecase(buildingsRepo,timeoutContext)
	buildingsCtrl := _buildingsController.NewBuildingsController(buildingsUsecase)

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userCtrl,
		BuildingsController:     *buildingsCtrl,
		RentsController:     *rentsCtrl,
		RoomsController: *roomsCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
