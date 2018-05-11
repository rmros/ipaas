package controllers

import (
	"ipaas/models/account"

	"github.com/astaxie/beego"
)

// BaseController the basic controller of all controller
type BaseController struct {
	beego.Controller
}

// HuaweiCloud assert the cloud is huawei cloud or not
func (c *BaseController) HuaweiCloud() bool {
	// clusterID := c.Ctx.Input.Param(":cluster")
	// clusterModel := &infrastructure.Cluster{}
	// if _, err := infrastructure.Get(clusterID); err != nil {
	// 	glog.Errorf("get cluster by id %v err: %v", clusterID, err)
	// 	return false
	// }
	// if clusterModel.Type == 1 {
	// 	return true
	// }
	// return false
	return true
}

// Prepare runs after Init before request function execution. (Interceptor)
func (c *BaseController) Prepare() {}

// Finish runs after request function execution.
func (c *BaseController) Finish() {}

func checkToken(username, token string) (*account.User, error) {
	return nil, nil
}
