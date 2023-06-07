package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	roleCtx             = "role"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == " " {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
	c.Set(roleCtx, role)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid type user id")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

func (h *Handler) onlyInstructor(c *gin.Context) {
	role, ok := c.Get(roleCtx)
	if !ok {
		newErrorResponse(c, http.StatusForbidden, "user role not found")
	}

	roleString, ok := role.(string)
	if !ok {
		newErrorResponse(c, http.StatusForbidden, "invalid type user role")
	}
	if roleString != "instructor" {
		newErrorResponse(c, http.StatusForbidden, "permission denied")
	}

}

func (h *Handler) onlyStudent(c *gin.Context) {
	role, ok := c.Get(roleCtx)
	if !ok {
		newErrorResponse(c, http.StatusForbidden, "user role not found")
	}

	roleString, ok := role.(string)
	if !ok {
		newErrorResponse(c, http.StatusForbidden, "invalid type user role")
	}
	if roleString != "student" {
		newErrorResponse(c, http.StatusForbidden, "permission denied")
	}

}
