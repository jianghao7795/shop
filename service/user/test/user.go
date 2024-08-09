package main

import (
	"context"
	v1 "user/api/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userCient v1.UserClient
var conn *grpc.ClientConn

func main() {
	Init()
	TestCreateUser()
	conn.Close()
}

func Init() {
	var err error
	// 使用insecure.NewCredentials()替换已弃用的grpc.WithInsecure()
	//
	conn, err = grpc.NewClient("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc link err " + err.Error())
	}
	userCient = v1.NewUserClient(conn)
}

func TestCreateUser() {
	resp, err := userCient.CreateUser(context.Background(), &v1.CreateUserInfo{
		Mobile:   "13800138800",
		Password: "123123123123123",
		NickName: "test",
	})

	if err != nil {
		panic("grpc 用户创建失败" + err.Error())
	}

	println(resp.Id)
}
