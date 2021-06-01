package handler

import (
	"net/http"
	"strconv"

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

type getAllTechsResponse struct {
	Data []packom.TechAll `json:"data"`
	CPS  []packom.CP_srv  `json:"cps"`
}

func (h *Handler) getAllTechs(c *gin.Context) {

	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	techs, cps, err := h.services.Tech.GetAll(O_Id /*, input*/)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTechsResponse{
		Data: techs,
		CPS:  cps,
	})

}

func (h *Handler) getTechById(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	tech, costs, calendars, err := h.services.Tech.GetById(O_Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tech.Calendars = calendars
	tech.Costs = costs

	c.JSON(http.StatusOK, tech)
}

func (h *Handler) updateTechById(c *gin.Context) {

}
