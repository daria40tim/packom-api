package handler

import (
	"net/http"
	"strconv"

	"github.com/daria40tim/packom"
	"github.com/gin-gonic/gin"
)

type getAllOrgsResponse struct {
	Data []packom.OrgAll `json:"data"`
}

func (h *Handler) getAllOrgs(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	techs, err := h.services.Org.GetAll(O_Id /*, input*/)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrgsResponse{
		Data: techs,
	})
}

func (h *Handler) getOrgById(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	org, err := h.services.Org.GetById(O_Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) updateOrgById(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	var input packom.OrgI
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Org.UpdateById(O_Id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"o_id": id,
	})
}

func (h *Handler) addOrgById(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	o_id, err := h.services.Org.AddById(O_Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, o_id)
}
