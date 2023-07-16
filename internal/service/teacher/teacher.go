package teacher

import (
	"class-management/errors"
	"class-management/internal/dto"
	"class-management/internal/models"
	"class-management/internal/utils"
	"log"
	"regexp"
)

type TeacherService interface {
	RegisterStudents(dto.RegisterStudentsRequest) error
	SuspendStudent(dto.SuspendRequest) error
	CommonStudentsOfTeachers([]string) ([]string, error)
	FetchStudentsForNotification(dto.FetchStudentsForNotificationRequest) ([]string, error)
	RegisterTeachers(dto.RegisterTeachersRequest) error
}

type teacherService struct {
	teacherRepo        models.TeacherRepo
	studentRepo        models.StudentRepo
	teacherStudentRepo models.TeacherStudentRepo
}

func NewTeacherService(teacherRepo models.TeacherRepo, studentRepo models.StudentRepo, teacherStudentRepo models.TeacherStudentRepo) TeacherService {
	return &teacherService{
		teacherRepo:        teacherRepo,
		studentRepo:        studentRepo,
		teacherStudentRepo: teacherStudentRepo,
	}
}

//RegisterStudents service for registering multiple students with a teacher. A student can also be registered to multiple teachers.
func (ts *teacherService) RegisterStudents(req dto.RegisterStudentsRequest) error {
	teacherDetails, err := ts.teacherRepo.GetTeacherByEmail(req.Teacher)
	if err != nil {
		return err
	}

	if teacherDetails == nil {
		return errors.ErrTeacherNotExists
	}

	for _, studentEmail := range req.Students {
		if utils.IsEmailValid(studentEmail) {
			tx := models.DB.Begin()
			studentDetails, err := ts.studentRepo.GetStudentByEmail(studentEmail)
			if err != nil {
				log.Printf("GetStudentByEmail error for email: %s", studentEmail)
				tx.Rollback()
				return err
			}
			if studentDetails == nil {
				studentObj := &models.Student{
					Email:  studentEmail,
					Status: "Active",
				}

				studentDetails, err = ts.studentRepo.CreateStudent(studentObj)
				if err != nil {
					log.Printf("CreateStudent error for email: %s", studentEmail)
					tx.Rollback()
					return err
				}
			}

			teacherStudentDetails, err := ts.teacherStudentRepo.IsStudentRegisteredForTeacher(teacherDetails.ID, studentDetails.ID)
			if err != nil {
				log.Printf("IsStudentRegisteredForTeacher error for email: %s", studentEmail)
				tx.Rollback()
				return err
			}

			if teacherStudentDetails == nil {
				teacherStudentObj := &models.TeacherStudent{
					TeacherID: teacherDetails.ID,
					StudentID: studentDetails.ID,
				}
				err := ts.teacherStudentRepo.CreateTeacherStudent(teacherStudentObj)
				if err != nil {
					log.Printf("CreateTeacherStudent error for email: %s", studentEmail)
					tx.Rollback()
					return err
				}
			}
			tx.Commit()
		} else {
			log.Printf("Invalid email: %s", studentEmail)
		}
	}

	return nil
}

//SuspendStudent service to suspend a student.
func (ts *teacherService) SuspendStudent(req dto.SuspendRequest) error {
	studentDetails, err := ts.studentRepo.GetStudentByEmail(req.Student)
	if err != nil {
		return err
	}
	if studentDetails == nil {
		return errors.ErrStudentNotExists
	}

	studentDetails.Status = models.StatusSuspended

	err = ts.studentRepo.UpdateStudentStatus(studentDetails)
	if err != nil {
		return err
	}
	return nil
}

// CommonStudentsOfTeachers service retrieves a list of students common to a given list of teachers.
func (ts *teacherService) CommonStudentsOfTeachers(teachers []string) ([]string, error) {
	//validate given teachers
	for _, teacher := range teachers {
		teacherDetails, err := ts.teacherRepo.GetTeacherByEmail(teacher)
		if err != nil {
			return nil, err
		}
		if teacherDetails == nil {
			return nil, errors.ErrTeacherNotExists
		}
	}

	students, err := ts.teacherStudentRepo.GetCommonStudents(teachers)
	if err != nil {
		return nil, err
	}
	return students, nil
}

//FetchStudentsForNotification service retrieve a list of students who can receive a given notification.
func (ts *teacherService) FetchStudentsForNotification(req dto.FetchStudentsForNotificationRequest) ([]string, error) {
	teacherDetails, err := ts.teacherRepo.GetTeacherByEmail(req.Teacher)
	if err != nil {
		return nil, err
	}

	if teacherDetails == nil {
		return nil, errors.ErrTeacherNotExists
	}

	//fetch students mentioned in the notification
	namedStudents := fetchMentionedStudents(req.Notification)

	registeredStudent, err := ts.teacherStudentRepo.GetAllStudentsByTeacher(req.Teacher)
	if err != nil {
		return nil, err
	}

	allStudents := filterStudents(namedStudents, registeredStudent)
	return allStudents, nil
}

//Retrieve email ids from given text
func fetchMentionedStudents(text string) []string {
	regexPattern := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`
	regexRule := regexp.MustCompile(regexPattern)

	//fetch all email address from notification text
	emails := regexRule.FindAllString(text, -1)
	return emails

}

//filter unique students from notification text and registered students of a teacher
func filterStudents(namedStudents []string, registeredStudent []models.Student) []string {
	studentSet := make(map[string]bool)
	students := []string{}

	for _, student := range namedStudents {
		studentSet[student] = true
	}

	for _, student := range registeredStudent {
		studentSet[student.Email] = true
	}

	for student := range studentSet {
		students = append(students, student)
	}

	return students

}

//RegisterTeachers service registers single or multiple teachers.
func (ts *teacherService) RegisterTeachers(req dto.RegisterTeachersRequest) error {

	for _, email := range req.Teachers {
		if utils.IsEmailValid(email) {
			teacherDetails, err := ts.teacherRepo.GetTeacherByEmail(email)
			if err != nil {
				return err
			}

			if teacherDetails == nil {
				teacherObj := &models.Teacher{
					Email: email,
				}
				_, err := ts.teacherRepo.CreateTeacher(teacherObj)
				if err != nil {
					return err
				}
			}

		} else {
			log.Printf("Invalid email: %s", email)
		}
	}

	return nil
}
