package repository

import (
	"excel-processor/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(student entity.Student) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository  {
	return &studentRepository{db: db}
}

func (r *studentRepository) Create(student entity.Student) error  {
	fmt.Printf("💾 Saving to DB: %s\n", student.Name)

	result := r.db.Create(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}