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

/*
Package account include user、team、space's controller basic operation of logic
*/
package account

import (
	"encoding/json"
	"fmt"
	"time"

	base "ipaas/controllers"
	"ipaas/models/account"
	"ipaas/pkg/k8s/client"
	"ipaas/pkg/k8s/typed/core/v1"
	"ipaas/pkg/tools/storage/redis"
	"ipaas/pkg/tools/uuid"
	"ipaas/pkg/tools/validate"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UserController user controller
type UserController struct {
	base.BaseController
}

// Login login
// @Title Login server
// @Description Login server by username and password
// @Success 201		{object}	models.account.User
// @Param	body		body 	models.account.User		true	"body for user content"
// @router /login [post]
func (c *UserController) Login() {
	user, err := validate.ValidateUser(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	if user.Name == "" || user.Password == "" {
		c.Response400(fmt.Errorf("name and password mustn't null"))
		return
	}
	if user.GetByNameAndPassword() != nil {
		c.ResponseErrorUnameOrPassword()
		return
	}
	token := uuid.Token()
	if err := redis.GetClient().Set(user.Name, token, base.TOKEN_EXPIRE_TIME).Err(); err != nil {
		c.Response500(err)
		return
	}
	c.Data["json"] = map[string]interface{}{"token": token}
	c.ServeJSON()
}

// Logout Logout
// @Title Logout server
// @Description Login server by username and password
// @Success 201		{object}	models.account.User
// @Param	body		body 	models.account.User		true	"body for user content"
// @router /logout [delete]
func (c *UserController) Logout() {
	if err := redis.GetClient().Del(c.Ctx.Input.Header("Username")).Err(); err != nil {
		c.Response500(err)
		return
	}
	c.Data["json"] = map[string]interface{}{"logout": true}
	c.ServeJSON()
}

// Create create user
// @Title Create server
// @Description create a user
// @Success 201		{object}	models.account.User
// @Param	body		body 	models.account.User		true	"body for user content"
// @router / [post]
func (c *UserController) Create() {
	user, err := validate.ValidateUser(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	if user.Name == "" || user.Password == "" {
		c.Response400(fmt.Errorf("name and password mustn't null"))
		return
	}
	user.CreationTime = time.Now()
	user.LastLoginTime = time.Now()
	if err := user.Create(); err != nil {
		c.Response500(err)
		return
	}
	c.Response(201, user)

	createnamespace := func() {
		for clusterID, client := range client.GetClientsets() {
			glog.Info(clusterID)
			_, err := v1.Namespaces(client.Clientset).Create(TOK8sNamespace(user.Name))
			if err != nil {
				glog.Errorf("when add user,create k8s namespace [%v] in cluster [%v] err: %v", user.Name, clusterID, err)
			}
		}
	}
	go createnamespace()
}

// Delete delete user
// step:
// 1. create user
// 2. create user namespace in kubernentes cluster
// @Title Delete server
// @Description delete a user
// @Success 201		{object}	models.account.User
// @Param	body		body 	models.account.User		true	"body for user content"
// @router /:user [delete]
func (c *UserController) Delete() {
	uname := c.GetString(":user")
	if uname == "" {
		c.Response400(fmt.Errorf("the request param user mustn't null"))
		return
	}
	user := new(account.User)
	user.Name = uname
	if err := user.Delete(); err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, "ok")

	deletenamespace := func() {
		for clusterID, client := range client.GetClientsets() {
			if err := v1.Namespaces(client.Clientset).Delete(uname, &metav1.DeleteOptions{}); err != nil {
				glog.Errorf("when add user,create k8s namespace [%v] in cluster [%v] err: %v", user.Name, clusterID, err)
			}
		}
	}
	go deletenamespace()
}

// ResetPassword update user password
// @Title CreateUser server
// @Description create a user
// @Success 201		{object}	models.account.User
// @Param	body		body 	models.account.User		true	"body for user content"
// @router / [put]
func (c *UserController) ResetPassword() {
	var user account.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.Response400(err)
		return
	}
	if err := (&user).Update(); err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, "ok")
}

// List list all  user
// @Title list server
// @Description list all user
// @Success 201		{object}	models.account.User
// @Param	body		body 	models.account.User		true	"body for user content"
// @router / [get]
func (c *UserController) List() {
	users, err := new(account.User).ListAll()
	if err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, users)
}
