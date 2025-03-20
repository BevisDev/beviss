package controller

import (
	"github.com/gin-gonic/gin"
	"gobaucua/src/main/consts"
	"gobaucua/src/main/dto/request"
	"gobaucua/src/main/dto/response"
	"gobaucua/src/main/service"
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
