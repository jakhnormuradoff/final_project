package db

import "strconv"

func TaskByID(id string) (*Task, error) {
	row := db.QueryRow(`SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`, id)
	var task Task
	var idInt int64
	err := row.Scan(&idInt, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}
	task.ID = strconv.FormatInt(idInt, 10)
	return &task, nil
}
