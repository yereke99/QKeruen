package controller

import (
	"fmt"
	"net/http"
	"qkeruen/dto"
	"qkeruen/help"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type authController struct {
	AuthService service.AuthService
	JWTService  service.JWTService
}

func NewAuthController(authservice service.AuthService, jwtService service.JWTService) authController {
	return authController{
		AuthService: authservice,
		JWTService:  jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RequestRegisterDTO
	var responseDTO dto.ResponseDTO
	if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}
	ok, err := c.AuthService.Check(registerDTO.PhoneNumber, registerDTO.Role)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "can not generate token"})
		return
	}
	if !ok {
		token_, err := c.AuthService.GiveTokenService(registerDTO.PhoneNumber, registerDTO.Role)

		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"message": "can not take token."})
			return
		}
		responseDTO.Token = token_
		responseDTO.IsAuthorized = true
		ctx.JSON(http.StatusAccepted, responseDTO)

	}

	code := help.GenerateRandomID(4)
	if smsServiceErr := c.AuthService.Create(registerDTO.PhoneNumber, code); smsServiceErr != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in sms service."})
		return
	}

	token, err := c.JWTService.GenerateToken(registerDTO.PhoneNumber, registerDTO.Role)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "can not generate token"})
		return
	}

	responseDTO.Token = token
	responseDTO.IsAuthorized = false
	ctx.JSON(201, responseDTO)

}

func (c *authController) ValidatorSMS(ctx *gin.Context) {
	var checkCode dto.CheckCodeRequest

	if err := ctx.ShouldBindJSON(&checkCode); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}
	code, _ := strconv.Atoi(checkCode.Code)
	ok, err := c.AuthService.ValidateSMS(checkCode.PhoneNumber, code)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}

	if !ok {
		ctx.JSON(http.StatusConflict, gin.H{"message": "wrong sms code."})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "Accepted."})
}

func (c *authController) ResendCode(ctx *gin.Context) {
	var registerDTO dto.RequestRegisterDTO
	var responseDTO dto.ResponseDTO
	if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}
	ok, err := c.AuthService.Check(registerDTO.PhoneNumber, registerDTO.Role)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "can not generate token"})
		return
	}
	if !ok {
		token_, err := c.AuthService.GiveTokenService(registerDTO.PhoneNumber, registerDTO.Role)

		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"message": "can not take token."})
			return
		}
		responseDTO.Token = token_
		responseDTO.IsAuthorized = true
		ctx.JSON(http.StatusAccepted, responseDTO)

	}

	code := help.GenerateRandomID(4)
	if smsServiceErr := c.AuthService.Create(registerDTO.PhoneNumber, code); smsServiceErr != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in sms service."})
		return
	}

	token, err := c.JWTService.GenerateToken(registerDTO.PhoneNumber, registerDTO.Role)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "can not generate token"})
		return
	}

	responseDTO.Token = token
	responseDTO.IsAuthorized = false
	ctx.JSON(201, responseDTO)

}

func (c *authController) CheckToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Empty token",
			},
		)
		// exit process
		return
	}

	which, _ := c.JWTService.Definition(token)
	fmt.Println(which)
	switch which {
	case "driver":
		data, err := c.AuthService.CheckTokenDriver(token)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest, gin.H{
					"error": "error in driver check token service.",
				},
			)
			// exit process
			return
		}
		ctx.JSON(200, data)

	case "user":
		datauser, err := c.AuthService.CheckTokenUser(token)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest, gin.H{
					"error": "error in user check token service.",
				},
			)
			// exit process
			return
		}
		ctx.JSON(200, datauser)

	}

}
