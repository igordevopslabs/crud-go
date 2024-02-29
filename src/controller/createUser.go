package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/crud-go/src/configuration/rest_err"
	"github.com/igordevopslabs/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	//variavel do tipo UserRequest para acessar os campos da struct de user_request.go
	var userRequest request.UserRequest

	//se der qualquer erro nos campos que estou recebendo atraves da req vou retornar um BadRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("there is something wrong with filed, error=%s", err.Error()))

		//aqui eu preciso dizer para o gin retornar em json qual o erro e status code.
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
