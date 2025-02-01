package controller

import (
	"github.com/gin-gonic/gin"
	"goiam/src/main/consts"
	"goiam/src/main/dto/request"
	"goiam/src/main/dto/response"
	"goiam/src/main/service"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateRole(c *gin.Context) {
	var roleDTO request.RoleDTO
	if err := c.ShouldBindJSON(&roleDTO); err != nil {
		response.BadRequest(c, consts.InvalidRequest)
		return
	}
	uc.userService.CreateRole()
}

func (uc *UserController) AssignRole() {

}
