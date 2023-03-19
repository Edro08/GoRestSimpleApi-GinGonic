package handler

import (
	"GoRestSimpleApi/internal/task/delete"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTaskHandler(taskDeleteService delete.TaskDeleteService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		req := TaskIDRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			err := taskDeleteService.Delete(ctx, req.Id)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
			} else {
				ctx.JSON(http.StatusAccepted, 1)
			}

		}
	}
}
