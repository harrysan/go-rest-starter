package api

import (
	"finance-tracker/internal/mods/auth/biz"
	"finance-tracker/internal/mods/auth/schema"
	"finance-tracker/pkg/util"

	"github.com/gin-gonic/gin"
)

type Login struct {
	LoginBIZ *biz.Login
}

func NewLoginApi(loginBIZ *biz.Login) *Login {
	return &Login{LoginBIZ: *&loginBIZ}
}

// @Tags LoginAPI
// @Summary Login system with username and password
// @Param body body schema.LoginForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.LoginToken}
// @Failure 400 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/login [post]
func (a *Login) Login(c *gin.Context) {
	item := new(schema.LoginForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	}

	data, err := a.LoginBIZ.Login(item.Trim())
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Logout system
// @Success 200 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/logout [post]
func (a *Login) Logout(c *gin.Context) {
	// ctx := c.Request.Context()
	// err := a.LoginBIZ.Logout(ctx)
	// if err != nil {
	// 	util.ResError(c, err)
	// 	return
	// }
	util.ResOK(c)
}
