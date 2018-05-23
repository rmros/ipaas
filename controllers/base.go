/*
Copyright [yyyy] [name of copyright owner]

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
	"fmt"
	"strings"

	"ipaas/pkg/tools/log"
	"ipaas/pkg/tools/storage/redis"
)

// BaseController the basic controller of all controller
type BaseController struct {
	Error
	Namespace string
}

// Prepare runs after Init before request function execution. (Interceptor)
func (c *BaseController) Prepare() {
	url, method := c.Ctx.Request.URL.String(), c.Ctx.Request.Method
	log.Info(url, method)
	if !isLogin(url, method) {
		uname, _, token := c.Ctx.Input.Header("Username"), c.Ctx.Input.Header("Teamspace"), c.Ctx.Input.Header("Authorization")
		if token == "" || !strings.HasPrefix(strings.ToLower(token), "token") {
			c.Response401(fmt.Errorf("invalid Authorization header"))
			return
		}

		if expiredToken(uname, token[len("token")+1:]) {
			c.Response401(fmt.Errorf("invalid Authorization header, token is expired"))
			return
		}
	}

}

// Finish runs after request function execution.
func (c *BaseController) Finish() {}

func (c *BaseController) Response(code int, data interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Data["json"] = map[string]interface{}{"code": code, "data": data}
	c.ServeJSON()
}

func expiredToken(username, token string) bool {
	return redis.GetClient().Get(username).Val() == token
}

func isLogin(url, method string) bool {
	return url == "/api/v1/account/login" && method == "POST"
}
