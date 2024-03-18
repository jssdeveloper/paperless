package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jssdeveloper/paperless/db"
	"github.com/labstack/echo/v4"
)

type Memos struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Date        string `json:"date"`
}

func CreateMemo(c echo.Context) error {
	memo := Memos{}
	err := c.Bind(&memo)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error binding data")
	}

	err = db.DB.Get(&memo, "INSERT INTO memos (title, description, category, date) VALUES (?, ?, ?, ?) RETURNING *", memo.Title, memo.Description, memo.Category, fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day()))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, memo)
}

func GetAllMemos(c echo.Context) error {
	memos := []Memos{}
	err := db.DB.Select(&memos, "SELECT * FROM memos")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, memos)
}

func GetMemosByCategory(c echo.Context) error {
	memos := []Memos{}
	category := c.Param("category")
	err := db.DB.Select(&memos, "SELECT * FROM memos WHERE category = ?", category)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, memos)
}

func PatchMemo(c echo.Context) error {
	id := c.Param("id")
	memo := Memos{}
	err := c.Bind(&memo)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error binding data")
	}

	_, err = db.DB.Exec("UPDATE memos SET title = ?, description = ?, category = ? WHERE id = ?", memo.Title, memo.Description, memo.Category, id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "memo has been updated")
}

func DeleteMemo(c echo.Context) error {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM memos WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "memo has been deleted")
}
