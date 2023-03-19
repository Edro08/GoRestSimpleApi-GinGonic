package mysql

import (
	"GoRestSimpleApi/internal/task"
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

func (r TaskRepository) Update(ctx context.Context, task task.Task) error {
	query := sqlbuilder.Update(sqlTaskTable).Set(
		"Name = \""+task.Names().String()+"\"",
		"Content = \""+task.Contents().String()+"\"",
	).Where(
		"Id = " + task.IDs().String(), //+ task.IDs().String(),
	).String()

	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Error intentando: %w", err)
	}
	return nil
}
