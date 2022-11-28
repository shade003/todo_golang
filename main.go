package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test") // ログの出力

	fmt.Println(models.Db) // init関数を呼び出す

	controllers.StartMainServer()

	// // user情報取得
	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// // session情報取得
	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// // session情報チェック
	// valid, _ := session.CheckSession()
	// fmt.Println(valid)

	// user
	// C
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// R
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// U
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// D
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// todo
	// C
	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third todo")

	// R
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// U
	// user, _ := models.GetUser(2)
	// user.CreateTodo("first todo")

	// todo, _ := models.GetTodo(1)
	// todo.Content = "update todo"
	// todo.UpdateTodo()

	// D
	// t, _ := models.GetTodo(3)
	// t.DeleteTodo()
}
