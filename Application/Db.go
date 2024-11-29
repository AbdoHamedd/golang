package Application

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func makeConnections() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database: %v", err.Error())
	}
	return db
}

func returnConnection(db *gorm.DB) *sql.DB {
	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error on get database connection: %v", err.Error())

	}
	return connection
}

// connect to database by the share interface that we do on request directory
func connectionToDataBase(share ShareResorces) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.DB = makeConnections()
		app.connection = returnConnection(app.DB)
	case *Request:
		req := share.(*Request)
		req.DB = makeConnections()
		req.connection = returnConnection(req.DB)

	}
}

/*
	closeDatabaseConnection

=> used one after the migrate
&
befor sending the Response
*/
func CloseDatabaseConnection(share ShareResorces) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.connection.Close()
	case *Request:
		req := share.(*Request)
		req.connection.Close()
	}

}
