package courses

import (
	"api_go/internal/create"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCourseRequest struct {
	Id       string `json:id binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateCourseHandler(courseCreator create.CourseCreator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createCourseRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := courseCreator.CreateCourse(ctx, req.Id, req.Name, req.Duration)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
