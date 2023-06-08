package database

import (
	"go-latihan/config"
	"go-latihan/models"
)

func GetBlogs() (interface{}, error) {
	var Blogs []models.Blog
	if e := config.DB.Preload("User").Find(&Blogs).Error; e != nil {
		return nil, e
	}
	return Blogs, nil
}

func GetBlogByUserId(id int) (interface{}, error) {
	var Blog models.Blog
	if e := config.DB.Preload("User").Find(&Blog, id).Error; e != nil {
		return nil, e
	}
	return Blog, nil
}

func CreateBlog(blog *models.Blog) (*models.Blog, error) {
	if e := config.DB.Create(blog).Error; e != nil {
		return nil, e
	}
	return blog, nil
}

func DeleteBlogById(id int) (interface{}, error) {
	var Blog models.Blog
	if e := config.DB.Delete(&Blog, id).Error; e != nil {
		return nil, e
	}
	return Blog, nil
}

func UpdateBlogById(id int, updatedBlog *models.Blog) (*models.Blog, error) {
	var Blog models.Blog
	if e := config.DB.First(&Blog, id).Error; e != nil {
		return nil, e
	}

	Blog.UserId = updatedBlog.UserId
	Blog.Judul = updatedBlog.Judul
	Blog.Konten = updatedBlog.Konten

	if e := config.DB.Save(&Blog).Error; e != nil {
		return nil, e
	}
	return &Blog, nil
}
