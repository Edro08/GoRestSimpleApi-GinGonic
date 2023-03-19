package handler

import (
	"GoRestSimpleApi/internal/task"
	"GoRestSimpleApi/internal/task/update"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(taskUpdateService update.TaskUpdateService) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		req := TaskRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {

			task, err := task.NewTask(strconv.Itoa(req.Id), req.Name, req.Content)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
			} else {

				err = taskUpdateService.Update(ctx, task)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, err.Error())
				} else {
					ctx.Status(http.StatusAccepted)
				}
			}
		}
	}
}
