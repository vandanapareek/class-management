package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type TeacherStudent struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TeacherID uint      `gorm:"primaryKey" json:"teacher_id"`
	StudentID uint      `gorm:"primaryKey" json:"student_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAT time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (TeacherStudent) TableName() string {
	return "teacher_students"
}

type teacherStudentRepo struct {
	db *gorm.DB
}

func NewTeacherStudentRepo(db *gorm.DB) TeacherStudentRepo {
	return &teacherStudentRepo{db}
}

type TeacherStudentRepo interface {
	CreateTeacherStudent(*TeacherStudent) error
	IsStudentRegisteredForTeacher(uint, uint) (*TeacherStudent, error)
	GetCommonStudents([]string) ([]string, error)
	GetAllStudentsByTeacher(string) ([]Student, error)
}

//Register a student with a teacher
func (ts *teacherStudentRepo) CreateTeacherStudent(teacherStudentObj *TeacherStudent) error {
	return ts.db.Create(teacherStudentObj).Error
}

//Check if given student is registered with given teacher
func (ts *teacherStudentRepo) IsStudentRegisteredForTeacher(teacherID uint, studentID uint) (*TeacherStudent, error) {
	var details TeacherStudent
	res := ts.db.Where("teacher_id = ?", teacherID).Where("student_id = ?", studentID).First(&details)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, res.Error
		}
	}
	return &details, nil
}

//Get list of common students to a given list of teachers
func (ts *teacherStudentRepo) GetCommonStudents(teachers []string) ([]string, error) {
	var students []string
	query := ts.db.Raw(`SELECT s.email
	                   FROM students AS s 
					   JOIN teacher_students AS ts ON s.id = ts.student_id
					   JOIN teachers AS t ON t.id = ts.teacher_id
					   WHERE t.email IN (?)
					   GROUP BY s.email
					   HAVING COUNT(s.email) = ?
					   `, teachers, len(teachers)).Scan(&students)
	if query.Error != nil {
		return nil, query.Error
	}
	return students, nil
}

//Get all registered students of a teacher
func (ts *teacherStudentRepo) GetAllStudentsByTeacher(teacher string) ([]Student, error) {
	var students []Student
	err := ts.db.
		Model(TeacherStudent{}).
		Select("students.*").
		Joins("JOIN students ON students.id = teacher_students.student_id").
		Joins("JOIN teachers ON teachers.id = teacher_students.teacher_id").
		Where("teachers.email = ?", teacher).
		Not("students.status = ?", StatusSuspended).
		Find(&students).Error

	if err != nil {
		return nil, err
	}

	return students, nil
}
