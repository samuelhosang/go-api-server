package controllers

//import (
//	"github.com/julienschmidt/httprouter"
//	"net/http"
//	"fmt"
//)
//
//type HelloController struct {
//	BaseController
//}
//
//func NewHelloController () *HelloController {
//	return &HelloController{}
//}
//
//func (c HelloController) Root (p httprouter.Params) (r Response) {
//	r.Status = http.StatusOK
//	r.Body = `Welcome!`
//	return r
//}
//
//func (c HelloController) Index (p httprouter.Params) (r Response) {
//	r.Status = http.StatusCreated
//	r.Body = fmt.Sprintf(`Hello, %s!`, p.ByName(`name`))
//	return r
//}