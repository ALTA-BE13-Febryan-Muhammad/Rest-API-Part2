package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "id must integer",
		})
	}
	var data = User{
		Id:       id,
		Name:     "Name 1",
		Email:    "Email 1",
		Password: "Password 1",
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"data":    data,
	})
}

// // delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var x int

	for k, v := range users {
		if v.Id == id {
			x = k
		}
	}
	users = removearr(users, x)
	return c.JSON(http.StatusOK, map[string]any{
		"status": "Delete user success",
		"user":   users,
	})
}
func removearr(x []User, y int) []User {
	return append(x[:y], x[y+1:]...)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	// var user User
	idParam := c.Param("id") //id yang akan diupdate
	id, _ := strconv.Atoi(idParam)

	user := User{}
	errBind := c.Bind(&user) //data yang diupdate
	// c.Request().FormFile("foto")
	fmt.Println("error", errBind)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error binding data" + errBind.Error(),
		})
	}

	if user.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "nama harus diisi",
		})
	}

	fmt.Println("nama user", user.Name)
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"data":    user,
		"id":      id,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
