package mysql

import (
	"GoRestSimpleApi/internal/task"
	"context"
	"fmt"
)

func (r TaskRepository) Delete(ctx context.Context, Id string) error {

	_, err := task.NewTaskID(Id)

	if err != nil {
		return err
	} else {
		sqlquery := "delete from task where Id = " + Id
		_, err = r.db.ExecContext(ctx, sqlquery)
		if err != nil {
			return fmt.Errorf("Error intentando: %w", err)
		}
		return nil
	}

}
