package configs

import (
	"rentRoom/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func InitDB(){
	dsn := "root:@tcp(localhost:3306)/rentroom?parseTime=true"
	var err error
   DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil){
		panic("Failed Connection Database")
	}
	Migrate()
}

func Migrate(){
	DB.AutoMigrate(&users.User{})
}