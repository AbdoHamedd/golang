package main

import "github.com/gin-gonic/gin"

type Request struct {
	context *gin.Context
}

func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.context = c
		return request
	}
}

func newRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	return req

}
