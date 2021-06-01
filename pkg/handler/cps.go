package handler

import (
	"net/http"
	"strconv"

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

type getAllCPsResponse struct {
	Data []packom.CPAll `json:"data"`
}

func (h *Handler) getAllCPs(c *gin.Context) {

	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	cps, err := h.services.CP.GetAll(O_Id /*, input*/)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCPsResponse{
		Data: cps,
	})
}

func (h *Handler) getCPById(c *gin.Context) {

	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	cp, err := h.services.CP.GetById(O_Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cp)
}

func (h *Handler) updateCPById(c *gin.Context) {

}
