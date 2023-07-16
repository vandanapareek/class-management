package handler

import (
	"class-management/internal/service/teacher"
)

type teacherHandler struct {
	service teacher.TeacherService
}

func NewTeacherHandler(s teacher.TeacherService) *teacherHandler {
	return &teacherHandler{
		service: s,
	}
}
