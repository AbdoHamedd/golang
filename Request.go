package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Request struct {
	context    *gin.Context
	DB         *gorm.DB
	connection *sql.DB
}

// handel request closure
func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.context = c
		request.DB, request.connection = connectionToDataBase()
		return request
	}
}

// Handel Response
func (req Request) Ok(body interface{}) {
	req.Response(200, body)
}
func (req Request) Response(code int, body interface{}) {
	req.connection.Close()
	req.context.JSON(code, body)
}

// init new request

func newRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	return req

}

// connect to data base
func connectionToDataBase() (*gorm.DB, *sql.DB) {

	dsn := "sqlserver://LAPTOP-KFK94AE1\\MSSQLSERVER01:LoremIpsum86@localhost:9930?database=golang"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database: %v", err.Error())
	}

	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error on get database connection: %v", err.Error())

	}
	return db, connection

}

// close database connection
func (req Request) closeDatabaseConnection() {
	req.connection.Close()
}
