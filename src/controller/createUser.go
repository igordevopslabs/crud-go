package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/crud-go/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Voce chamou a rota incorretamente")
	c.JSON(err.Code, err)
}
