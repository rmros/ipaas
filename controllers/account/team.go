/*
Copyright [huangjia] [name of copyright owner]

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
	"fmt"

	base "ipaas/controllers"
	"ipaas/models"
	"ipaas/pkg/tools/validate"

	"github.com/golang/glog"
)

// TeamController team controller
type TeamController struct {
	base.BaseController
}

// CreateTeam CreateTeam
// @Title CreateTeam server
// @Description create team
// @Success 200		{object}	models.Team
// @Param	body		body 	models.Team		true	"body for team content"
// @router / [post]
func (c *TeamController) CreateTeam() {
	team, err := validate.ValidateTeam(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	if team.Exsit() {
		c.Response400(fmt.Errorf("team [%v] was exsit in db", team.Name))
		return
	}
	if err = team.Create(); err != nil {
		glog.Errorf("create team [%v] err: %v", team.Name, err)
		c.Response500(err)
		return
	}
	c.Response(200, "ok")
}

// DeleteTeam DeleteTeam
// @Title DeleteTeam server
// @Description delete team
// @Success 200
// @router /:team [delete]
func (c *TeamController) DeleteTeam() {
	teamName := c.GetString(":team")
	if teamName == "" {
		c.Response400(fmt.Errorf("team name mustn't be null"))
		return
	}
	team := &models.Team{Name: teamName}
	if err := team.Delete(); err != nil {
		c.Response500(err)
	}
	c.Response(200, "ok")
}

// ListTeam ListTeam
// @Title ListTeam server
// @Description list team
// @Success 200		{object}	[]models.Team
// @router / [get]
func (c *TeamController) ListTeam() {
	teams, err := new(models.Team).ListAll()
	if err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, teams)
}

// AddUsers AddUsers
// @Title AddUsers server
// @Description add users to Team
// @Success 200		{object}	models.User
// @router /:team/users [post]
func (c *TeamController) AddUsers() {
	users, err := validate.ValidateTeamAddUsers(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	team := &models.Team{Name: c.GetString(":team"), Users: users}
	if err = team.Update(); err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, "ok")
}

// GetUsers GetUsers
// @Title GetUsers server
// @Description get users from Team
// @Success 200		{object}	[]models.User
// @router /:team/users [get]
func (c *TeamController) GetUsers() {
	id, _ := c.GetUint32(":team")
	team := &models.Team{}
	team.ID = uint(id)
	users, err := team.GetTeamUsers()
	if err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, users)
}

// AddSpace AddSpace
// @Title AddSpace server
// @Description add namespace to team
// @Param	body		body 	models.Space	true	"body for user content"
// @Success 200		{object}	models.Space
// @router /:team/spaces [post]
func (c *TeamController) AddSpace() {
	space, err := validate.ValidateSpace(c.Ctx.Request)
	if err != nil {
		c.Response400(err)
		return
	}
	if err = createNamespace(c.GetString(":cluster"), space.TOK8sNamespace()); err != nil {
		c.Response500(err)
		return
	}
	c.Response(200, "ok")
}
