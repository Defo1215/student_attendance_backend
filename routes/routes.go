package routes

import (
	"github.com/gin-gonic/gin"
	"student_attendance/app/controllers"
	"student_attendance/app/middleware/cors"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/public", "./public") // 静态文件服务

	r.Use(cors.Cors()) //启用跨域中间件

	//路由组件
	root := r.Group("/api")
	{
		//学生路由组
		student := root.Group("/student")
		{
			student.GET("/hello", func(c *gin.Context) {
				c.JSON(200, "这是一个测试接口")
			})
			student.POST("/save", controllers.AddStudent)
			student.GET("/findAll", controllers.FindAllStudent)
			student.GET("/findByGradeAndStudentId", controllers.FindStudentByGradeAndStudentId)
			student.GET("/findByGradeAndClassId", controllers.FindStudentByGradeAndClassId)
		}
		//年级路由组
		grade := root.Group("/grade")
		{
			grade.POST("/save", controllers.AddGrade)
			grade.GET("/findAll", controllers.FindAllGrade)
		}
		//课程路由组
		course := root.Group("/course")
		{
			course.POST("/save", controllers.AddCourse)
			course.GET("/findAll", controllers.FindAllCourse)
			course.GET("/findByGrade", controllers.FindCourseByGrade)
		}
		//记录路由组
		record := root.Group("/record")
		{
			record.POST("/save", controllers.AddRecord)
			record.POST("/update", controllers.UpdateRecord)
			record.POST("/delete", controllers.DeleteRecord)
			record.GET("/findAllTheSame", controllers.FindAllTheSameRecord)
			record.GET("/findByDateAndGradeAndCourseIdAndStatus", controllers.FindByDateAndGradeAndCourseIdAndStatus)
			record.GET("/findNotCheckedStudents", controllers.FindNotCheckedStudents)

		}
	}

	return r
}
