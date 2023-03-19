package bootstrap

import (
	"GoRestSimpleApi/internal/task/create"
	"GoRestSimpleApi/internal/task/delete"
	"GoRestSimpleApi/internal/task/platform/server/handler"
	"GoRestSimpleApi/internal/task/platform/storage/mysql"
	"GoRestSimpleApi/internal/task/read"
	"GoRestSimpleApi/internal/task/update"
	"GoRestSimpleApi/kit/platform/server"
	mysqlglobal "GoRestSimpleApi/kit/platform/storage/mysql"
)

const (
	host = "192.168.1.3"
	port = 8000
)

func Run() error {

	srv := server.New(host, port)

	mysqlURI := mysqlglobal.GetmysqlURI()
	db, err := mysqlglobal.SetupStorage(mysqlURI)

	if err != nil {
		return err
	}

	taskRepository := mysql.TaskNewRepository(db)
	taskGetIdService := read.NewTaskServiceID(taskRepository)
	taskGetAllService := read.NewTaskService(taskRepository)
	newtaskService := create.NewTaskService(taskRepository)
	updatetaskService := update.NewTaskService(taskRepository)
	deletetaskService := delete.NewTaskService(taskRepository)

	srv.RegisterRoute("GET", "/getIdtask", handler.TaskReadIDHandler(taskGetIdService))
	srv.RegisterRoute("GET", "/getalltask", handler.TaskReadAllHandler(taskGetAllService))
	srv.RegisterRoute("POST", "/newtask", handler.NewTaskHandler(newtaskService))
	srv.RegisterRoute("POST", "/updatetask", handler.UpdateTaskHandler(updatetaskService))
	srv.RegisterRoute("POST", "/deletetask", handler.DeleteTaskHandler(deletetaskService))

	return srv.Run()
}
