package api

import (
	"github.com/gin-gonic/gin"
	"island/serializer"
	"island/service"
	"island/utils"
)

// @Summary 用户注册
// @Tags 用户
// @Accept  json
// @Produce json
// @Param register body service.UserRegisterService true "注册"
// @Success 200 {object} serializer.UserResponse
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var s service.UserRegisterService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	if user, err := s.Register(); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildUserResponse(user)
		c.JSON(200, res)
	}
}

// @Summary 用户登录
// @Tags 用户
// @Accept  json
// @Produce json
// @Param login body service.UserLoginService true "登录"
// @Success 200 {object} serializer.Response
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var s service.UserLoginService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	if token, err := s.Login(); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, Response(token))
	}
}

// @Summary 当前用户详情
// @Tags 用户
// @Accept  json
// @Produce json
// @Success 200 {object} serializer.UserResponse
// @Router /api/v1/user [get]
// @Security ApiKeyAuth
func UserInfo(c *gin.Context) {
	var s service.UserDetailService
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if user, err := s.Get(claims.Id); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildUserResponse(user)
		c.JSON(200, res)
	}
}

// @Summary 指定用户详情
// @Tags 用户
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} serializer.UserResponse
// @Router /api/v1/user/{id} [get]
func UserDetail(c *gin.Context) {
	var s service.UserDetailService
	if user, err := s.Get(c.Param("id")); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildUserResponse(user)
		c.JSON(200, res)
	}
}

// @Summary 用户列表
// @Tags 用户
// @Accept  json
// @Produce json
// @Param limit query int false "每页数量"
// @Param page query int false "当前页数"
// @Success 200 {object} serializer.UserListResponse
// @Router /api/v1/users [get]
func UserList(c *gin.Context) {
	var s service.UserListService
	if users, count, err := s.List(c); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildUserListResponse(users, count)
		c.JSON(200, res)
	}
}

// @Summary 更新用户
// @Tags 用户
// @Accept  json
// @Produce json
// @Param update body service.UserUpdateService true "更新"
// @Success 200 {object} serializer.UserResponse
// @Router /api/v1/user [put]
// @Security ApiKeyAuth
func UserUpdate(c *gin.Context) {
	var s service.UserUpdateService
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if user, err := s.Update(claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildUserResponse(user)
		c.JSON(200, res)
	}
}

// @Summary 删除用户
// @Tags 用户
// @Accept  json
// @Produce json
// @Success 200 {object} serializer.Response
// @Router /api/v1/user [delete]
// @Security ApiKeyAuth
func UserDelete(c *gin.Context) {
	var s service.UserDeleteService
	claims := c.MustGet("claims").(*utils.CustomClaims)
	if err := s.Delete(claims.ID); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, Response(nil))
	}
}
