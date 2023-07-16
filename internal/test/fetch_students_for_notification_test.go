package handler

import (
	"bytes"
	"class-management/internal/handler"
	"class-management/internal/mocks"
	"class-management/internal/service/teacher"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetrieveNotifications(t *testing.T) {
	// Create a new instance of the teacher service and mock the dependencies
	teacherRepo := &mocks.MockTeacherRepo{}
	studentRepo := &mocks.MockStudentRepo{}
	teacherStudentRepo := &mocks.MockTeacherStudentsRepo{}
	teacherService := teacher.NewTeacherService(teacherRepo, studentRepo, teacherStudentRepo)
	teacherHandler := handler.NewTeacherHandler(teacherService)

	// Test case: Empty request body
	t.Run("EmptyRequestBody_BadRequest", func(t *testing.T) {
		// Prepare the request URL and body
		reqURL := "/api/retrievefornotifications"
		reqBody := []byte(`{}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.FetchStudentsForNotification)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Empty teacher email
	t.Run("EmptyTeacherEmail_BadRequest", func(t *testing.T) {
		// Prepare the request URL and body
		reqURL := "/api/retrievefornotifications"
		reqBody := []byte(`{"teacher": "", "notification": "Hello students!"}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.FetchStudentsForNotification)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Invalid teacher email
	t.Run("InvalidTeacherEmail_BadRequest", func(t *testing.T) {
		// Prepare the request URL and body
		reqURL := "/api/retrievefornotifications"
		reqBody := []byte(`{"teacher": "invalid_email", "notification": "Hello students!"}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.FetchStudentsForNotification)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Empty notification
	t.Run("EmptyNotification_BadRequest", func(t *testing.T) {
		// Prepare the request URL and body
		reqURL := "/api/retrievefornotifications"
		reqBody := []byte(`{"teacher": "teacherken@gmail.com", "notification": ""}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.FetchStudentsForNotification)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Missing teacher email
	t.Run("MissingTeacherEmail_BadRequest", func(t *testing.T) {
		// Prepare the request URL and body
		reqURL := "/api/retrievefornotifications"
		reqBody := []byte(`{"notification": "Hello students!"}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.FetchStudentsForNotification)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})

	// Test case: Missing notification message
	t.Run("MissingNotificationMessage_BadRequest", func(t *testing.T) {
		// Prepare the request URL and body
		reqURL := "/api/retrievefornotifications"
		reqBody := []byte(`{"teacher": "teacherken@gmail.com"}`)

		// Create a new HTTP request
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.FetchStudentsForNotification)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})
}
