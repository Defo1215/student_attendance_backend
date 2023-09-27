package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"student_attendance/app/models"
	"student_attendance/app/result"
)

// AddStudent 新增学生
func AddStudent(c *gin.Context) {
	var student models.Students

	err := c.ShouldBind(&student) //获取参数

	if err != nil {
		c.JSON(200, result.Fail("获取参数失败")) //返回错误信息
		return
	}

	student.Id = uuid.New().String() //生成uuid
	student.Status = 1               //默认状态为1

	id, err := models.AddStudent(student) //调用模型层的添加学生方法

	if err != nil {
		c.JSON(200, result.Fail("新增失败")) //返回错误信息
		return
	}

	c.JSON(200, result.Success(id)) //返回成功信息
}

// FindAllStudent 查询所有学生
func FindAllStudent(c *gin.Context) {
	students, err := models.FindAllStudent()

	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(students))
}

// FindStudentByGradeAndStudentId 根据年级和学号查询学生
func FindStudentByGradeAndStudentId(c *gin.Context) {
	grade, _ := strconv.Atoi(c.Query("grade")) //获取参数,并将字符串转换为int类型
	studentId, _ := strconv.Atoi(c.Query("studentId"))

	student, err := models.FindStudentByGradeAndStudentId(grade, studentId)

	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(student))
}

// FindStudentByGradeAndClassId 根据年级和班级查询学生
func FindStudentByGradeAndClassId(c *gin.Context) {
	grade, _ := strconv.Atoi(c.Query("grade")) //获取参数,并将字符串转换为int类型
	classId, _ := strconv.Atoi(c.Query("classId"))

	student, err := models.FindStudentByGradeAndClassId(grade, classId)

	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(student))
}
