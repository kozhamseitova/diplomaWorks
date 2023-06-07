package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"net/http"
)

func (h *Handler) createRequest(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var student model.Student
	student, err = h.services.User.GetStudentByUserId(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.RequestInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Request.CreateRequest(model.RequestInput{
		StudentId:   student.Id,
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
