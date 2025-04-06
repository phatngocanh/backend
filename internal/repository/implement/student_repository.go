package repositoryimplement

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/phat-ngoc-anh/backend/internal/database"
	"github.com/phat-ngoc-anh/backend/internal/domain/entity"
	"github.com/phat-ngoc-anh/backend/internal/repository"
)

type StudentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(db database.Db) repository.StudentRepository {
	return &StudentRepository{db: db}
}

func (repo StudentRepository) GetAllStudentQuery(c context.Context) []entity.Student {
	return []entity.Student{
		{Name: "John"},
		{Name: "Khoa"},
		{Name: "Lindan"},
	}
}
