package db

import "errors"

func UpdateDate(date string, id string) error {
	result, err := db.Exec(`UPDATE scheduler SET date = ? WHERE id = ?`, date, id)
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