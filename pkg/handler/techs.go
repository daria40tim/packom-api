package handler

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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
}

func (h *Handler) getAllTechs(c *gin.Context) {

	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	techs, err := h.services.Tech.GetAll(O_Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTechsResponse{
		Data: techs,
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

	var input packom.Tech
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	tz_id, err := h.services.Tech.UpdateById(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"tz_id": tz_id,
	})
}

type getSelectResponse struct {
	Data packom.Select `json:"data"`
}

func (h *Handler) getSelect(c *gin.Context) {

	data, err := h.services.Tech.SelectAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getSelectResponse{
		Data: data,
	})

}

type delCal struct {
	Tz_id   string `json:"tz_id"`
	Task    string `json:"task_name"`
	History string `json:"history"`
}

func (h *Handler) deleteCal(c *gin.Context) {
	var input delCal
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	a, err := strconv.Atoi(input.Tz_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.Tech.DeleteCal(a, input.Task, input.History)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

type delCst struct {
	Tz_id   string `json:"tz_id"`
	Task    string `json:"task"`
	History string `json:"history"`
}

func (h *Handler) deleteCst(c *gin.Context) {

	var input delCst
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	a, err := strconv.Atoi(input.Tz_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	data, err := h.services.Tech.DeleteCost(a, input.Task, input.History)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)

}

func (h *Handler) getTechDocsById(c *gin.Context) {
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

	err = c.SaveUploadedFile(input.Docs, "techs/"+strconv.Itoa(id)+"/"+name)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Tech.AddTechDoc(name, O_Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"o_id": O_Id,
	})
}

func (h *Handler) getTechDoc(c *gin.Context) {
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

	target := "techs/" + strconv.Itoa(id) + "/" + name

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

type getTechFilterResponse struct {
	Data packom.TechFilterData `json:"data"`
}

func (h *Handler) getTechFilterData(c *gin.Context) {
	filters, err := h.services.Tech.GetFilterData()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getTechFilterResponse{
		Data: filters,
	})
}

func (h *Handler) getFilteredTechs(c *gin.Context) {
	O_Id, err := getOId(c)
	if err != nil {
		return
	}

	var input packom.TechFilterParams
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	techs, err := h.services.Tech.GetAllTechsFiltered(O_Id, input.EDate, input.SDate, input.Clients, input.Projs, input.TZ_STS, input.CP_STS, input.Tender_STS)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTechsResponse{
		Data: techs,
	})
}
