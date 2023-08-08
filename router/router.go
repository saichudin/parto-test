package router

import (
	"saichudin/parto-test/contact"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/get-all-contact", contact.GetAllContact)
	r.GET("/get-contact/:id", contact.GetDetailContact)
	r.POST("/add-contact", contact.AddContact)
	r.PUT("/update-contact/:id", contact.UpdateContact)
	r.DELETE("/delete-contact/:id", contact.DeleteContact)
}