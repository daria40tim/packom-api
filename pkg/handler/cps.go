package handler

import (
	"net/http"

	"github.com/daria40tim/packom"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCP(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	var input packom.CP
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CP.Create(O_Id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"cp_id": id,
	})
}

func (h *Handler) getAllCPs(c *gin.Context) {

}

func (h *Handler) getCPById(c *gin.Context) {

}

func (h *Handler) updateCPById(c *gin.Context) {

}
