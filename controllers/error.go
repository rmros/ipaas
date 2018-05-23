package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type Error struct {
	beego.Controller
}

func (e *Error) ResponseError(code int, err error) {
	e.Ctx.ResponseWriter.WriteHeader(code)
	e.Data["json"] = map[string]interface{}{"code": code, "errorMessage": err.Error()}
	e.ServeJSON()
}

func (e *Error) Response400(err error) {
	e.ResponseError(http.StatusBadRequest, err)
}

func (e *Error) Response401(err error) {
	e.ResponseError(http.StatusUnauthorized, err)
}

func (e *Error) Response500(err error) {
	e.ResponseError(http.StatusInternalServerError, err)
}

func (e Error) ResponseErrorUnameOrPassword() {
	e.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	e.Ctx.ResponseWriter.Write([]byte("username or password is not correct"))
}
