package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vivekv96/go-admin/database"
	"github.com/vivekv96/go-admin/models"
	"github.com/vivekv96/go-admin/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// return error message if `password` & `confirmPassword` fields are not identical
	if data["password"] != data["confirmPassword"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "passwords do not match!",
		})
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

	user := models.User{
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
		Password:  string(passwordHash),
	}
	database.Gorm.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	err := database.Gorm.Where("email = ?", data["email"]).First(&user).Error
	if err != nil || user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "no user found with that email!",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "entered wrong password!",
		})
	}

	token, err := utils.GenerateJWT(fmt.Sprint(user.ID))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"jwt": token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJWT(cookie)

	var user models.User
	database.Gorm.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
