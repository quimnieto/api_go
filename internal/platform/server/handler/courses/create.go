package courses

import (
	"api_go/internal/create"
	"api_go/kit/command"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCourseRequest struct {
	Id       string `json:id binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateCourseHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createCourseRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		command := create.NewCreateCourseCommand(req.Id, req.Name, req.Duration)
		err := commandBus.Dispatch(ctx, command)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
