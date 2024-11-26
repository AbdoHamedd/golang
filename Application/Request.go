package Application

import (
	"1/Models"
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
	User       Models.User
	IsAuth     bool
}

func (req Request) Auth() Request {
	req.IsAuth = false
	authHeader := req.context.GetHeader("Authorization")
	if authHeader != "" {
		req.DB.Where("token = ?", authHeader).First(&req.User) //token is not empty find the first token in the table that token is equal to the authHeader
		if req.User.ID != 0 {
			req.IsAuth = true
		}
	}
	return req
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
