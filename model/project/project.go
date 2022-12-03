package project

import "time"

type Project struct {
	OwnerID         int
	Id              int
	Name            string
	StartDate       time.Time
	EndDate         time.Time
	Description     string
	MetaDescription string
	Tech            []string
	Duration        string
	FileDir         string
}

func GetAll() {

}
