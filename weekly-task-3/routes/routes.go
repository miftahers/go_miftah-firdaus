package routes

import (
	"weekly-task-3/configs"
	controller "weekly-task-3/controllers"
	"weekly-task-3/repositories"
	"weekly-task-3/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	middleware.Logger()

	repo := repositories.NewGorm(db)
	blogS, userS, categoryS := services.NewServices(repo)
	blogHandler := controller.BlogHandler{BlogService: blogS}
	userHandler := controller.UserHandler{UserService: userS}
	categoryHandler := controller.CategoryHandler{CategoryService: categoryS}

	blogs := e.Group("/blogs")

	e.POST("/signup", userHandler.SignUp)
	e.POST("/login", userHandler.Login)
	blogs.POST("", blogHandler.NewBlog, middleware.JWT([]byte(configs.Token)))
	blogs.POST("/category", categoryHandler.NewCategory, middleware.JWT([]byte(configs.Token)))

	blogs.GET("", blogHandler.GetBlogs)
	blogs.GET("/:id", blogHandler.GetBlogById)
	blogs.GET("/category/:category_id", blogHandler.GetBlogByCategory)

	blogs.PUT("/:id", blogHandler.UpdateBlog, middleware.JWT([]byte(configs.Token)))
	blogs.DELETE("/:id", blogHandler.DeleteBlog, middleware.JWT([]byte(configs.Token)))

	return e
}
