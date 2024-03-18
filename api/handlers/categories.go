package handlers

import (
	"fmt"
	"net/http"

	"github.com/jssdeveloper/paperless/db"
	"github.com/labstack/echo/v4"
)

type Categories struct {
	Category string `json:"category"`
}

type CategoriesWithCount struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

func CreateCategory(c echo.Context) error {
	category := Categories{}
	err := c.Bind(&category)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error binding data")
	}

	err = db.DB.Get(&category, "INSERT INTO categories (category) VALUES (?) RETURNING *", category.Category)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func GetCategories(c echo.Context) error {
	// get all categories with count
	categories := []CategoriesWithCount{}
	err := db.DB.Select(&categories, "SELECT c.category, COUNT(m.id) as count FROM categories c LEFT JOIN memos m ON c.category = m.category GROUP BY c.category")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, categories)
}

func PatchCategory(c echo.Context) error {
	oldCategory := c.Param("category")
	newCategory := Categories{}
	err := c.Bind(&newCategory)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error binding data")
	}

	_, err = db.DB.Exec("UPDATE categories SET category = ? WHERE category = ?", newCategory.Category, oldCategory)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "category has been updated")
}

func DeleteCategory(c echo.Context) error {
	category := c.Param("category")

	// check if category is used in memos
	var count int
	err := db.DB.Get(&count, "SELECT COUNT(*) FROM memos WHERE category = ?", category)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if count > 0 {
		return c.String(http.StatusBadRequest, "category is used in memos")
	}

	_, err = db.DB.Exec("DELETE FROM categories WHERE category = ?", category)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "category has been deleted")
}
