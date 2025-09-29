package main

import (
	"Common/appconfig"
	"Common/initialize"
	"google.golang.org/grpc"
	"stu_srv/grpc_stu"
	"stu_srv/internal"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos(func() {
		initialize.MysqlInit()
		initialize.RedisInit()
	})
	initialize.MysqlInit()
	initialize.RedisInit()
	initialize.ZapInit()
	//global.DB.AutoMigrate(&mysql.Student{})
	grpc_stu.RegisterStuGrpc(func(server *grpc.Server) {
		internal.RegisterStudentServer(server)
	})
}
