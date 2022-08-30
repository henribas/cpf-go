package cpf

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidarEndpoint(ctx *gin.Context) {

	var cpf CPFDidatico
	if err := ctx.ShouldBind(&cpf); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": "request inválida",
		})
	}
	valido := cpf.Valido()
	ctx.JSON(http.StatusOK, gin.H{
		"CpfValido": valido,
	})

}
