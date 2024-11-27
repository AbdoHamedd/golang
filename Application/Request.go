package Application

import (
	"database/sql"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"project1/Models"
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
	Lang       string
}

// handel request closure          (1)
func req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.context = c
		connectionToDataBase(&request)
		setLang(&request)
		return &request
	}
}

func setLang(request *Request) {
	lang := gotrans.DetectLanguage(request.context.GetHeader("Accept-Language"))
	gotrans.SetDefaultLocale(lang)
	request.Lang = lang
}

// init new request for a new Api Like in The Users
func NewRequest(c *gin.Context) *Request {
	request := req()
	return request(c)

}

// New
func NewREQUESTwithAUTN(c *gin.Context) Request {
	return NewRequest(c).Auth()
}

func AuthRequest(c *gin.Context) (*Request, bool) {
	request := NewREQUESTwithAUTN(c)
	if !request.IsAuth {
		request.NotAuth()
		return &request, false
	}
	return &request, true
}

func (req Request) Response(code int, body interface{}) {
	CloseDatabaseConnection(&req)
	req.context.JSON(code, body)
}

func (req *Request) Share() {

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
