package Application

import (
	"github.com/bykovme/gotrans"
)

func (req Request) Ok(body interface{}) {
	req.Response(200, buildResponse(body, gotrans.T("Ok"), 200, nil))
}

func (req Request) Created(body interface{}) {
	req.Response(201, buildResponse(body, gotrans.T("Created"), 201, nil))
}

func (req Request) NotAuth() {
	req.Response(401, buildResponse(nil, gotrans.T("Unauthorized"), 401, nil))
}

func (req Request) BadRequest(error interface{}) {
	req.Response(422, buildResponse(nil, gotrans.T("Somthing wrong in your request"), 422, nil))
}

func (req Request) UserNoutFound() {
	req.Response(404, buildResponse(nil, gotrans.T("We Not Found this user in our system"), 404, nil))
}

func buildResponse(payload interface{}, message string, code int, err interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = code
	response["message"] = message
	response["payload"] = payload
	response["error"] = err
	return response
}
