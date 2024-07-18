package database

import (
	"database/sql"

	"github.com/ThalesLoreto/go-graphql/graph/model"
	"github.com/google/uuid"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (c *CategoryDB) CreateCategory(name string, description *string) (*model.Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return &model.Category{}, err
	}
	return &model.Category{ID: id, Name: name, Description: description}, nil
}

func (c *CategoryDB) FindByCourseID(id string) (*model.Category, error) {
	var category model.Category
	err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryDB) FindAll() ([]*model.Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
