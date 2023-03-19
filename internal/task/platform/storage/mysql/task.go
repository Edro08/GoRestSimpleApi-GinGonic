package mysql

import "database/sql"

var sqlTaskTable string = "task"

type sqlTask struct {
	ID      string `db:"Id"`
	Name    string `db:"Name"`
	Content string `db:"Content"`
}

type TaskRepository struct {
	db *sql.DB
}

func TaskNewRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}
