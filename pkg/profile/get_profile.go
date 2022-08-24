package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-neves/golang-restApi-with-docker/pkg/common/models"
)

func (h handler) getProfilebyId(c *gin.Context) {
	id := c.Param("id")

	var profile models.Profile

	if result := h.DB.First(&profile, id); result.Error != nil {
		c.AbortWithError(http.StatusNotAcceptable, result.Error)
		return
	}

	c.JSON(http.StatusOK, profile)

}
