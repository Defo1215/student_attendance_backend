package main

import (
	"student_attendance/database"
	"student_attendance/routes"
)

func main() {
	database.InitMySQL() // 初始化MySQL数据库连接

	r := routes.InitRouter() // 初始化路由

	_ = r.Run(":8088")
}
