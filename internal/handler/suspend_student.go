package handler

import (
	"class-management/errors"
	"class-management/internal/dto"
	"class-management/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//SuspendStudent handler suspends a student.
//It expects the student email as input param and return error if any
func (th teacherHandler) SuspendStudent(writer http.ResponseWriter, request *http.Request) {

	//validate params
	suspendReq, err := processSuspendParams(request)
	if err != nil {
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	//Suspend Student
	err = th.service.SuspendStudent(suspendReq)
	if err != nil {
		fmt.Println("err in suspension process", err)
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	//prepare output
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(http.StatusNoContent)
	json.NewEncoder(writer).Encode("Students has been suspended successfully.")
}

//validate input parameters
func processSuspendParams(request *http.Request) (dto.SuspendRequest, error) {
	var params dto.SuspendRequest

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		return params, errors.ErrDecodingRequest
	}

	if params.Student == "" {
		return params, errors.ErrStudentRequired
	}

	if !utils.IsEmailValid(params.Student) {
		return params, errors.ErrInvalidStudentEmail
	}

	return params, nil
}
