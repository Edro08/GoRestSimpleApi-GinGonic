package mysql

import (
	"GoRestSimpleApi/internal/task"
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

func (r TaskRepository) Save(ctx context.Context, task task.Task) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlTask))
	query, arg := courseSQLStruct.InsertInto(sqlTaskTable, sqlTask{
		ID:      task.IDs().String(),
		Name:    task.Names().String(),
		Content: task.Contents().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, arg...)
	if err != nil {
		return fmt.Errorf("Error intentando: %w", err)
	}
	return nil
}
