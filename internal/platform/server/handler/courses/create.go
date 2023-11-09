package courses

import (
	mooc "api_go/internal"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCourseRequest struct {
	Id       string `json:id binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateCourseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createCourseRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course, err := mooc.NewCourse(req.Id, req.Name, req.Duration)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println(course)

		ctx.Status(http.StatusCreated)
	}
}
