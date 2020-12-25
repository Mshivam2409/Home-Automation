package auth

import (
	"encoding/json"
	"time"

	"github.com/Mshivam2409/Home-Automation/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// SignIn Signs in a new User
func SignIn(c *fiber.Ctx) error {
	existingUser := new(models.User)
	if err := c.BodyParser(existingUser); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})

	}
	user := &models.User{}
	coll := mgm.Coll(user)
	err := coll.First(bson.M{"username": existingUser.Username}, user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Check password
	if user.Password != existingUser.Password {
		return c.SendStatus(fiber.StatusForbidden)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{"message": "success", "token": t})
}

// SignUp Signs Up a New User
func SignUp(c *fiber.Ctx) error {
	payLoad := new(models.User)
	err := json.Unmarshal([]byte(c.Body()), &payLoad)
	if err != nil {
		panic(err)
	}
	user := models.NewUser(payLoad.Username, payLoad.Username, payLoad.Details)
	err = mgm.Coll(user).Create(user)
	if err != nil {
		panic(err)
	}
	return c.Status(201).JSON(fiber.Map{"message": "success"})
}
