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

	//************************************************************************************//
	//	ユーザー														　					//
	//************************************************************************************//

	// ユーザーを作成する
	u := &models.User{}
	u.Name = "test3"
	u.Email = "test3@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()

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

	//************************************************************************************//
	//	Todo														　  				  //
	//************************************************************************************//

	// Todoを作成する
	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// Todoを取得する（単数）
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// Todoを取得する（複数）
	user, _ := models.GetUser(3)
	user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// Todoを取得する（ユーザー絞り込み）
	user2, _ := models.GetUser(2)
	todos, _ := user2.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}
}
