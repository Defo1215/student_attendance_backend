package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"student_attendance/app/models"
	"student_attendance/app/result"
)

// AddGrade 添加年级
func AddGrade(c *gin.Context) {
	var grade models.Grade

	err := c.ShouldBind(&grade) //获取参数

	if err != nil {
		c.JSON(200, result.Fail("获取参数失败")) //返回错误信息
		return
	}

	grade.Id = uuid.New().String() //生成uuid

	id, err := models.AddGrade(grade) //调用模型层的添加年级方法

	if err != nil {
		c.JSON(200, result.Fail("新增失败")) //返回错误信息
		return
	}

	c.JSON(200, result.Success(id)) //返回成功信息
}

// FindAllGrade 查询所有年级
func FindAllGrade(c *gin.Context) {
	grade, err := models.FindAllGrade()

	if err != nil {
		c.JSON(200, "查询失败")
		return
	}

	c.JSON(200, result.Success(grade))
}
