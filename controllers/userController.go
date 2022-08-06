package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rekib0023/go-crud-apis/database"
	"github.com/rekib0023/go-crud-apis/models"
	"github.com/rekib0023/go-crud-apis/serializers"
	"golang.org/x/crypto/bcrypt"
)

func UserResponse(user models.User) serializers.UserSerializer {
	return serializers.UserSerializer{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
}

func CreateUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	database.Database.Db.Create(&user)

	responseUser := UserResponse(user)

	return c.Status(fiber.StatusCreated).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	database.Database.Db.Find(&users)
	responseUsers := []serializers.UserSerializer{}

	for _, user := range users {
		responseUser := UserResponse(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(fiber.StatusOK).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(err.Error())
	}

	responseUser := UserResponse(user)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	var updateUser UpdateUser

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(err.Error())
	}

	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName
	user.Email = updateUser.Email

	database.Database.Db.Save(&user)
	responseUser := UserResponse(user)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(err.Error())
	}

	return c.Status(fiber.StatusNoContent).JSON("Successfully Deleted User")
}
