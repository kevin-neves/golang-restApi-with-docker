package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-neves/go-databases/pkg/common/models"
)

type AddProfileRequestBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Cpf  string `json:"cpf"`
}

func (h handler) addProfile(c *gin.Context) {
	var profileRequest AddProfileRequestBody

	if err := c.BindJSON(&profileRequest); err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
		return
	}

	var profile models.Profile

	profile.Name = profileRequest.Name
	profile.Age = profileRequest.Age
	profile.Cpf = profileRequest.Cpf

	if result := h.DB.Create(&profile); result.Error != nil {
		c.AbortWithError(http.StatusNotAcceptable, result.Error)
		return
	}

	c.JSON(http.StatusAccepted, profile)
}
