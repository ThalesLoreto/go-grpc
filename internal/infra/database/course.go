package database

import (
	"database/sql"

	"github.com/ThalesLoreto/go-graphql/graph/model"
	"github.com/google/uuid"
)

type CourseDB struct {
	db *sql.DB
}

func NewCourseDB(db *sql.DB) *CourseDB {
	return &CourseDB{db: db}
}

func (c *CourseDB) CreateCourse(title string, description *string, categoryID string) (*model.Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, title, description, category_id) VALUES ($1, $2, $3, $4)", id, title, description, categoryID)
	if err != nil {
		return &model.Course{}, err
	}
	return &model.Course{
		ID:          id,
		Title:       title,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *CourseDB) FindByCategoryID(id string) ([]*model.Course, error) {
	rows, err := c.db.Query("SELECT id, title, description FROM courses WHERE category_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course
	for rows.Next() {
		var course model.Course
		err := rows.Scan(&course.ID, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (c *CourseDB) FindAll() ([]*model.Course, error) {
	rows, err := c.db.Query("SELECT id, title, description FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course
	for rows.Next() {
		var course model.Course
		err := rows.Scan(&course.ID, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}
