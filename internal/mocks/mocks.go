package mocks

import (
	"class-management/internal/models"
	"errors"
)

// MockTeacherRepo is a mock implementation of the TeacherRepo interface
type MockTeacherRepo struct {
	TeacherByEmailFn func(email string) (*models.Teacher, error)
	CreateTeacherFn  func(teacher *models.Teacher) (*models.Teacher, error)
}

func (m *MockTeacherRepo) GetTeacherByEmail(email string) (*models.Teacher, error) {
	if m.TeacherByEmailFn != nil {
		return m.TeacherByEmailFn(email)
	}

	// Default behavior: Return a mock teacher with the provided email
	mockTeacher := &models.Teacher{
		ID:    1,
		Email: email,
		// Set other properties as needed
	}

	// Return the mock teacher and nil error
	return mockTeacher, nil
}

func (m *MockTeacherRepo) CreateTeacher(teacher *models.Teacher) (*models.Teacher, error) {
	createdTeacher := &models.Teacher{
		ID:    1,
		Email: teacher.Email,
	}

	// Return the created teacher and nil error
	return createdTeacher, nil
}

// MockStudentRepo is a mock implementation of the StudentRepo interface
type MockStudentRepo struct {
	CreateStudentFn       func(student *models.Student) (*models.Student, error)
	GetStudentByEmailFn   func(email string) (*models.Student, error)
	UpdateStudentStatusFn func(student *models.Student) error
}

func (m *MockStudentRepo) CreateStudent(student *models.Student) (*models.Student, error) {
	if m.CreateStudentFn != nil {
		return m.CreateStudentFn(student)
	}

	// Default behavior: Return the provided student as is
	return student, nil
}

func (m *MockStudentRepo) GetStudentByEmail(email string) (*models.Student, error) {
	if m.GetStudentByEmailFn != nil {
		return m.GetStudentByEmailFn(email)
	}

	// Default behavior: Return a mock student with the provided email
	mockStudent := &models.Student{
		ID:     1,
		Email:  email,
		Status: models.StatusActive,
	}

	// Return the mock student and nil error
	return mockStudent, nil
}

func (m *MockStudentRepo) UpdateStudentStatus(student *models.Student) error {
	if m.UpdateStudentStatusFn != nil {
		return m.UpdateStudentStatusFn(student)
	}

	// Default behavior: Return nil error
	return nil
}

// MockTeacherStudentsRepo is a mock implementation of the TeacherStudentsRepo interface
type MockTeacherStudentsRepo struct {
	CreateTeacherStudentFn          func(*models.TeacherStudent) error
	IsStudentRegisteredForTeacherFn func(uint, uint) (*models.TeacherStudent, error)
	GetAllStudentsByTeacherFn       func(string) ([]models.Student, error)
	GetCommonStudentsFn             func([]string) ([]string, error)
}

func (m *MockTeacherStudentsRepo) CreateTeacherStudent(student *models.TeacherStudent) error {
	if m.CreateTeacherStudentFn != nil {
		return m.CreateTeacherStudentFn(student)
	}

	// Default behavior: Return an error
	return errors.New("failed to create teacher student")
}

func (m *MockTeacherStudentsRepo) IsStudentRegisteredForTeacher(teacherID uint, studentID uint) (*models.TeacherStudent, error) {
	if m.IsStudentRegisteredForTeacherFn != nil {
		return m.IsStudentRegisteredForTeacherFn(teacherID, studentID)
	}

	// Default behavior: Return a mock TeacherStudent object or an error based on your test case scenario
	// Replace the code below with the desired behavior of the mock
	mockTeacherStudent := &models.TeacherStudent{
		ID:        1,
		TeacherID: teacherID,
		StudentID: studentID,
	}

	// Return the mock TeacherStudent object and nil error
	return mockTeacherStudent, nil
}

func (m *MockTeacherStudentsRepo) GetAllStudentsByTeacher(teacherEmail string) ([]models.Student, error) {
	if m.GetAllStudentsByTeacherFn != nil {
		return m.GetAllStudentsByTeacherFn(teacherEmail)
	}

	// Default behavior: Return an empty slice of students
	return []models.Student{}, nil
}

func (m *MockTeacherStudentsRepo) GetCommonStudents(teacherEmails []string) ([]string, error) {
	if m.GetCommonStudentsFn != nil {
		return m.GetCommonStudentsFn(teacherEmails)
	}

	// Default behavior: Return an empty slice of common students
	return []string{}, nil
}
