package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
  Id    int    `json:"id" form:"id"`
  Name  string `json:"name" form:"name"`
  Email string `json:"email" form:"email"`
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
  userID := c.Param("id")

  var foundUser *User
  for _, user := range users {
    if strconv.Itoa(user.Id) == userID {
      foundUser = &user
      break
    }
  }

  if foundUser != nil {
    return c.JSON(http.StatusOK, map[string]interface{}{
      "message": "success get user by ID",
      "user":   foundUser,
    })
  }
  
  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "message": "user not found",
  })
}

// delete user by id
func DeleteUserController(c echo.Context) error {
  // your solution here
  userID := c.Param("id")

  index := -1
  for i, user := range users {
    if strconv.Itoa(user.Id) == userID {
      index = i
      break
    }
  }

  if index != -1 {
    users = append(users[:index], users[index+1:]...)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user by ID",
		})
  }

  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "message": "user not found",
  })
}

// update user by id
func UpdateUserController(c echo.Context) error {
  // your solution here
  userID := c.Param("id")

  var foundUser *User
  for _, user := range users {
    if strconv.Itoa(user.Id) == userID {
      foundUser = &user
      break
    }
  }

  if foundUser != nil {
    if err := c.Bind(foundUser); err != nil {
      return err
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
      "message": "success update user by ID",
      "user":    foundUser,
    })
  }

  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "message": "user not found",
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
  e.PUT("/users/:id", UpdateUserController)
  e.DELETE("/users/:id", DeleteUserController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}