package repository

import (
	"context"

	"github.com/phat-ngoc-anh/backend/internal/domain/entity"
)

type StudentRepository interface {
	GetAllStudentQuery(c context.Context) []entity.Student
}
