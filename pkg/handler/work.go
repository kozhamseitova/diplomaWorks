package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/diplomaWorks/pkg/model"
	"net/http"
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

//func (h *Handler) getAllWorksById(c *gin.Context) {
//	userId, err := getUserId(c)
//	if err != nil {
//		return
//	}
//
//	works, err := h.services.Work.GetAll(userId)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//}

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
