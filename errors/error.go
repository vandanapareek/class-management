package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err ApiError) Error() string {
	return fmt.Sprintf("error_code = %v, error_message = %v", err.Code, err.Message)
}

func CreateError(code int, message string) ApiError {
	return ApiError{Code: code, Message: message}
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

var Success = ApiError{Code: 200, Message: "success"}

var ErrAccountNotFound = ApiError{Code: 422, Message: "Invalid account!"}
var ErrDecodingRequest = ApiError{Code: 422, Message: "Please pass valid parameters!"}
var ErrTeacherRequired = ApiError{Code: 422, Message: "A valid teacher email is required!"}
var ErrTeachersRequired = ApiError{Code: 422, Message: "No teachers provided for registration!"}
var ErrStudentRequired = ApiError{Code: 422, Message: "A valid student email is required!"}
var ErrStudentsRequired = ApiError{Code: 422, Message: "No students provided for registration!"}
var ErrTeacherNotExists = ApiError{Code: 422, Message: "Teacher's email you provided doesn't exists!"}
var ErrStudentNotExists = ApiError{Code: 422, Message: "Student's email you provided doesn't exists!"}
var ErrMissingTeacherParam = ApiError{Code: 422, Message: "Teacher parameter is missing in the request!"}
var ErrNotificationRequired = ApiError{Code: 422, Message: "Please enter notification text!"}
var ErrInvalidTeacherEmail = ApiError{Code: 422, Message: "Please enter valid teacher's email!"}
var ErrInvalidStudentEmail = ApiError{Code: 422, Message: "Please enter valid student's email!"}
