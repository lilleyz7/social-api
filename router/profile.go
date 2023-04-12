package router

import (
	"auth-service/initializers"
	"auth-service/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddProfileGroup(app *fiber.App) {
	profileGroup := app.Group("/profile/:username")

	profileGroup.Get("/", getUserProfile)
	profileGroup.Post("/addFriend/:friend", followUser)
	profileGroup.Post("/update", updateProfile)
}

func getUserProfile(c *fiber.Ctx) error {
	user := models.User{}
	c.Params("username")
	collection := initializers.GetDBCollection("Users")
	response := collection.FindOne(c.Context(), bson.M{"username": user.Username})
	err := response.Decode(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"data": user,
		})
	}
}

func followUser(c *fiber.Ctx) error {
	self := c.Params("username")
	friend := c.Params("friend")
	user := models.User{}
	collection := initializers.GetDBCollection("Users")
	response := collection.FindOne(c.Context(), bson.M{"username": self})
	err := response.Decode(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user.Friends = append(user.Friends, friend)
	updated, err := collection.UpdateOne(c.Context(), bson.M{"username": self}, bson.M{"$set": user})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "update failed",
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": updated,
	})
}
func updateProfile(c *fiber.Ctx) error {
	user := models.User{}
	collection := initializers.GetDBCollection("Users")

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "body required",
		})
	}
	response, err := collection.UpdateOne(c.Context(), bson.M{"username": user.Username}, bson.M{"$set": user})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "failed to update",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": response,
	})
}
