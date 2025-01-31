package controller

import (
	"goauth/src/main/consts"
	request2 "goauth/src/main/dto/request"
	"goauth/src/main/dto/response"
	"goauth/src/main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(
	authService service.IAuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// SignIn godoc
// @Summary Sign In API
// @Description sign in web app
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Data "Successful"
// @Failure 400 {object} response.DataError "Client Error"
// @Failure 500 {object} response.DataError "Server Error"
// @Router /signin [post]
// @Security AccessTokenAuth
func (a *AuthController) SignIn(c *gin.Context) {
	var signInDTO request2.SignInDTO
	if err := c.ShouldBindJSON(&signInDTO); err != nil {
		response.BadRequest(c, consts.InvalidRequest)
		return
	}
	a.authService.SignIn(signInDTO)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// SignUp godoc
// @Summary Sign Up API
// @Description sign up web app
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Data "Successful"
// @Failure 400 {object} response.DataError "Client Error"
// @Failure 500 {object} response.DataError "Server Error"
// @Router /signup [post]
// @Security AccessTokenAuth
func (a *AuthController) SignUp(c *gin.Context) {
	var signUpDTO request2.SignUpDTO
	if err := c.ShouldBindJSON(&signUpDTO); err != nil {
		response.BadRequest(c, consts.InvalidRequest)
		return
	}
	a.authService.SignUp(signUpDTO)
	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})
}
