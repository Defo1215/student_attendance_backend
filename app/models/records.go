package models

import "student_attendance/database"

// Records 考勤记录模型
type Records struct {
	Id         string `json:"id"`         //主键ID
	Grade      int    `json:"grade"`      //年级
	StudentId  int    `json:"studentId"`  //学号
	CourseId   string `json:"courseId"`   //课程ID
	Status     int    `json:"status"`     //状态
	CreateTime int64  `json:"createTime"` //创建时间
	UpdateTime int64  `json:"updateTime"` //更新时间
}

// RecordsGroupResponse 考勤记录分组返回结构体
type RecordsGroupResponse struct {
	CourseId   string `json:"courseId"`   //课程ID
	CourseName string `json:"courseName"` //课程名
	Grade      int    `json:"grade"`      //年级
	CreateTime int64  `json:"createTime"` //创建时间
}

type RecordsResponse struct {
	Id          string `json:"id"`          //主键ID
	Grade       int    `json:"grade"`       //年级
	StudentId   int    `json:"studentId"`   //学号
	StudentName string `json:"studentName"` //学生姓名
	CourseId    string `json:"courseId"`    //课程ID
	Status      int    `json:"status"`      //状态
	CreateTime  int64  `json:"createTime"`  //创建时间
	UpdateTime  int64  `json:"updateTime"`  //更新时间
}

// AddRecords 添加考勤记录
func AddRecords(records Records) (id string, err error) {
	result := database.GetMySQL().Create(&records)

	if result.Error != nil {
		err = result.Error
	}

	return records.Id, err
}

// UpdateRecord 更新考勤记录
func UpdateRecord(record Records) (err error) {
	result := database.GetMySQL().Save(&record)

	if result.Error != nil {
		err = result.Error
	}

	return err

}

// DeleteRecord 删除考勤记录
func DeleteRecord(record Records) (err error) {
	result := database.GetMySQL().Delete(&record)

	if result.Error != nil {
		err = result.Error
	}

	return err
}

// FindAllTheSameRecord 查询同一日期、年级、课程的记录
func FindAllTheSameRecord() (records []RecordsGroupResponse, err error) {

	sql := `SELECT DISTINCT 
				records.course_id, 
				course.name AS course_name, 
				course.grade, 
				records.create_time
			FROM records
			LEFT JOIN course ON course.id = records.course_id`

	result := database.GetMySQL().Raw(sql).Scan(&records)

	if result.Error != nil {
		err = result.Error
	}

	return records, err
}

// FindByDateAndGradeAndCourseIdAndStatus 根据日期、年级、课程ID、状态查询考勤记录
func FindByDateAndGradeAndCourseIdAndStatus(date int64, grade int, courseId string, Status int) (record []RecordsResponse, err error) {
	result := database.GetMySQL().
		Table("records").
		Select("records.id, records.grade, records.student_id, students.name AS student_name, records.course_id, records.status, records.create_time, records.update_time").
		Where("records.create_time = ? AND records.grade = ? AND records.course_id = ? AND records.status = ?", date, grade, courseId, Status).
		Joins("JOIN students ON students.student_id = records.student_id").
		Find(&record)

	if result.Error != nil {
		err = result.Error
	}

	return record, err
}

// FindNotCheckedStudents 查询指定课程中未签到的学生
func FindNotCheckedStudents(date int64, grade int, courseId string) (students []Students, err error) {
	sql := `SELECT student_id, name
			FROM students
			WHERE student_id NOT IN (
			  SELECT student_id
			  FROM records
			  WHERE grade = ? 
				AND course_id = ?
			 	AND  create_time = ?
			)`

	result := database.GetMySQL().Raw(sql, grade, courseId, date).Scan(&students)

	if result.Error != nil {
		err = result.Error
	}

	return students, err
}

// IsCheck 判断该学生是否已经签到
func IsCheck(grade int, courseId string, studentId int, timestamp int64) (record []Records, err error) {
	result := database.GetMySQL().Where("grade = ? AND AND course_id AND student_id = ? AND create_time > ?", grade, courseId, studentId, timestamp-4*60*60*1000).Find(&record)

	if result.Error != nil {
		err = result.Error
	}

	return record, err
}
