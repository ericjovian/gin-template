package db

import (
	"fmt"
	"log"
	"os"

	"github.com/ericjovian/gin-template/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	c  = config.DBConfig
	db *gorm.DB
)

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() (err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
		c.Port,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: getLogger(),
	})
	return
}

func Get() *gorm.DB {
	return db
}
