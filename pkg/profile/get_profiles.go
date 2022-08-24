package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-neves/go-databases/pkg/common/models"
)

func (h handler) getProfiles(c *gin.Context) {
	var profiles []models.Profile

	if result := h.DB.Find(&profiles); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, profiles)
}
