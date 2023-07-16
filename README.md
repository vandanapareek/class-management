# Class Management

Class Management is a service for managing students and teachers in a school or educational institution. It provides functionality to create, retrieve, update, and delete students and teachers, as well as assign students to teachers.


## Quickstart

### Configuration
1. Make sure you have Docker installed.

2. Clone this repo to your local machine.

```bash
git clone https://github.com/vandanapareek/class-management.git

```
3. Change to the project directory.

```bash
cd class-management
```

4. Create a copy of the .env.example file located in the project root and rename it to .env. 

By default, the .env.example file provides the recommended default values for a local development environment using Docker. Feel free to modify these values according to your specific requirements.

### Run the app using Docker

1. Start the app using Docker Compose:

```bash
docker-compose up --build
```

2. The app will be running on http://localhost:8080.


### Run the test-app using Docker

1. Run the following command to start the test environment:

```bash
docker-compose -f docker-compose.test.yml up  --build
```

2. Wait for the containers to start and initialize.

3. The test-app will be running on http://localhost:8081

4. Execute the tests using the following command:

```bash
docker-compose -f docker-compose.test.yml exec api-test go test -v ./internal/test/...
```

5. You can stop and remove the test environment by running:

```bash
docker-compose -f docker-compose.test.yml down
```

## API Endpoints

### Note: There is one additional API (Teacher Registration) for registering multiple teachers. Use this to feed few teachers before running other APIs.

### 1. Teacher Registration
* Description: This is one additional API for registering multiple teachers. Use this to feed few teachers before running any other API.
* Endpoint: `POST /api/registerteachers`
* Headers: `Content-Type: application/json`
* Success response status: HTTP 204
* Request body example:
```
{
  "teachers":
    [
      "teacher1@gmail.com",
      "teacher2@gmail.com"
    ]
}
```

### 2. Student Registration
* Description: A teacher can register multiple students. A student can also be registered to multiple teachers.
* Endpoint: `POST /api/register`
* Headers: `Content-Type: application/json`
* Success response status: HTTP 204
* Request body example:
```
{
  "teacher": "teacherken@gmail.com"
  "students":
    [
      "studentjon@gmail.com",
      "studenthon@gmail.com"
    ]
}
```

### 3. Get Common Students
* Description: A teacher can register multiple students. A student can also be registered to multiple teachers.
* Endpoint: `GET /api/commonstudents`
* Success response status: HTTP 200
* Request example 1: `GET /api/commonstudents?teacher=teacherken%40gmail.com`
* Success response body 1:
```
{
  "students" :
    [
      "commonstudent1@gmail.com", 
      "commonstudent2@gmail.com",
      "student_only_under_teacher_ken@gmail.com"
    ]
}
```
* Request example 2: `GET /api/commonstudents?teacher=teacherken%40gmail.com&teacher=teacherjoe%40gmail.com`
* Success response body 2:
```
{
  "students" :
    [
      "commonstudent1@gmail.com", 
      "commonstudent2@gmail.com"
    ]
}
```

### 4. Suspend Student
* Description: A teacher can suspend a specified student.
* Endpoint: `POST /api/suspend`
* Headers: `Content-Type: application/json`
* Success response status: HTTP 204
* Request body example:
```
{
  "student" : "studentmary@gmail.com"
}
```

### 5. Retrieve Students for Notification
* Description: A teacher can retrieve a list of students who can receive a given notification.
* Endpoint: `POST /api/retrievefornotifications`
* Headers: `Content-Type: application/json`
* Success response status: HTTP 200
* Request body example 1:
```
{
  "teacher":  "teacherken@gmail.com",
  "notification": "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com"
}
```
* Success response body 1:
```
{
  "recipients":
    [
      "studentbob@gmail.com",
      "studentagnes@gmail.com", 
      "studentmiche@gmail.com"
    ]   
}
```
In the example above, studentagnes@gmail.com and studentmiche@gmail.com can receive the notification from teacherken@gmail.com, regardless whether they are registered to him, because they are @mentioned in the notification text. studentbob@gmail.com however, has to be registered to teacherken@gmail.com.
* Request body example 2:
```
{
  "teacher":  "teacherken@gmail.com",
  "notification": "Hey everybody"
}
```
* Success response body 2:
```
{
  "recipients":
    [
      "studentbob@gmail.com"
    ]   
}
```

## Postman Collection
[Postman Collection](postman_collection.json)