package db

import "errors"

func AddTask(task *Task) (int64, error) {
	result, err := db.Exec(`INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("failed to get last insert id")
	}
	return id, nil
}