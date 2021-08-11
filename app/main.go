package main

import (
	_categoryUsecase "rentRoom/businesses/category"
	_categoryController "rentRoom/controllers/category"
	_categoryRepo "rentRoom/drivers/databases/category"

	_userUsecase "rentRoom/businesses/users"
	_userController "rentRoom/controllers/users"
	_userRepo "rentRoom/drivers/databases/users"

	_dbDriver "rentRoom/drivers/mysql"

	// _ipLocatorDriver "rentRoom/drivers/thirdparties/iplocator"

	_middleware "rentRoom/app/middleware"
	_routes "rentRoom/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
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

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	// iplocatorRepo := _ipLocatorDriver.NewIPLocator()

	categoryRepo := _categoryRepo.NewCategoryRepository(db)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(timeoutContext, categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	// newsRepo := _newsRepo.NewMySQLNewsRepository(db)
	// newsUsecase := _newsUsecase.NewNewsUsecase(newsRepo, categoryUsecase, timeoutContext, iplocatorRepo)
	// newsCtrl := _newsController.NewNewsController(newsUsecase)

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userCtrl,
		// NewsController:     *newsCtrl,
		CategoryController: *categoryCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
