package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	orgCtx              = "O_Id"
)

func (h *Handler) orgIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	O_Id, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(orgCtx, O_Id)
}

func getOId(c *gin.Context) (int, error) {
	O_Id, ok := c.Get(orgCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "org id not found")
		return 0, errors.New("org id not found")
	}

	o_idInt, ok := O_Id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "org id invalid")
		return 0, errors.New("org id invalid")
	}

	return o_idInt, nil
}
