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

	var input packom.CPIns
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input packom.CPIns
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cp_id, err := h.services.CP.UpdateById(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"cp_id": cp_id,
	})

}

func (h *Handler) cpDeleteCal(c *gin.Context) {
	var input int
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	/*id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}*/

	data, err := h.services.CP.DeleteCal(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) cpDeleteCst(c *gin.Context) {
	var input int
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	/*id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}*/

	data, err := h.services.CP.DeleteCst(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
