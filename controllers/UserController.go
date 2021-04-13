package controllers

import (
	"boilerplate-golang/models"
	"boilerplate-golang/services"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (t *UserController) GetUsers() {
	users, err := services.GetUsers()
	if err != nil {
		getUserErrorResponse(t, http.StatusNotFound, err.Error())
	}

	getUserOkResponse(t, http.StatusOK, users)
}

func (t *UserController) GetUserByUsername() {
	name := t.GetString("name")
	user, err := services.GetUserByUsername(&name)
	if err != nil {
		getUserErrorResponse(t, http.StatusNotFound, err.Error())
	}

	getUserOkResponse(t, http.StatusOK, user)
}

func (t *UserController) GetUserById() {
	id, _ := t.GetInt("id")
	user, err := services.GetUserById(&id)
	if err != nil {
		getUserErrorResponse(t, http.StatusNotFound, err.Error())
	}

	getUserOkResponse(t, http.StatusOK, user)
}

func (t *UserController) UpdateUserById() {
	id, _ := t.GetInt("id")
	num, err := services.UpdateUserById(&id)
	if err != nil {
		getUserErrorResponse(t, http.StatusBadRequest, err.Error())
	}

	getUserOkResponse(t, http.StatusOK, models.OkResponse{
		Message: fmt.Sprintf("%d record/s have been updated", num),
		Status:  http.StatusOK,
	})
}

func (t *UserController) DeleteUserById() {
	id, _ := t.GetInt("id")
	num, err := services.DeleteUserById(&id)
	if err != nil {
		getUserErrorResponse(t, http.StatusBadRequest, err.Error())
	}

	getUserOkResponse(t, http.StatusOK, models.OkResponse{
		Message: fmt.Sprintf("%d record/s have been deleted", num),
		Status:  http.StatusOK,
	})
}

func getUserOkResponse(t *UserController, status int, data interface{}) {
	t.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	t.Ctx.ResponseWriter.WriteHeader(status)
	t.Data["json"] = data
	t.ServeJSON()
}

func getUserErrorResponse(t *UserController, status int, message string) {
	var errorResponse = models.ErrorResponse{}

	errorResponse.Status = status
	errorResponse.Message = message

	t.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	t.Ctx.ResponseWriter.WriteHeader(status)
	t.Data["json"] = errorResponse
	t.ServeJSON()
}
