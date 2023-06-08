package controllers

import (
	"go-latihan/database"
	"go-latihan/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	blogs, e := database.GetBlogs()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	BlogId := GetIdFromParam(c)
	blog, e := database.GetBlogByUserId(BlogId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"blogs":   blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := &models.Blog{}

	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdBlog, err := database.CreateBlog(blog)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    createdBlog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	BlogId := GetIdFromParam(c)
	blog, e := database.DeleteBlogById(BlogId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Delete blog",
		"blog":    blog,
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	BlogId := GetIdFromParam(c)
	updatedBlog := &models.Blog{}

	if e := c.Bind(updatedBlog); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	blog, e := database.UpdateBlogById(BlogId, updatedBlog)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})
}

func GetBlogUserIdFromParam(c echo.Context) int {
	BlogUserId, _ := strconv.Atoi(c.Param("id"))
	return BlogUserId
}
