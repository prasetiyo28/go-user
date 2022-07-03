package account

import (
	"net/http"

	"go-user/pkg"
	"go-user/presenter"

	"github.com/gin-gonic/gin"
)

type GetUserRequestBody struct {
	UserID string `json:"userId" validate:"required"`
}

func (handler *AccountController) DisplayAllUsers(ctx *gin.Context) {
	user, err := handler.as.GetUser()
	if err != nil {
		result := presenter.Response{
			Status: pkg.HTTP_STATUS_FAIL,
		}
		ctx.JSON(http.StatusCreated, result)
		return
	}
	result := presenter.Response{
		Status: pkg.HTTP_STATUS_SUCCESS,
		Data:   user,
	}
	ctx.JSON(http.StatusCreated, result)
}

func (handler *AccountController) DisplayUser(ctx *gin.Context) {
	var input GetUserRequestBody
	err := ValidateRequest(pkg.BIND_TYPE_JSON, ctx, &input)
	if err != nil {
		result := presenter.Response{
			Status:  pkg.HTTP_STATUS_FAIL,
			Message: err.Message,
			Data:    nil,
		}
		ctx.JSON(err.Code, result)
		return
	}

	user, err := handler.as.GetUserByUserID(input.UserID)
	if err != nil {
		result := presenter.Response{
			Status: pkg.HTTP_STATUS_FAIL,
		}
		ctx.JSON(http.StatusCreated, result)
		return
	}
	result := presenter.Response{
		Status: pkg.HTTP_STATUS_SUCCESS,
		Data:   user,
	}
	ctx.JSON(http.StatusCreated, result)
}
