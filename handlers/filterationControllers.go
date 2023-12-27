package handlers

import (
	"mostafaba29/intialization"
	"mostafaba29/models"

	"github.com/gofiber/fiber/v2"
)

// func GetCurrentUser(c *fiber.Ctx) error {
// 	sess, err := middleware.Store.Get(c)
// 	if err != nil {
// 		c.JSON("failed to get user from session")
// 	}
// 	fmt.Println(sess)

// 	username := sess.Get(middleware.User)
// 	if username == nil {
// 		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": "not autherized cant get id",
// 		})
// 	}

// 	var user models.User
// 	intialization.DB.Where("username=?", username).First(&user)
// 	if err != nil {
// 		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": "not autherized something is wrong",
// 		})
// 	}

// 	return c.JSON(user)
// }

func ShowAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	intialization.DB.Find(&appointments)
	return c.JSON(appointments)
}

// func ShowManagerAppointments(c *fiber.Ctx) error {
//user := GetCurrentUser(c)
// sess, err := middleware.Store.Get(c)
// if err != nil {
// 	c.JSON("failed to get user from session")
// }
// fmt.Println(sess)
// username := sess.Get(middleware.User)
// if username == nil {
// 	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"error": "not autherized cant get id",
// 	})
// }

// var user models.User
// manager_name := intialization.DB.Where("username=?", username).First(&user)
// if err != nil {
// 	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"error": "not autherized something is wrong",
// 	})
// }
//fmt.Println(username)

// var appointments []models.Appointment
// // managerName := user
// //fmt.Println(managerName)
// if err := intialization.DB.Where("manager_name = ?", username).Find(&appointments).Error; err != nil {
// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// }

// return c.JSON(appointments)
//}

// func ShowApprovedAppointments(c *fiber.Ctx) error {
// 	var approvedAppointments []models.Appointment
// 	if err := intialization.DB.Where("status=?", "Confirmed").First(&approvedAppointments); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"massege": "no approved appointments found",
// 		})
// 	}
// 	return c.JSON(approvedAppointments)
// }

// func ShowCanceledAppointments(c *fiber.Ctx) error {
// 	var canceledAppointments []models.Appointment
// 	if err := intialization.DB.Where("status=?", "Declined").First(&canceledAppointments); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"massege": "no canceled appointments found",
// 		})

// 	}
// 	return c.JSON(canceledAppointments)
// }
