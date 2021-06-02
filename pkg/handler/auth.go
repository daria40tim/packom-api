package handler

import (
	"net/http"

	"github.com/daria40tim/packom"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input packom.Org

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	o_id, err := h.services.Authorization.CreateOrg(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"o_id": o_id,
	})

}

type signInInput struct {
	Login string `json:"email"`
	Pwd   string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err, org := h.services.Authorization.GenerateToken(input.Login, input.Pwd)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":    token,
		"o_id":     org.O_id,
		"group_id": org.Group,
		"name":     org.Name,
	})

}
