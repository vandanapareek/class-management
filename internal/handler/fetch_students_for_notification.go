package handler

import (
	"class-management/errors"
	"class-management/internal/dto"
	"class-management/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//FetchStudentsForNotification handler retrieve a list of students who can receive a given notification.
//It expects the teacher email and notification text as input and return error if any or returns list of recipients if successful.
func (th teacherHandler) FetchStudentsForNotification(writer http.ResponseWriter, request *http.Request) {
	//validate params
	reqData, err := processStudentsForNotificationsParams(request)
	if err != nil {
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	//Fetch students for notification
	students, err := th.service.FetchStudentsForNotification(reqData)
	if err != nil {
		fmt.Println("err in getting common students", err)
		errors.JSONError(writer, err, http.StatusUnprocessableEntity)
		return
	}

	response := struct {
		Recipients []string `json:"recipients"`
	}{
		Recipients: students,
	}

	//prepare output
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)

}

func processStudentsForNotificationsParams(request *http.Request) (dto.FetchStudentsForNotificationRequest, error) {
	var params dto.FetchStudentsForNotificationRequest

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

	if params.Notification == "" {
		return params, errors.ErrNotificationRequired
	}

	return params, nil
}
