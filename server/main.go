package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GroceryList struct {
	ID           int      `json:"id"`
	GroceryItems []string `json:"groceryItems"`
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500", "http://localhost:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/grocery-list", func(c *gin.Context) {

		id := c.Query("id")

		res := GroceryList{}
		switch id {
		case "1":
			res = GroceryList{
				ID: 1,
				GroceryItems: []string{
					"Avocados",
					"Ground Beef",
					"Spinach",
					"Blueberries",
					"Coffee",
					"Almond Milk",
					"Tortillas",
					"Black Beans",
					"Salsa",
					"Shredded Cheese",
					"Frozen Broccoli",
					"Orange Juice",
					"Peanut Butter",
					"Crackers",
					"Ice Cream",
				},
			}

		case "2":
			res = GroceryList{
				ID: 2,
				GroceryItems: []string{
					"Bread",
					"Milk",
					"Eggs",
					"Butter",
					"Cheese",
					"Chicken Breast",
					"Rice",
					"Pasta",
					"Tomatoes",
					"Onions",
					"Garlic",
					"Bell Peppers",
					"Bananas",
					"Apples",
					"Yogurt",
				},
			}

		default:
			res = GroceryList{}
		}

		// Return JSON response
		c.JSON(http.StatusOK, res)
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
