package main

import (
	"github.com/jssdeveloper/paperless/db"
	"github.com/jssdeveloper/paperless/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db.Connect()
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5: true,
		Root:  "dist", // because files are located in `web` directory in `webAssets` fs
	}))

	api := e.Group("/api")

	api.POST("/categories", handlers.CreateCategory)             // Create a category
	api.GET("/categories", handlers.GetCategories)               // Get all categories
	api.PATCH("/categories/:category", handlers.PatchCategory)   // Update a category
	api.DELETE("/categories/:category", handlers.DeleteCategory) // Delete a category
	api.POST("/memos", handlers.CreateMemo)                      // Create a memo
	api.GET("/memos", handlers.GetAllMemos)                      // Get all memos
	api.GET("/memos/:category", handlers.GetMemosByCategory)     // Get memos by category
	api.PATCH("/memos/:id", handlers.PatchMemo)                  // Update a memo
	api.DELETE("/memos/:id", handlers.DeleteMemo)                // Delete a memo

	e.Logger.Fatal(e.Start(":1323"))
}
