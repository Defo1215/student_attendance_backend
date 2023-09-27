package models

import "student_attendance/database"

// Course 课程模型
type Course struct {
	Id    string `json:"id"`    //课程ID(主键)
	Name  string `json:"name"`  //课程名
	Grade int    `json:"grade"` //年级
}

// AddCourse 添加课程
func AddCourse(course Course) (id string, err error) {
	result := database.GetMySQL().Create(&course)

	if result.Error != nil {
		err = result.Error
	}

	return course.Id, err
}

// FindAllCourse 查询所有课程
func FindAllCourse() (course []Course, err error) {
	result := database.GetMySQL().Find(&course)

	if result.Error != nil {
		err = result.Error
	}

	return course, err
}

// FindCourseByGrade 根据年级查询课程
func FindCourseByGrade(grade int) (course []Course, err error) {
	result := database.GetMySQL().Where("grade = ?", grade).Find(&course)

	if result.Error != nil {
		err = result.Error
	}

	return course, err
}
