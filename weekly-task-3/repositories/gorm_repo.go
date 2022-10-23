package repositories

import (
	"net/http"
	"weekly-task-3/models"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type GormSQL struct {
	DB *gorm.DB
}

func NewGorm(db *gorm.DB) Database {
	return &GormSQL{
		DB: db,
	}
}

func (gdb *GormSQL) SignUp(user models.User) error {

	if err := gdb.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (gdb *GormSQL) Login(ctx echo.Context) (models.User, error) {
	var user models.User

	ctx.Bind(&user)

	err := gdb.DB.Where("username = ? AND password = ?",
		user.Username,
		user.Password,
	).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (gdb *GormSQL) GetBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := gdb.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (gdb *GormSQL) GetBlogByID(uuid string) (models.Blog, error) {
	var blog models.Blog

	if err := gdb.DB.First(&blog, "uuid = ?", uuid).Error; err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func (gdb *GormSQL) NewBlog(item models.Blog) error {

	if err := gdb.DB.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

func (gdb *GormSQL) UpdateBlog(new models.Blog) error {

	if err := gdb.DB.Model(&new).Where("uuid = ?", new.UUID).
		Updates(models.Blog{Title: new.Title, Content: new.Content, CategoryID: new.CategoryID}).
		Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (gdb *GormSQL) DeleteBlog(uuid string) error {
	var item models.Blog

	if err := gdb.DB.Unscoped().Where("uuid = ?", uuid).Delete(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (gdb *GormSQL) GetBlogByCategory(category_id uint) ([]models.Blog, error) {
	var items []models.Blog

	if err := gdb.DB.Model(&models.Blog{}).Where("category_id = ?", category_id).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (gdb *GormSQL) NewCategory(category models.Category) error {
	if err := gdb.DB.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func (gdb *GormSQL) GetBlogByTitle(item_name string) (models.Blog, error) {
	var item models.Blog

	if err := gdb.DB.Where("title = ?", item_name).First(&item).Error; err != nil {
		return models.Blog{}, err
	}
	return item, nil
}
