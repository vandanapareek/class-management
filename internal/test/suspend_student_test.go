package handler

import (
	"bytes"
	"class-management/internal/handler"
	"class-management/internal/mocks"
	"class-management/internal/models"
	"class-management/internal/service/teacher"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuspendStudents(t *testing.T) {

	// Create a new instance of the teacher service and mock the dependencies
	teacherRepo := &mocks.MockTeacherRepo{}
	studentRepo := &mocks.MockStudentRepo{}
	teacherStudentRepo := &mocks.MockTeacherStudentsRepo{}
	teacherService := teacher.NewTeacherService(teacherRepo, studentRepo, teacherStudentRepo)
	teacherHandler := handler.NewTeacherHandler(teacherService)

	// Test case 1: Suspend an existing student successfully
	t.Run("SuspendExistingStudent_Success", func(t *testing.T) {
		// Register the student
		studentEmail := "student1@example.com"
		mockStudent := &models.Student{
			ID:     1,
			Email:  studentEmail,
			Status: models.StatusActive,
		}
		studentRepo.CreateStudentFn = nil

		// Mock the CreateStudent and return the mockStudent
		createdStudent, err := studentRepo.CreateStudent(mockStudent)
		if err != nil {
			log.Println("Failed to register student:", err)
		} else {
			log.Println("Student registered successfully:", createdStudent.ID)
		}

		// Prepare the request payload
		payload := []byte(fmt.Sprintf(`{"student": "%s"}`, studentEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/suspend", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.SuspendStudent)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusNoContent {
			t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, rr.Code)
		}
	})

	// Test case: Empty request body
	t.Run("EmptyRequestBody_BadRequest", func(t *testing.T) {
		// Prepare the request URL and body
		reqURL := "/api/suspend"
		reqBody := []byte(`{}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.SuspendStudent)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	//Test case 2: Suspend a non-existing student
	t.Run("SuspendNonExistingStudent_NotFound", func(t *testing.T) {
		studentEmail := "nonexistingstudent@gmail.com"

		studentRepo.GetStudentByEmailFn = func(email string) (*models.Student, error) {
			return nil, errors.New("student not found")
		}

		// Prepare the request payload
		payload := []byte(fmt.Sprintf(`{"student": "%s"}`, studentEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/suspend", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.SuspendStudent)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case 3: Suspend with empty student email
	t.Run("SuspendWithEmptyStudent_BadRequest", func(t *testing.T) {
		studentEmail := ""

		// Prepare the request payload
		payload := []byte(fmt.Sprintf(`{"student": "%s"}`, studentEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/suspend", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.SuspendStudent)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case 4: Suspend a student with an invalid email
	t.Run("SuspendInvalidStudent_InvalidRequest", func(t *testing.T) {
		studentEmail := "invalid_email"

		// Prepare the request payload
		payload := []byte(fmt.Sprintf(`{"student": "%s"}`, studentEmail))

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/suspend", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.SuspendStudent)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case 5: Suspend a student with missing email
	t.Run("SuspendMissingStudent_MissingRequest", func(t *testing.T) {
		payload := []byte(`{}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "/api/suspend", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.SuspendStudent)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})
}
