package handler

import (
	"net/http"

	"github.com/daria40tim/packom"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTech(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	var input packom.Tech
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Tech.Create(O_Id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"tz_id": id,
	})

}

func (h *Handler) getAllTechs(c *gin.Context) {

}

func (h *Handler) getTechById(c *gin.Context) {

}

func (h *Handler) updateTechById(c *gin.Context) {

}
