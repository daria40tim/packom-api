package handler

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

type getSpecsResponse struct {
	Data packom.Specs `json:"data"`
}

func (h *Handler) getSelectSpecs(c *gin.Context) {

	data, err := h.services.Org.SelectAllSpecs()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getSpecsResponse{
		Data: data,
	})

}

func (h *Handler) getDocsById(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input packom.Org_docs

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

	err = c.SaveUploadedFile(input.Docs, "assets/"+strconv.Itoa(O_Id)+"/"+name)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Org.AddDoc(name, O_Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"o_id": O_Id,
	})
}

func (h *Handler) getDoc(c *gin.Context) {
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

	target := "assets/" + strconv.Itoa(id) + "/" + name

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

func (h *Handler) deleteTrusted(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Org.DeleteTrustedOrg(O_Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"o_id": O_Id,
	})
}

type getOrgFilterResponse struct {
	Data packom.OrgFilterData `json:"data"`
}

func (h *Handler) getFilterData(c *gin.Context) {

	filters, err := h.services.Org.GetFilterData()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrgFilterResponse{
		Data: filters,
	})
}

func (h *Handler) getFilteredOrgs(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	var input packom.OrgFilterParams
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	techs, err := h.services.Org.GetAllFiltered(O_Id, input.Names, input.Groups, input.Specs, input.Countries)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrgsResponse{
		Data: techs,
	})
}
