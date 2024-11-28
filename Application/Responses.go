package Application

import "github.com/gin-gonic/gin"

func (req Request) Ok(body interface{}) {
	req.Response(200, body)
}

func (req Request) Created(body interface{}) {
	req.Response(201, body)
}

func (req Request) NotAuth() {
	req.Response(401, gin.H{
		"message": "Yoy are Not Authorized",
	})
}
func (req Request) BadRequest(error interface{}) {
	req.Response(422, error)
}
func (req Request) UserNoutFound() {
	req.Response(404, gin.H{
		"message": "We Not Found this user in our system",
	})
}
