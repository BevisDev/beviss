package controller

import (
	"github.com/gin-gonic/gin"
	"goauth/src/main/consts"
	"goauth/src/main/dto/request"
	"goauth/src/main/dto/response"
	"goauth/src/main/service"
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
