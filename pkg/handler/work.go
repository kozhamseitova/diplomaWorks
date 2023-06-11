package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"net/http"
	"strconv"
)

func (h *Handler) createWork(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.WorkInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Work.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllWorksResponse struct {
	Data []model.Work `json:"data"`
}

type getWorksResponse struct {
	Data []model.WorkInstructor `json:"data"`
}

func (h *Handler) getAllWorks(c *gin.Context) {
	works, err := h.services.Work.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWorksResponse{
		Data: works,
	})
}

func (h *Handler) getWorkById(c *gin.Context) {

}

func (h *Handler) getAllWorksForAdmin(c *gin.Context) {
	works, err := h.services.Work.GetAllWorksForAdmin()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWorksResponse{
		Data: works,
	})
}

func (h *Handler) approveWork(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.WorkUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Work.UpdateWork(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getWorksByInstructorId(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("instructor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if userId != id {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	works, err := h.services.Work.GetWorksByInstructorId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorksResponse{
		Data: works,
	})
}

func (h *Handler) updateWork(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	instructorId, err := strconv.Atoi(c.Param("instructor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid instructor id param")
		return
	}

	if userId != instructorId {
		newErrorResponse(c, http.StatusBadRequest, "invalid instructor id param")
		return
	}

	var input model.WorkUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	workId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid work id param")
		return
	}

	if err := h.services.Work.UpdateWork(userId, workId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteWork(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	instructorId, err := strconv.Atoi(c.Param("instructor_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid instructor id param")
		return
	}

	if userId != instructorId {
		newErrorResponse(c, http.StatusBadRequest, "invalid instructor id param")
		return
	}

	workId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid work id param")
		return
	}

	if err := h.services.Work.DeleteWork(workId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
