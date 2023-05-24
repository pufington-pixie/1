package models

import "time"

type Project struct {
	ID        int
	Name      string
	Title     string
	Date      time.Time
	SAPNumber string
	Notes     string
	BranchID  int
	StatusID  int
	Services  []Service
}

type Service struct {
	ID   int
	Name string
}
