package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-neves/go-databases/pkg/common/models"
)

type UpdateProfileRequestBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Cpf  string `json:"cpf"`
}

func (h handler) updateProfile(c *gin.Context) {
	id := c.Param("id")
	var updateProfileRequest UpdateProfileRequestBody
	var profileToUpdate models.Profile

	if err := c.BindJSON(&updateProfileRequest); err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
		return
	}

	if result := h.DB.First(&profileToUpdate, id); result.Error != nil {
		c.AbortWithError(http.StatusNotAcceptable, result.Error)
		return
	}

	profileToUpdate.Name = updateProfileRequest.Name
	profileToUpdate.Age = updateProfileRequest.Age
	profileToUpdate.Cpf = updateProfileRequest.Cpf

	h.DB.Save(profileToUpdate)
	c.JSON(http.StatusOK, profileToUpdate)

}
