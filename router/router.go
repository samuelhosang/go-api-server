package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"awesomeProject/api/controllers"
	"encoding/json"
	"time"
	"fmt"
	"strconv"
	"reflect"
	"github.com/davecgh/go-spew/spew"
)

type Router struct {
	Router httprouter.Router
}

func New () (r Router) {
	r.Router = *httprouter.New()
	return r
}

// Report the request to the console.
func Log (m string, u string, s int) {
	fmt.Println(fmt.Sprintf(
		`[app] %s %s %s [status=%s]`,
		time.Now().Format(time.RFC3339),
		m,
		u,
		strconv.Itoa(s),
	))
}

func wrapper (h func () controllers.Response) httprouter.Handle {

	return func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// Format the response as JSON.
		result := h()
		resp, _ := json.Marshal(result)

		// Set the status code.
		w.WriteHeader(result.Status)

		Log(r.Method, r.URL.String(), result.Status)

		// Send the response
		fmt.Fprint(w, string(resp))

	}

}

func (r *Router) Get (path string, endpoint controllers.Endpoint) () {
	r.Router.GET(path, wrapper(endpoint))
}

func (r *Router) Post (path string, endpoint controllers.Endpoint) () {
	r.Router.POST(path, wrapper(endpoint))
}

func (r *Router) Put (path string, endpoint controllers.Endpoint) () {
	r.Router.PUT(path, wrapper(endpoint))
}

func (r *Router) Delete (path string, endpoint controllers.Endpoint) () {
	r.Router.DELETE(path, wrapper(endpoint))
}

func (r *Router) NewGet (path string, con controllers.Controller, method string) {
	fmt.Println(`path`, path)
	r.Router.GET(path, RestDispatchMethod(con, method))
}

func RestDispatchMethod(sample interface{}, method string) httprouter.Handle {
	t := reflect.ValueOf(sample).Type()

	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		v := reflect.New(t)

		fmt.Println(`~~~~~~~~~~~~~~~~~~~~~~`)
		spew.Dump(v)
		fmt.Println(`~~~~~~~~~~~~~~~~~~~~~~`)

		CallControllerMethod(method, v, w, req, p)
	}
}

func CallControllerMethod(methodName string, cont reflect.Value, w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	method := cont.MethodByName(methodName)
	controller := cont.Interface()

	celem := cont.Elem()
	if method.IsValid() {
		//vresp := method.Call([]reflect.Value{0: vw, 1: vreq, 2: vp})
		vresp := method.Call([]reflect.Value{})
		fmt.Println(`---------------`)
		fmt.Println(`celem`)
		spew.Dump(celem)
		fmt.Println(`---------------`)
		fmt.Println(`controller`)
		spew.Dump(controller)
		fmt.Println(`---------------`)
		fmt.Println(`vresp`)
		spew.Dump(vresp)


	} else {
		fmt.Println(`Unable to find method`, methodName)
	}
}
