package repository

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Database struct {
}

var logMode = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

func InitDB() {
	var err error
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	debug := os.Getenv("DB_DEBUG_MYSQL")
	mode := os.Getenv("LOG_MODE_MYSQL")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logMode[mode]),
	})

	if err != nil {
		fmt.Println("Error Connect database", err)
		panic(500)
	}

	fmt.Println("Connection Successfully")

	if debug == "true" {
		DB.Debug()
		return
	}
}

func (db *Database) GetHealthCheck(ctx *gin.Context) string {
	return "Alhamdulillah Sehat"
}
