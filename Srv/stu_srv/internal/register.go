package internal

import (
	"google.golang.org/grpc"
	"stu_srv/proto_stu/student"
	server2 "stu_srv/server"
)

func RegisterStudentServer(server grpc.ServiceRegistrar) {
	student.RegisterStudentServer(server, server2.ServerStudent{})
}
