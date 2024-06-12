package controller

// import (
// 	"time"

// 	"github.com/JacobRWebb/InventoryManagement/pkg/models"
// 	"github.com/JacobRWebb/InventoryManagement/pkg/queries"
// 	"github.com/JacobRWebb/InventoryManagement/pkg/utils"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// )

// func ViewRegisterAccount(c *fiber.Ctx) error {
// 	return nil
// }

// func RegisterAccount(c *fiber.Ctx) error {
// 	registerAccountRequest := &models.RegisterAccountRequest{}

// 	if err := c.BodyParser(registerAccountRequest); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "Please try again.",
// 		})
// 	}

// 	validate := utils.NewValidator()

// 	if err := validate.Struct(registerAccountRequest); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   utils.ValidatorErrors(err),
// 		})
// 	}

// 	user := &models.User{
// 		Id:           uuid.New(),
// 		CreatedAt:    time.Now(),
// 		ModifiedAt:   time.Now(),
// 		Email:        registerAccountRequest.Email,
// 		PasswordHash: utils.GeneratePasswordHash(registerAccountRequest.Password),
// 	}

// 	err := queries.CreateUser(user)

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	return c.SendStatus(fiber.StatusNotImplemented)
// }

// func GetUserByEmail(c *fiber.Ctx) error {
// 	getUserByEmailRequest := &models.GetUserByEmailRequest{}

// 	if err := c.BodyParser(getUserByEmailRequest); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "Unable to parse request" + err.Error(),
// 		})
// 	}

// 	validate := utils.NewValidator()

// 	if err := validate.Struct(getUserByEmailRequest); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   utils.ValidatorErrors(err),
// 		})
// 	}

// 	user, err := queries.GetUserByEmail(getUserByEmailRequest.Email)

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   utils.ValidatorErrors(err),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"user": user,
// 	})
// }
