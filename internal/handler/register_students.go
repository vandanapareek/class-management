package handler

import (
	"class-management/errors"
	"class-management/internal/dto"
	"class-management/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//RegisterStudents handler registers single or multiple students with a teacher.A student can also be registered to multiple teachers.
//It expects the teacher email and student(s) email as input and return error if any.
func (th teacherHandler) RegisterStudents(writer http.ResponseWriter, request *http.Request) {

	//validate params
	registerReq, err := processRegisterParams(request)
	if err != nil {
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	//register all students
	err = th.service.RegisterStudents(registerReq)
	if err != nil {
		fmt.Println("err in registration process", err)
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	//prepare output
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(http.StatusNoContent)
	json.NewEncoder(writer).Encode("Students has been registered successfully.")
}

//validate input parameters
func processRegisterParams(request *http.Request) (dto.RegisterStudentsRequest, error) {
	var params dto.RegisterStudentsRequest

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		return params, errors.ErrDecodingRequest
	}

	if params.Teacher == "" {
		return params, errors.ErrTeacherRequired
	}

	if !utils.IsEmailValid(params.Teacher) {
		return params, errors.ErrInvalidTeacherEmail
	}

	if len(params.Students) == 0 {
		return params, errors.ErrStudentsRequired
	}

	return params, nil
}
