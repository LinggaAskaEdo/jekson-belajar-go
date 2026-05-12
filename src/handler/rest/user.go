package rest

import (
	"errors"
	"net/http"

	"belajar-go/src/domain"
	"belajar-go/src/dto"
	"belajar-go/src/util"

	"github.com/gin-gonic/gin"
)

func (e *rest) ListUsers(c *gin.Context) {

	var filter dto.UserFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		util.ResponseError(c, http.StatusBadRequest, "invalid parameter cause : "+err.Error())
		return
	}

	users, total, err := e.svc.User.ListAllDataUser(&filter)

	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError, "internal server error cause : "+err.Error())
		return
	}

	util.ResponseOk(c, &total, users)
}

func (e *rest) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var request []*domain.UserCreateDomain

	if err := c.ShouldBindJSON(&request); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			util.ResponseError(c, http.StatusRequestEntityTooLarge,
				"request body too large, maximum allowed is 2 MB",
			)
			return
		}
		util.ResponseError(c, http.StatusBadRequest, "invalid request body cous")
		return
	}

	users, err := e.svc.User.CreateDataUser(ctx, request)
	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError, "internal server error cause")
		return
	}

	util.ResponseOk(c, nil, users)
}
