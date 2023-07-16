package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAT time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Teacher) TableName() string {
	return "teachers"
}

type teacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return &teacherRepo{db}
}

type TeacherRepo interface {
	CreateTeacher(teacher *Teacher) (*Teacher, error)
	GetTeacherByEmail(email string) (*Teacher, error)
}

//Create a new teacher
func (s *teacherRepo) CreateTeacher(teacher *Teacher) (*Teacher, error) {
	err := s.db.Create(teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

//Get teacher's detail by its email id
func (t *teacherRepo) GetTeacherByEmail(email string) (*Teacher, error) {
	var details Teacher
	res := t.db.Where("email = ?", email).First(&details)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, res.Error
		}
	}
	return &details, nil
}
