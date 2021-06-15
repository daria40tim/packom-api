package handler

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

func (h *Handler) getCpDocsById(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input packom.Org_docs

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = c.ShouldBind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBindUri(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	filename := strings.Split(input.Docs.Filename, "/")

	name := filename[len(filename)-1]

	err = c.SaveUploadedFile(input.Docs, "cps/"+strconv.Itoa(id)+"/"+name)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.CP.AddCPDoc(name, O_Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"o_id": O_Id,
	})
}

func (h *Handler) getCpDoc(c *gin.Context) {
	/*O_Id, err := getOId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}*/
	name := c.Param("name")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	target := "cps/" + strconv.Itoa(id) + "/" + name

	if strings.HasPrefix(filepath.Clean(target), "assets/") {
		c.String(403, "Look like you attacking me")
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+name)
	c.Header("Content-Type", "application/octet-stream")
	c.FileAttachment(target, name)
}
