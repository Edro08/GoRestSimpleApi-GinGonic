package mysql

import (
	"GoRestSimpleApi/internal/task"
	"context"
)

func (r TaskRepository) ReadAll(ctx context.Context) ([]task.TaskDB, error) {

	sqlquery := "Select * from task"
	rows, err := r.db.Query(sqlquery)

	if err != nil {
		return nil, err
	}

	tasks := make([]task.TaskDB, 0)

	for rows.Next() {
		var t task.TaskDB
		err := rows.Scan(&t.ID, &t.Name, &t.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r TaskRepository) ReadID(ctx context.Context, Id string) ([]task.TaskDB, error) {

	sqlquery := "Select * from task where Id = " + Id
	rows, err := r.db.Query(sqlquery)

	if err != nil {
		return nil, err
	}

	tasks := make([]task.TaskDB, 0)

	for rows.Next() {
		var t task.TaskDB
		err := rows.Scan(&t.ID, &t.Name, &t.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
