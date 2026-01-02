package db

import (
	"errors"
)

func UpdateTask(task *Task) error {
	update, err := db.Exec(`UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`,
		task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err	
	}
	if update != nil {
		if rows, _ := update.RowsAffected(); rows == 0 {
			return errors.New("no task found with the given ID")
		}
	}
	return nil
}
