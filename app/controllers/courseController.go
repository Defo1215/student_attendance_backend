package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"student_attendance/app/models"
	"student_attendance/app/result"
)

// AddCourse 新增课程
func AddCourse(c *gin.Context) {
	var course models.Course

	err := c.ShouldBind(&course) //获取参数

	if err != nil {
		c.JSON(200, result.Fail("获取参数失败")) //返回错误信息
		return
	}

	course.Id = uuid.New().String() //生成uuid

	id, err := models.AddCourse(course) //调用模型层的添加课程方法

	if err != nil {
		c.JSON(200, result.Fail("新增失败")) //返回错误信息
		return
	}

	c.JSON(200, result.Success(id)) //返回成功信息
}

// FindAllCourse 查询所有课程
func FindAllCourse(c *gin.Context) {
	courses, err := models.FindAllCourse()

	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(courses))
}

// FindCourseByGrade 根据年级查询课程
func FindCourseByGrade(c *gin.Context) {
	grade, _ := strconv.Atoi(c.Query("grade")) //获取参数,并将字符串转换为int类型

	courses, err := models.FindCourseByGrade(grade)

	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(courses))
}
