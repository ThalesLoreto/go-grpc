package database

import "github.com/ThalesLoreto/go-graphql/graph/model"

type ICategory interface {
	CreateCategory(name string, description *string) (*model.Category, error)
	FindByCourseID(id string) (*model.Category, error)
	FindAll() ([]*model.Category, error)
}

type ICourse interface {
	CreateCourse(title string, description *string, categoryID string) (*model.Course, error)
	FindByCategoryID(id string) ([]*model.Course, error)
	FindAll() ([]*model.Course, error)
}
