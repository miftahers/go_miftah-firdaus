package services

import "weekly-task-3/repositories"

type BlogServ struct {
	repositories.Database
}
type UserServ struct {
	repositories.Database
}
type CategoryServ struct {
	repositories.Database
}

func NewServices(db repositories.Database) (BlogService, UserService, CategoryService) {
	return &BlogServ{Database: db}, &UserServ{Database: db}, &CategoryServ{Database: db}
}
