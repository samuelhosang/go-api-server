package controllers

import (
	"net/http"
	"awesomeProject/api/models"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

//type UserController struct {
//	BaseController
//}

func NewUserController () UserController {
	return UserController{}
}

func (c UserController) New () Controller {
	return Controller(UserController{})
}

func (c *UserController) Index () Response {
	var (
		users models.Users
		err   error
	)


	err = users.Get()
	if err != nil {
		return c.HandleError(
			err,
			http.StatusInternalServerError,
		)
	}

	fmt.Println(`users`, users)

	resp := c.renderJSON(
		users,
		http.StatusOK,
	)

	fmt.Println("Controller Resp", resp)
	spew.Dump(resp)

	return resp
}



func (c *UserController) Create () Response {
	var (
		user models.User
		err error
	)

	//err = c.decodeJSON(*r, &user)
	if err != nil {
		return c.HandleError(
			err,
			http.StatusBadRequest,
		)
	}

	err = user.Save()
	if err != nil {
		return c.HandleError(
			err,
			http.StatusInternalServerError,
		)
	}

	return c.renderJSON(
		user,
		http.StatusCreated,
	)
}

func (c *UserController) Show () Response {
	var (
		user models.User
		userID string
		err error
	)

	//userID = p.ByName(`id`)

	err = user.Get(userID)
	if err != nil {
		return c.HandleError(
			err,
			http.StatusNotFound,
		)
	}

	return c.renderJSON(
		user,
		http.StatusOK,
	)
}

func (c *UserController) Update () Response {
	var (
		userID string
		user models.User
		err error
	)

	//userID = p.ByName(`id`)

	err = user.Get(userID)
	if err != nil {
		return c.HandleError(
			err,
			http.StatusNotFound,
		)
	}

	//err = c.decodeJSON(*r, &user)
	if err != nil {
		return c.HandleError(
			err,
			http.StatusBadRequest,
		)
	}

	err = user.Save()
	if err != nil {
		return c.HandleError(
			err,
			http.StatusInternalServerError,
		)
	}

	return c.renderJSON(
		user,
		http.StatusAccepted,
	)
}

func (c *UserController) Delete () Response {
	var (
		userID string
		user models.User
		err error
	)

	//userID = p.ByName(`id`)

	err = user.Get(userID)
	if err != nil {
		return c.HandleError(
			err,
			http.StatusNotFound,
		)
	}

	err = user.Delete()
	if err != nil {
		return c.HandleError(
			err,
			http.StatusInternalServerError,
		)
	}

	return c.renderJSON(
		user.ID,
		http.StatusOK,
	)
}