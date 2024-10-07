package api

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// User represents a basic user structure
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "John Doe", Age: 25},
	{ID: 2, Name: "Jane Doe", Age: 30},
}

func Main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize a new Fiber app
	app := fiber.New()

	// Route to serve static files (like HTML, CSS, etc.)
	app.Static("/", "./public")

	// Simple route: GET request for root path "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the User API!")
	})

	// Route to get all users
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	// Route to get a single user by ID
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
		}

		for _, user := range users {
			if user.ID == id {
				return c.JSON(user)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	})

	// Route to create a new user
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)

		// Parse the incoming request body
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request")
		}

		user.ID = len(users) + 1
		users = append(users, *user)
		return c.Status(fiber.StatusCreated).JSON(user)
	})

	// Route to update an existing user
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
		}

		updatedUser := new(User)
		if err := c.BodyParser(updatedUser); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request")
		}

		for i, user := range users {
			if user.ID == id {
				users[i].Name = updatedUser.Name
				users[i].Age = updatedUser.Age
				return c.JSON(users[i])
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	})

	// Route to delete a user by ID
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
		}

		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				return c.SendString("User deleted successfully")
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	})

	app.Listen(os.Getenv("API_LISTEN_URL"))
}
