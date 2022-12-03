package project

import (
	"myserver/model/project"
	"net/http"
	"time"
)

// struktur data project

// penampang data project

var ProjectData = []project.Project{
	{
		OwnerID:         1,
		Id:              0,
		Name:            "Proyek Terkadang",
		StartDate:       time.Now(),
		EndDate:         time.Now(),
		Description:     "Halo bang Aaoisjdoijoiujokmnokinmoijdoaijdoiankjnkjnee",
		MetaDescription: "Yes It Is!",
		Tech:            []string{"Golang", "NodeJS", "React", "Snowpack"},
		Duration:        "5 Months",
	},
	{
		OwnerID:         2,
		Id:              1,
		Name:            "Proyek Terkadang",
		StartDate:       time.Now(),
		EndDate:         time.Now(),
		Description:     "Halo bang",
		MetaDescription: "Yes It is!",
		Tech:            []string{"Golang", "NodeJS", "React", "Snowpack"},
		Duration:        "5 Months",
	},
}

func ProjectCtrl(w http.ResponseWriter, r *http.Request) {

}
