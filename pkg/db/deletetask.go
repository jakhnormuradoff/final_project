package db

import "errors"

func DeleteTask(id string) error {
	result, err := db.Exec(`DELETE FROM scheduler WHERE id = ?`, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to get rows affected")
	}
	if rowsAffected == 0 {
		return errors.New("no task found with the given id")
	}
	return nil
}