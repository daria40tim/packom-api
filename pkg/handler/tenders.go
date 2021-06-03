package handler

import (
	"net/http"
	"strconv"

	"github.com/daria40tim/packom"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTender(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	var input packom.Tender
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Tender.Create(O_Id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"tender_id": id,
	})

}

type getAllTendersResponse struct {
	Data []packom.TenderAll `json:"data"`
}

func (h *Handler) getAllTenders(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	techs, err := h.services.Tender.GetAll(O_Id /*, input*/)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTendersResponse{
		Data: techs,
	})

}

func (h *Handler) getTenderById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	tender, err := h.services.Tender.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tender)
}

type getFullCostsResponse struct {
	Data []packom.FullCost `json:"data"`
}

func (h *Handler) getMinandMax(c *gin.Context) {
	var input string
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tz_id, err := strconv.Atoi(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	res, err := h.services.Tender.GetFullCosts(tz_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getFullCostsResponse{
		Data: res,
	})

}
