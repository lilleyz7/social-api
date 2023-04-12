package router

import (
	"auth-service/initializers"
	"auth-service/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func AddAuthGroup(app *fiber.App) {
	authGroup := app.Group("/auth")

	authGroup.Get("/allUsers", getAllUsers)
	authGroup.Post("/register", registerUser)
	authGroup.Post("/login", authenticateUser)
	authGroup.Delete("/delete/:username", deleteUser)

}

func getAllUsers(c *fiber.Ctx) error {
	collection := initializers.GetDBCollection("Users")

	allUsers := make([]models.User, 0)
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for cursor.Next(c.Context()) {
		singleUser := models.User{}
		err := cursor.Decode(&singleUser)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		allUsers = append(allUsers, singleUser)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": allUsers,
	})
}

func registerUser(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "code not hash pass",
		})
	}

	collection := initializers.GetDBCollection("Users")
	check := collection.FindOne(c.Context(), bson.M{"username": user.Username})
	if check != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "username already exists",
		})
	}

	user.Password = string(hashedPassword)
	result, err := collection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "failed to add user",
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": result,
	})

}

func deleteUser(c *fiber.Ctx) error {
	collection := initializers.GetDBCollection("Users")
	username := c.Params("username")
	if username == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "name was not given",
		})
	}

	result, err := collection.DeleteOne(c.Context(), bson.M{"username": username})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "name does not exist in db",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

func authenticateUser(c *fiber.Ctx) error {
	user := models.User{}
	check := models.User{}

	c.BodyParser(user)
	collection := initializers.GetDBCollection("Users")
	existingUser := collection.FindOne(c.Context(), bson.M{"username": user.Username})
	err := existingUser.Decode(&check)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "could not bind find to user",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(check.Password))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "mismatched passwords",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": "success",
	})
}
