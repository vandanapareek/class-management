package main

import (
	"class-management/internal/handler"
	"class-management/internal/models"
	"class-management/internal/service/teacher"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//connect to db
	db, err := getDb()
	if err != nil {
		log.Fatal(err)
	}
	models.DB = db

	//close the db connection when application exits
	dbClose, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbClose.Close()

	teacherRepo := models.NewTeacherRepo(db)
	studentRepo := models.NewStudentRepo(db)
	teacherStudentRepo := models.NewTeacherStudentRepo(db)
	teacherService := teacher.NewTeacherService(teacherRepo, studentRepo, teacherStudentRepo)
	teacherHandler := handler.NewTeacherHandler(teacherService)

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/", homeHandler).Methods("GET")

	router.HandleFunc("/register", teacherHandler.RegisterStudents).Methods(http.MethodPost)
	router.HandleFunc("/suspend", teacherHandler.SuspendStudent).Methods(http.MethodPost)
	router.HandleFunc("/commonstudents", teacherHandler.CommonStudentsOfTeachers).Methods(http.MethodGet)
	router.HandleFunc("/retrievefornotifications", teacherHandler.FetchStudentsForNotification).Methods(http.MethodPost)
	router.HandleFunc("/registerteachers", teacherHandler.RegisterTeachers).Methods(http.MethodPost)

	log.Println("Application has started. Listening port is 8080")
	http.ListenAndServe(":8080", router)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(os.Getenv("DB_USER"))
	log.Println(os.Getenv("MYSQL_ROOT_PASSWORD"))
	fmt.Fprintf(w, "Hello, World!")
}

func getDb() (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
