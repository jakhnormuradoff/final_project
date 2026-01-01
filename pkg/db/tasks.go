package db

import "strconv"

func Tasks(limit int) ([]Task, error) {
	rows, err := db.Query(`SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var idInt int64
		err := rows.Scan(&idInt, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		task.ID = strconv.FormatInt(idInt, 10)
		tasks = append(tasks, task)
	}
	if tasks == nil {
		tasks = []Task{}
	}
	return tasks, nil
}