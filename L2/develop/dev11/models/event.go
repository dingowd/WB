package models

type Event struct {
	Owner     int    `json:"owner" db:"owner"`
	Title     string `json:"title" db:"title"`
	Descr     string `json:"descr" db:"descr"`
	StartDate string `json:"start_date" db:"start_date"`
	StartTime string `json:"start_time" db:"start_time"`
	EndDate   string `json:"end_date" db:"end_date"`
	EndTime   string `json:"end_time" db:"end_time"`
}

type DBEvent struct {
	ID        int    `json:"id" db:"id"`
	Owner     int    `json:"owner" db:"owner"`
	Title     string `json:"title" db:"title"`
	Descr     string `json:"descr" db:"descr"`
	StartDate string `json:"start_date" db:"start_date"`
	StartTime string `json:"start_time" db:"start_time"`
	EndDate   string `json:"end_date" db:"end_date"`
	EndTime   string `json:"end_time" db:"end_time"`
}
