package service

import (
	"context"

	"github.com/phat-ngoc-anh/backend/internal/domain/model"
)

type StudentService interface {
	GetAllStudent(ctx context.Context) []model.Student
}
