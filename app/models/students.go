package models

import "student_attendance/database"

// Students 学生模型
type Students struct {
	Id        string `json:"id"`        //主键ID
	StudentId int    `json:"studentId"` //学号
	Name      string `json:"name"`      //姓名
	Grade     int    `json:"grade"`     //年级
	ClassID   int    `json:"classId"`   //班级
	Gender    int    `json:"gender"`    //性别
	Status    int    `json:"status"`    //状态
}

// AddStudent 添加学生
func AddStudent(student Students) (id string, err error) {
	result := database.GetMySQL().Create(&student) //  这里的DB变量是 database 包里定义的，Create 函数是 gorm包的创建数据API

	if result.Error != nil {
		err = result.Error
	}

	return student.Id, err // 返回新建数据的id 和 错误信息，在控制器里接收
}

// FindAllStudent 查询所有学生
func FindAllStudent() (student []Students, err error) {
	result := database.GetMySQL().Find(&student)

	if result.Error != nil {
		err = result.Error
	}

	return student, err
}

// FindStudentByGradeAndStudentId 根据年级和学号查询学生
func FindStudentByGradeAndStudentId(gradeId, studentId int) (student Students, err error) {
	result := database.GetMySQL().Where("grade = ? AND student_id = ?", gradeId, studentId).Find(&student)

	if result.Error != nil {
		err = result.Error
	}

	return student, err
}

// FindStudentByGradeAndClassId 根据年级和班级查询学生
func FindStudentByGradeAndClassId(gradeId, classId int) (student []Students, err error) {
	result := database.GetMySQL().Where("grade = ? AND class_id = ?", gradeId, classId).Find(&student)

	if result.Error != nil {
		err = result.Error
	}

	return student, err
}
