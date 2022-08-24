package profile

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/profiles")
	routes.POST("/", h.addProfile)
	routes.GET("/", h.getProfiles)
	routes.GET("/:id", h.getProfilebyId)
	routes.PUT("/:id", h.updateProfile)
	routes.DELETE("/:id", h.deleteProfile)
}
