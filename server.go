package main

import (
	"net/http"
	"fmt"
	"log"
	"awesomeProject/api/controllers"
	"awesomeProject/api/router"
)

func main() {
	fmt.Println(`Starting Development API Server.`)
	fmt.Println()

	defer fmt.Println(`Server exited.`)

	// Create the router.
	myRouter := router.New()

	// Create controller instances.
	//user := controllers.NewUserController()
	userController := controllers.NewUserController()

	//myRouter.Get(`/users`, func() controller)
	//
	//func Get(func() controller) {
	//	controller.New()
	//	controller.BaseController.Something = r
	//}

	// Establish the routes.
	//myRouter.Get(`/users`, user.Index)
	//myRouter.Post(`/users`, user.Create)
	//myRouter.Get(`/users/:id`, user.Show)
	//myRouter.Put(`/users/:id`, user.Update)
	//myRouter.Delete(`/users/:id`, user.Delete)
	myRouter.NewGet(`/users`, userController, `Index`)

	// Listen for traffic.
	log.Fatal(http.ListenAndServe(`:3001`, &myRouter.Router))
}
