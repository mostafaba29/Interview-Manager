package handlers

import (
	"mostafaba29/intialization"
	"mostafaba29/middleware"
	"time"

	"mostafaba29/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	var userData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Position string `json:"position"`
	}
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 14)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to hash password",
		})
	}

	user := models.User{Username: userData.Username, Password: hashedPass, Position: userData.Position}
	result := intialization.DB.Create(&user)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"massege": "failed to create user",
		})
	}

	return c.Status(200).JSON("account created")
}

func Login(c *fiber.Ctx) error {

	var userInfo struct {
		Username string
		Password string
		Position string
	}
	var user models.User

	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "couldn't read user data",
		})
	}

	intialization.DB.Where("username=?", userInfo.Username).First(&user)

	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"massege": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(userInfo.Password)); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"massege": "incorrect password",
		})
	}

	session, err := middleware.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong" + err.Error(),
		})
	}
	sessionID := session.ID()

	cookie := fiber.Cookie{
		Name:     "session",
		Value:    sessionID,
		Expires:  time.Now().Add(time.Hour * 24 * 10),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	session.Set(middleware.USER_ID, user.ID)
	session.Set(middleware.AUTH_KEY, true)
	sessERR := session.Save()
	if sessERR != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong" + err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"user": "logged in as " + user.Position, "session": sessionID})
}

func Logout(c *fiber.Ctx) error {
	session, err := middleware.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "no session",
		})
	}
	sessERR := session.Destroy()
	if sessERR != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "something is wrong" + err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"massege": "logged out",
	})
}
