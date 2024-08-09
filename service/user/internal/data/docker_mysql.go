package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	dockertest "github.com/ory/dockertest/v3"
)

func DockerMysql() (string, func()) {
	return innerDockerMysql()
}

func innerDockerMysql() (string, func()) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	pool.MaxWait = 2 * time.Second
	if err != nil {
		log.Fatalf("1 Could not connect to docker: %s", err)
	}

	// resource, err := pool.Run(img, version, []string{"MYSQL_ROOT_PASSWORD=123456", "MYSQL_ROOT_HOST=%"})
	// if err != nil {
	// 	log.Fatalf("Could not start resource: %s", err)
	// }
	// conStr := fmt.Sprintf("root:123456@(192.168.56.162:%s)/mysql?charset=utf8mb4&parseTime=True&loc=Local", resource.GetPort("3306/tcp"))
	conStr := fmt.Sprintf("root:123456@(192.168.56.162:%s)/shop?charset=utf8mb4&parseTime=True&loc=Local", "3306")
	if err = pool.Retry(func() error {
		var err error
		db, err := sql.Open("mysql", conStr)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("2 Could not connect to docker: %s", err)
	}
	return conStr, func() {
		// if err := pool.Purge(resource); err != nil {
		// 	log.Fatalf("Could not purge resource: %s", err)
		// }
	}
}
