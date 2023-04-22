package router

import (
	"auth-service/initializers"
	"auth-service/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddPostGroup(app *fiber.App) {
	postGroup := app.Group("/post")

	postGroup.Get("/", getAllPosts)
	postGroup.Get("/:username", getPostsByUser)
	postGroup.Get("/:username/:title", getPostByTitle)
	postGroup.Post("/new", createNewPost)
	postGroup.Put("/update", updatePost)

}

func getAllPosts(c *fiber.Ctx) error {
	collection := initializers.GetDBCollection("Posts")
	allPosts := make([]models.Post, 0)
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for cursor.Next(c.Context()) {
		singlePost := models.Post{}
		err := cursor.Decode(&singlePost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		allPosts = append(allPosts, singlePost)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": allPosts,
	})

}

func getPostsByUser(c *fiber.Ctx) error {
	posts := make([]models.Post, 0)
	username := c.Params("username")
	collection := initializers.GetDBCollection("Posts")

	cursor, err := collection.Find(c.Context(), bson.M{"username": username})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for cursor.Next(c.Context()) {
		singlePost := models.Post{}
		err := cursor.Decode(&singlePost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		posts = append(posts, singlePost)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": posts,
	})

}

func getPostByTitle(c *fiber.Ctx) error {
	username := c.Params("username")
	title := c.Params("title")
	post := models.Post{}

	collection := initializers.GetDBCollection("posts")
	result := collection.FindOne(c.Context(), bson.M{"username": username, "title": title})
	err := result.Decode(&post)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid username, title combination",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": post,
	})
}

func createNewPost(c *fiber.Ctx) error {
	newPost := models.Post{}
	err := c.BodyParser(&newPost)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	collection := initializers.GetDBCollection("posts")
	result, err := collection.InsertOne(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "unable to insert into db",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data":    result,
		"message": "successfully entered into db",
	})

}

func updatePost(c *fiber.Ctx) error {
	return nil

}
