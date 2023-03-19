package handler

import (
	"GoRestSimpleApi/internal/task/create"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewTaskHandler(taskCreateService create.TaskCreateService) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		req := TaskRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {

			err = taskCreateService.Save(ctx, strconv.Itoa(req.Id), req.Name, req.Content)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.Status(http.StatusAccepted)
			}
		}
	}
}
