package api

import (
	"finance-tracker/internal/auth/biz"
	"finance-tracker/internal/auth/schema"
	"finance-tracker/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserBIZ biz.User
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Query user list
// @Param username query string false "Username for login"
// @Param name query string false "Name of user"
// @Param status query string false "Status of user (activated, freezed)"
// @Success 200 {object} util.ResponseResult{data=[]schema.User}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users [get]
func (a *User) Query(c *gin.Context) {
	result, err := a.UserBIZ.Query()
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Get user record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.User}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id} [get]
func (a *User) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := a.UserBIZ.Get(id)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Create user record
// @Param body body schema.UserForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.User}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users [post]
func (a *User) Create(c *gin.Context) {
	item := new(schema.UserForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.UserBIZ.Create(item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Update user record by ID
// @Param id path string true "unique id"
// @Param body body schema.UserForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id} [put]
func (a *User) Update(c *gin.Context) {
	item := new(schema.UserForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err := a.UserBIZ.Update(id, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Delete user record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id} [delete]
func (a *User) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := a.UserBIZ.Delete(id)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
