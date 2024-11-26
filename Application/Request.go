package Application

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShareResorces interface {
	Share()
}

type Request struct {
	context    *gin.Context
	DB         *gorm.DB
	connection *sql.DB
}

func (req *Request) Share() {

}

// handel request closure
func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.context = c
		connectionToDataBase(&request)
		return request
	}
}

func (req Request) Response(code int, body interface{}) {
	CloseDatabaseConnection(&req)
	req.context.JSON(code, body)
}

// init new request
func NewRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	return req

}
