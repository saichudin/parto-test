package contact

import (
	"net/http"
	"saichudin/parto-test/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllContact(c *gin.Context) {
	contacts := GetAllContactRepo()

	resp := helper.DefaultResponse{
		Message: "Success",
		Data:    contacts,
	}
	c.JSON(http.StatusOK, &resp)
}

func GetDetailContact(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	contact, err := GetDetailContactRepo(id)
	if err != nil {
		resp := helper.DefaultResponse{
			Message: err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, &resp)
		return
	}

	resp := helper.DefaultResponse{
		Message: "Success",
		Data:    contact,
	}
	c.JSON(http.StatusOK, &resp)
}

func AddContact(c *gin.Context) {
	var contact Contact
	errBind := c.ShouldBindJSON(&contact)
	if errBind != nil {
		resp := helper.DefaultResponse{
			Message: "Error validation, name and phone is required",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, &resp)
		return
	}

	err := AddContactRepo(&contact)
	if err != nil {
		resp := helper.DefaultResponse{
			Message: err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, &resp)
		return
	}

	resp := helper.DefaultResponse{
		Message: "Created",
		Data:    contact,
	}
	c.JSON(http.StatusCreated, &resp)
}

func UpdateContact(c *gin.Context) {
	var contact Contact
	errBind := c.ShouldBindJSON(&contact)
	if errBind != nil {
		resp := helper.DefaultResponse{
			Message: "Error validation, name and phone is required",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))

	err := UpdateContactRepo(id, &contact)
	if err != nil {
		resp := helper.DefaultResponse{
			Message: err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, &resp)
		return
	}

	resp := helper.DefaultResponse{
		Message: "Updated",
		Data:    contact,
	}
	c.JSON(http.StatusOK, &resp)
}

func DeleteContact(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := DeleteContactRepo(id)
	if err != nil {
		resp := helper.DefaultResponse{
			Message: err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, &resp)
		return
	}

	resp := helper.DefaultResponse{
		Message: "Deleted",
	}
	c.JSON(http.StatusOK, &resp)
}