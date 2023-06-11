package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"net/http"
	"strconv"
)

type getAllRequestsResponse struct {
	Data []model.Request `json:"data"`
}

func (h *Handler) createRequest(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid student id param")
		return
	}
	if userId != studentId {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.RequestInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Request.CreateRequest(model.RequestInput{
		StudentId:   studentId,
		WorkId:      input.WorkId,
		Description: input.Description,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllRequestsByStudentId(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid student id param")
		return
	}

	if userId != id {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	requests, err := h.services.Request.GetRequestsByStudentId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllRequestsResponse{
		Data: requests,
	})
}

func (h *Handler) closeRequest(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid student id param")
		return
	}

	if userId != id {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	requestsId, err := strconv.Atoi(c.Param("requests_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request id param")
		return
	}

	input := model.RequestStatus{
		Id:       requestsId,
		StatusId: 3,
	}

	if err := h.services.Request.ChangeStatus(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) changeRequestStatus(c *gin.Context) {
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

	requestsId, err := strconv.Atoi(c.Param("requests_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request id param")
		return
	}

	var input model.RequestStatus
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Request.ChangeStatus(
		model.RequestStatus{
			Id:       requestsId,
			StatusId: input.StatusId,
		},
	); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteRequest(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid student id param")
		return
	}

	if userId != id {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	requestsId, err := strconv.Atoi(c.Param("requests_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request id param")
		return
	}

	if err := h.services.Request.DeleteRequest(requestsId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getAllRequestsByWorkId(c *gin.Context) {
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

	workId, err := strconv.Atoi(c.Param("workId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid work id param")
		return
	}

	requests, err := h.services.Request.GetRequestsByWorkId(workId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllRequestsResponse{
		Data: requests,
	})
}
