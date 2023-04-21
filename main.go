package main

import (
	"github.com/gin-gonic/gin"
	"go-clean-translation/controller/httpapi"
	"go-clean-translation/infras/googlesv"
	mysqlRepo "go-clean-translation/infras/mysql"
	translateServ "go-clean-translation/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func main() {
	db, err := connectDBWithRetry(5)

	if err != nil {
		log.Fatalln(err)
	}

	// Setup Dependencies
	repository := mysqlRepo.NewMySQLRepo(db)
	googleTranslate := googlesv.New()
	service := translateServ.NewService(repository, googleTranslate)
	controller := httpapi.NewAPIController(service)

	engine := gin.Default()

	v1 := engine.Group("/v1")
	controller.SetUpRoute(v1)

	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}

}

func connectDBWithRetry(times int) (*gorm.DB, error) {
	var e error

	for i := 1; i <= times; i++ {
		dsn := os.Getenv("MYSQL_DSN")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			return db, nil
		}

		e = err

		time.Sleep(time.Second * 2)
	}

	return nil, e
}
