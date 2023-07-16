package handler

import (
	"class-management/errors"
	"class-management/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

//RegisterTeachers handler registers single or multiple teachers.
func (th teacherHandler) RegisterTeachers(writer http.ResponseWriter, request *http.Request) {

	//validate params
	registerReq, err := processRegisterTeachersParams(request)
	if err != nil {
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	//register all teachers
	err = th.service.RegisterTeachers(registerReq)
	if err != nil {
		fmt.Println("err in registration process")
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	//prepare output
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(http.StatusNoContent)
	json.NewEncoder(writer).Encode("Students has been registered successfully.")
}

func processRegisterTeachersParams(request *http.Request) (dto.RegisterTeachersRequest, error) {
	var params dto.RegisterTeachersRequest

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		return params, errors.ErrDecodingRequest
	}

	if len(params.Teachers) == 0 {
		return params, errors.ErrTeachersRequired
	}

	return params, nil
}
