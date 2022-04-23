package main

import (
	"fmt"
	"todo_bygolang/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("testだよ")

	fmt.Println(models.Db)
}
