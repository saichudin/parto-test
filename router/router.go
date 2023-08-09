package router

import (
	"saichudin/parto-test/auth"
	"saichudin/parto-test/contact"
	"saichudin/parto-test/middleware"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.POST("/login", auth.Login)

	r.GET("/get-all-contact", middleware.RequireAuth, contact.GetAllContact)
	r.GET("/get-contact/:id", middleware.RequireAuth, contact.GetDetailContact)
	r.POST("/add-contact", middleware.RequireAuth, contact.AddContact)
	r.PUT("/update-contact/:id", middleware.RequireAuth, contact.UpdateContact)
	r.DELETE("/delete-contact/:id", middleware.RequireAuth, contact.DeleteContact)
}
