package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"student_attendance/app/models"
	"student_attendance/app/result"
	"time"
)

// getTodayTimestamp 获取今天00:00的时间戳
func getTodayTimestamp() int64 {
	t := time.Now()
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return today.UnixMilli()
}

// AddRecord 新增记录
func AddRecord(c *gin.Context) {
	var record models.Records

	err := c.ShouldBind(&record) //获取参数

	if err != nil {
		c.JSON(200, result.Fail("获取参数失败")) //返回错误信息
		return
	}

	//判断该学生是否存在
	student, err := models.FindStudentByGradeAndStudentId(record.Grade, record.StudentId)

	if student.Id == "" {
		c.JSON(200, result.Fail("该学生不存在")) //返回错误信息
		return
	}

	//判断该学生是否已经签到
	temp, err := models.IsCheck(record.Grade, record.CourseId, record.StudentId, time.Now().UnixMilli())

	if temp != nil {
		c.JSON(200, result.Fail("该学生已经签到")) //返回错误信息
		return
	}

	record.Id = uuid.New().String()         //生成uuid
	record.CreateTime = getTodayTimestamp() //获取今天00:00的时间戳
	record.UpdateTime = getTodayTimestamp() //获取今天00:00的时间戳

	id, err := models.AddRecords(record) //调用模型层的添加课程方法

	if err != nil {
		c.JSON(200, result.Fail("新增失败")) //返回错误信息
		return
	}

	c.JSON(200, result.Success(id)) //返回成功信息
}

// UpdateRecord 更新记录
func UpdateRecord(c *gin.Context) {
	var record models.Records

	err := c.ShouldBind(&record) //获取参数

	if err != nil {
		c.JSON(200, result.Fail("获取参数失败")) //返回错误信息
		return
	}

	record.UpdateTime = time.Now().UnixMilli() //int64类型的时间戳

	err = models.UpdateRecord(record) //调用模型层的添加课程方法

	if err != nil {
		c.JSON(200, result.Fail("更新失败")) //返回错误信息
		return
	}

	c.JSON(200, result.Success("更新成功")) //返回成功信息
}

// DeleteRecord 删除考勤记录
func DeleteRecord(c *gin.Context) {
	var record models.Records

	err := c.ShouldBind(&record) //获取参数

	if err != nil {
		c.JSON(200, result.Fail("获取参数失败")) //返回错误信息
		return
	}

	err = models.DeleteRecord(record) //调用模型层的添加课程方法

	if err != nil {
		c.JSON(200, result.Fail("删除失败")) //返回错误信息
		return
	}

	c.JSON(200, result.Success("删除成功")) //返回成功信息
}

// FindAllTheSameRecord 查询同一日期、年级、课程的记录
func FindAllTheSameRecord(c *gin.Context) {
	records, err := models.FindAllTheSameRecord()

	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(records))
}

// FindByDateAndGradeAndCourseIdAndStatus 根据日期、年级、课程ID、状态查询考勤记录
func FindByDateAndGradeAndCourseIdAndStatus(c *gin.Context) {
	// 获取参数
	date, _ := strconv.ParseInt(c.Query("date"), 10, 64) //获取参数,并将字符串转换为int64类型
	grade, _ := strconv.Atoi(c.Query("grade"))
	courseId := c.Query("courseId")
	status, _ := strconv.Atoi(c.Query("status"))

	records, err := models.FindByDateAndGradeAndCourseIdAndStatus(date, grade, courseId, status)

	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(records))
}

// FindNotCheckedStudents 查询指定课程中未签到的学生
func FindNotCheckedStudents(c *gin.Context) {
	// 获取参数
	date, _ := strconv.ParseInt(c.Query("date"), 10, 64) //获取参数,并将字符串转换为int64类型
	grade, _ := strconv.Atoi(c.Query("grade"))
	courseId := c.Query("courseId")

	students, err := models.FindNotCheckedStudents(date, grade, courseId)

	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(students))
}
