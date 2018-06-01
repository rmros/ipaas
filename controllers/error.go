/*
Copyright 2018 huangjia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	e.Ctx.ResponseWriter.Write([]byte(err.Error()))
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
