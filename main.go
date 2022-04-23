package main

import (
	"fmt"
	"todo_bygolang/app/models"
)

func main() {
	// ログファイルを作成しログを書き込む
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("testだよ")

	fmt.Println(models.Db)

	// ユーザーを作成する
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// ユーザーを呼び出す
	// u, _ := models.GetUser(1)

	// fmt.Println(u)

	// ユーザー情報を更新する
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UPdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)
}
