package service

import (
	"context"

	"github.com/ThalesLoreto/go-grpc/internal/infra/database"
	"github.com/ThalesLoreto/go-grpc/internal/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (cs *CategoryService) CreateCategory(ctx context.Context,
	in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := cs.CategoryDB.CreateCategory(in.Name, &in.Description)
	if err != nil {
		return nil, err
	}
	res := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: *category.Description,
	}
	return res, nil
}

func (cs *CategoryService) ListCategories(ctx context.Context,
	in *emptypb.Empty) (*pb.CategoryList, error) {
	categories, err := cs.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesResponse []*pb.Category
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: *category.Description,
		})
	}
	return &pb.CategoryList{Categories: categoriesResponse}, nil
}

func (cs *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := cs.CategoryDB.FindByCourseID(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: *category.Description,
	}, nil
}
