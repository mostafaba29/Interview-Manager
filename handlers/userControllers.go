package handlers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"mostafaba29/intialization"
	"net/http"
	"time"

	"mostafaba29/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var PrivateKey *ecdsa.PrivateKey

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

	return c.Status(200).JSON(fiber.Map{
		"message":  "account created",
		"position": user.Position})
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

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"issuer":  user.Username,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	})

	PrivateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	token, err := claims.SignedString(PrivateKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not login " + err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message":  "logged in",
		"position": user.Position,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logged Out",
	})
}
