package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	Email     string        `gorm:"unique;not null" json:"email"`
	Status    StatusStudent `gorm:"default:ACTIVE;not null" json:"status"`
	CreatedAt time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAT time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Student) TableName() string {
	return "students"
}

type studentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentRepo {
	return &studentRepo{db}
}

type StudentRepo interface {
	GetStudentByEmail(email string) (*Student, error)
	CreateStudent(student *Student) (*Student, error)
	UpdateStudentStatus(student *Student) error
}

//status of students
type StatusStudent string

const (
	StatusActive    StatusStudent = "ACTIVE"
	StatusSuspended StatusStudent = "SUSPENDED"
	StatusGraduated StatusStudent = "GRADUATED"
)

//Get student detail by its email id
func (s *studentRepo) GetStudentByEmail(email string) (*Student, error) {
	var details Student
	res := s.db.Where("email = ?", email).First(&details)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, res.Error
		}
	}
	return &details, nil
}

//create a new student
func (s *studentRepo) CreateStudent(student *Student) (*Student, error) {
	err := s.db.Create(student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

//Update student's status
func (s *studentRepo) UpdateStudentStatus(student *Student) error {
	err := s.db.Save(student).Error
	if err != nil {
		return err
	}
	return nil
}
