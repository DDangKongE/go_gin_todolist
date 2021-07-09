package main

import (
	"fmt"
	"go_board/src/Config"
	"go_board/src/Models"
	"go_board/src/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()

	r.Run()
}
