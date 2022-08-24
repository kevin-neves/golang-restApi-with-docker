package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-neves/golang-restApi-with-docker/pkg/common/models"
)

func (h handler) deleteProfile(c *gin.Context) {
	id := c.Param("id")
	var profile models.Profile

	if result := h.DB.First(&profile, id); result.Error != nil {
		c.AbortWithError(http.StatusNotAcceptable, result.Error)
		return
	}

	h.DB.Delete(&profile)

	c.Status(http.StatusOK)
}
