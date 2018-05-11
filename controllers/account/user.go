/*
Package account include user、team、space's controller basic operation of logic
*/
package account

import (
	base "ipaas/controllers"
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

}

// Logout Logout
// @Title Logout server
// @Description Login server by username and password
// @Success 201		{object}	models.account.User
// @Param	body		body 	models.account.User		true	"body for user content"
// @router /login [post]
func (c *UserController) Logout() {

}
