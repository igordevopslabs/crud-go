package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/crud-go/src/configuration/logger"
	"github.com/igordevopslabs/crud-go/src/configuration/validation"
	"github.com/igordevopslabs/crud-go/src/controller/model/request"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	//informar via logs a jornada que esta sendo feita
	logger.Info("Init create user on controller", zap.String("journey", "createUser"))
	//variavel do tipo UserRequest para acessar os campos da struct de user_request.go
	var userRequest request.UserRequest

	//se der qualquer erro nos campos que estou recebendo atraves da req vou retornar um BadRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		//aqui eu preciso dizer para o gin retornar em json qual o erro e status code.
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	logger.Info("User created", zap.String("journey", "createUser"))

}
