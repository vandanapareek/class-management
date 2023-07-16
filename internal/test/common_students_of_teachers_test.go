package handler

import (
	"class-management/internal/handler"
	"class-management/internal/mocks"
	"class-management/internal/service/teacher"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetCommonStudents(t *testing.T) {
	// Create a new instance of the teacher service and mock the dependencies
	teacherRepo := &mocks.MockTeacherRepo{}
	studentRepo := &mocks.MockStudentRepo{}
	teacherStudentRepo := &mocks.MockTeacherStudentsRepo{}
	teacherService := teacher.NewTeacherService(teacherRepo, studentRepo, teacherStudentRepo)
	teacherHandler := handler.NewTeacherHandler(teacherService)

	// Test case: No teacher found in query parameters
	t.Run("GetCommonStudents_NoTeacher", func(t *testing.T) {
		// Prepare the request URL and parameters with no teacher
		reqURL := "/api/commonstudents"
		reqParams := url.Values{}

		// Create a new HTTP request
		req, err := http.NewRequest("GET", reqURL+"?"+reqParams.Encode(), nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP test recorder
		rr := httptest.NewRecorder()

		// Handle the request
		handler := http.HandlerFunc(teacherHandler.CommonStudentsOfTeachers)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if rr.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnprocessableEntity, rr.Code)
		}
	})
}
