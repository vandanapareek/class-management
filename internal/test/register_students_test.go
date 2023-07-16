package handler

import (
	"bytes"
	"class-management/internal/handler"
	"class-management/internal/mocks"
	"class-management/internal/models"
	"class-management/internal/service/teacher"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func toJSON(data interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}

func TestRegisterStudent(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Set up the mock expectations
	mock.ExpectExec(".*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Create a new gorm.DB instance using the mockDB connection
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create gorm.DB instance: %v", err)
	}

	// Set the DB instance in the model package
	models.DB = db

	// Create a new instance of the teacher service and mock the dependencies
	teacherRepo := &mocks.MockTeacherRepo{}
	studentRepo := &mocks.MockStudentRepo{}
	teacherStudentRepo := &mocks.MockTeacherStudentsRepo{}
	teacherService := teacher.NewTeacherService(teacherRepo, studentRepo, teacherStudentRepo)
	teacherHandler := handler.NewTeacherHandler(teacherService)

	// Test case: Registering one student successfully
	t.Run("RegisterStudents_Success", func(t *testing.T) {
		// Reset the expectations on the mock DB
		mock.ExpectationsWereMet()

		// Prepare the request payload
		teacherEmail := "teacher@example.com"
		studentEmails := []string{"student1@example.com", "student2@example.com"}
		payload := []byte(fmt.Sprintf(`{"teacher": "%s", "students": %s}`, teacherEmail, toJSON(studentEmails)))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.RegisterStudents)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusNoContent {
			t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, rr.Code)
		}
	})

	// Test case: Registering multiple students successfully
	t.Run("RegisterMultipleStudents_Success", func(t *testing.T) {
		// Prepare the request payload
		teacherEmail := "t1@gmail.com"
		_, err := teacherRepo.CreateTeacher(&models.Teacher{
			Email: teacherEmail,
		})
		if err != nil {
			t.Fatal(err)
		}

		studentEmails := []string{"studentjon@gmail.com", "studenthon@gmail.com"}
		payload := []byte(fmt.Sprintf(`{"teacher": "%s", "students": ["%s", "%s"]}`, teacherEmail, studentEmails[0], studentEmails[1]))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.RegisterStudents)
		handler.ServeHTTP(rr, req)
		// Check the response status code
		if rr.Code != http.StatusNoContent {
			t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, rr.Code)
		}
	})

	// Test case: Empty request body
	t.Run("EmptyRequestBody_BadRequest", func(t *testing.T) {
		// Prepare the request URL and body
		reqURL := "/api/register"
		reqBody := []byte(`{}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.RegisterStudents)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Registering with missing teacher email
	t.Run("RegisterMissingTeacherEmail_BadRequest", func(t *testing.T) {
		// Prepare the request payload with missing teacher email
		studentEmail := "studentjon@gmail.com"
		payload := []byte(fmt.Sprintf(`{"students": ["%s"]}`, studentEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.RegisterStudents)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Registering with missing student emails
	t.Run("RegisterMissingStudentEmails_BadRequest", func(t *testing.T) {
		// Prepare the request payload
		teacherEmail := "t1@gmail.com"
		_, err := teacherRepo.CreateTeacher(&models.Teacher{
			Email: teacherEmail,
			// Set other properties of the teacher
		})
		if err != nil {
			t.Fatal(err)
		}
		payload := []byte(fmt.Sprintf(`{"teacher": "%s"}`, teacherEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.RegisterStudents)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Registering with empty teacher email
	t.Run("RegisterEmptyTeacherEmail_BadRequest", func(t *testing.T) {
		// Prepare the request payload with empty teacher email
		teacherEmail := ""
		studentEmail := "studentjon@gmail.com"
		payload := []byte(fmt.Sprintf(`{"teacher": "%s", "students": ["%s"]}`, teacherEmail, studentEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.RegisterStudents)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Registering with an invalid teacher email
	t.Run("RegisterInvalidTeacherEmail_BadRequest", func(t *testing.T) {
		// Prepare the request payload with an invalid teacher email
		teacherEmail := "invalid_email"
		studentEmail := "studentjon@gmail.com"
		payload := []byte(fmt.Sprintf(`{"teacher": "%s", "students": ["%s"]}`, teacherEmail, studentEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.RegisterStudents)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Registering with empty student emails
	t.Run("RegisterEmptyStudentEmails_BadRequest", func(t *testing.T) {
		// Prepare the request payload with empty student emails
		teacherEmail := "t1@gmail.com"
		_, err := teacherRepo.CreateTeacher(&models.Teacher{
			Email: teacherEmail,
			// Set other properties of the teacher
		})
		if err != nil {
			t.Fatal(err)
		}
		payload := []byte(fmt.Sprintf(`{"teacher": "%s", "students": []}`, teacherEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.RegisterStudents)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})
}
