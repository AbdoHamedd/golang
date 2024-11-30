package Application

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func makeConnections() *gorm.DB {
	user := os.Getenv("DATABASE_USER_NAME")
	//password := os.Getenv("DATABASE_USER_PASSWORD")
	host := os.Getenv("DATABASE_USER_HOST")
	port := os.Getenv("DATABASE_USER_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	// Construct the DSN (Data Source Name)
	dsn := user + ":" + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
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
