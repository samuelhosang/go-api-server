package controllers

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
)

type (
	BaseController struct {
		Response Response
		Request interface{}
		Params interface{}
	}

	UserController struct {
		BaseController
	}

	Response struct {
		Body interface{} `json:"body"`
		Status int `json:"status"`
		Error error `json:"error"`
	}

	Endpoint func () Response

	Controller interface{
		New() Controller
		SetParams ()
		SetRequest ()
	}
)

func (c *BaseController) New (c2 Controller) () {

}

func (c *BaseController) SetParams (params httprouter.Params) {
	c.Params = params
}

func (c *BaseController) SetRequest (request *http.Request) {
	c.Request = request
}

func (c *BaseController) HandleError (err error, status int) Response {
	c.Response.Error = err
	c.Response.Status = status
	return c.Response
}

func (c *BaseController) renderJSON (body interface{}, status int) Response {
	c.Response.Body = body
	c.Response.Status = status
	return c.Response
}

func (c *BaseController) decodeJSON (req http.Request, a interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &a)
	if err != nil {
		return err
	}

	return nil
}