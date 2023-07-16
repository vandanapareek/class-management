package handler

import (
	"class-management/errors"
	"class-management/internal/dto"
	"class-management/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// CommonStudentsOfTeachers handler retrieves a list of students common to a given list of teachers.
//It expects the teacher email query param containing the email address(es) and return error if any or returns the list of common students if successful.
func (th teacherHandler) CommonStudentsOfTeachers(writer http.ResponseWriter, request *http.Request) {

	//validate params
	allTeachers, isOk := request.URL.Query()["teacher"]
	if !isOk || len(allTeachers) == 0 {
		errors.JSONError(writer, errors.ErrMissingTeacherParam, http.StatusUnprocessableEntity)
		return
	}

	//validate given teachers
	for _, teacher := range allTeachers {
		if !utils.IsEmailValid(teacher) {
			errors.JSONError(writer, errors.ErrInvalidTeacherEmail, http.StatusUnprocessableEntity)
			return
		}
	}

	//fetch common students of given teachers
	students, err := th.service.CommonStudentsOfTeachers(allTeachers)
	if err != nil {
		fmt.Println("err in getting common students", err)
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	response := dto.CommonStudentsResponse{
		Students: students,
	}

	//prepare output
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}
