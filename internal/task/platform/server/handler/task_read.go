package handler

import (
	"GoRestSimpleApi/internal/task/read"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaskReadAllHandler(taskGetAllService read.TaskGetAllService) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		tasks, err := taskGetAllService.ReadAll(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.JSON(http.StatusOK, tasks)
		}
	}
}

func TaskReadIDHandler(taskGetIdService read.TaskGetIdService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		req := TaskIDRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {

			tasks, err := taskGetIdService.ReadID(ctx, req.Id)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, tasks)
			}
		}
	}
}
