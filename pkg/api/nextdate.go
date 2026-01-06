package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const formatted = "20060102"

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	now := r.FormValue("now")
	if now == "" {
		now = time.Now().Format(formatted)
	}
	parseNow, err := time.Parse(formatted, now)
	if err != nil {
		http.Error(w, "failed to parse date", http.StatusBadRequest)
		return
	}
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	result, err := NextDate(parseNow, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = fmt.Fprint(w, result)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	
	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	repeat = strings.TrimSpace(repeat)

	parsedDstart, err := time.Parse(formatted, dstart)
	if err != nil {
		return "", errors.New("failed to parse dstart")
	}

	if now.IsZero() {
		now = parsedDstart
	}

	now, err = time.Parse(formatted,now.Format(formatted))
	if err != nil {
		return "", errors.New("failed to parse now")
	}
	
	if repeat == "y" {
		next := parsedDstart.AddDate(1, 0, 0)
		if parsedDstart.Month() == time.February &&
		 parsedDstart.Day() == 29 &&
		 next.Month() == time.February && 
		 next.Day() == 28 {
			next = next.AddDate(0, 0, 1)
		}
		for next.Before(now) || next.Equal(now) {
			next = next.AddDate(1, 0, 0)
			if parsedDstart.Month() == time.February &&
		 parsedDstart.Day() == 29 &&
		 next.Month() == time.February && 
		 next.Day() == 28 {
			next = next.AddDate(0, 0, 1)
		}	
		}
		return next.Format(formatted), nil
	}

	parts := strings.Fields(repeat)

	 if len(parts) == 2 && parts[0] == "d" {
		day, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", errors.New("failed to parse day to int")
		}
		if day >= 1 && day <= 400 {
			next := parsedDstart.AddDate(0, 0, day)
			for next.Before(now) || next.Equal(now) {
				next = next.AddDate(0, 0, day)
			}
			return next.Format(formatted), nil
		}
		return "", errors.New("day is out of range")
	}
	
	return "", errors.New("invalid repeat format")

}
