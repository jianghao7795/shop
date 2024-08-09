package data_test

import (
	"context"
	"testing"
	"user/internal/conf"
	"user/internal/data"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}

var cleaner func()      // 定义删除 mysql 容器的回调函数
var Db *data.Data       // 用于测试的 data
var ctx context.Context // 上下文

// initialize  AutoMigrate gorm 自动建表的方法
func initialize(db *gorm.DB) error {
	err := db.AutoMigrate(
		&data.User{},
	)
	return errors.WithStack(err)
}

// ginkgo 使用 BeforeEach 为您的 Specs 设置状态
var _ = BeforeSuite(func() {
	// 执行测试数据库操作之前，链接之前 docker 容器创建的 mysql
	//con, f := data.DockerMysql("mysql", "latest")
	con, f := data.DockerMysql() //修改后的
	cleaner = f                  // 测试完成，关闭容器的回调方法
	config := &conf.Data{Database: &conf.Data_Database{Driver: "mysql", Source: con}}
	db := data.NewDB(config)
	mySQLDb, _, err := data.NewData(config, nil, db, nil)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	Db = mySQLDb
	err = initialize(db)
	if err != nil {
		return
	}
	Expect(err).NotTo(HaveOccurred())
})

// 测试结束后 通过回调函数，关闭并删除 docker 创建的容器
var _ = AfterSuite(func() {
	cleaner()
})
