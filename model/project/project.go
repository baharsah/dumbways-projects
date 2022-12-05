package project

import (
	"context"
	"fmt"
	"log"
	"math"
	"myserver/conn"
	"strings"
	"time"
)

type Project struct {
	OwnerID             int
	ProjectID           int
	OwnerName           string
	Id                  int
	Name                string
	StartDate           time.Time
	EndDate             time.Time
	Description         string
	MetaDescription     string
	Tech                []string
	Duration            int
	FileDir             string
	ParsedStartDate     string
	ParsedEndDate       string
	ParsedFormStartDate string
	ParsedFormEndDate   string
}

func SelectProject(project Project) []Project {

	conn.DatabaseConnect()
	if project.OwnerID == 0 {
		rows, err := conn.Conn.Query(context.Background(), `select projects.name ,  users.username , projects.id , projects.imagedir , projects.startdate , projects.enddate , projects.desc ,projects.stack , projects.ownerid from projects left join users on projects.ownerid =users.id`)
		log.Println(err)
		var result []Project
		for rows.Next() {
			var each = Project{}

			var err = rows.Scan(&each.Name, &each.OwnerName, &each.ProjectID, &each.FileDir, &each.StartDate, &each.EndDate, &each.Description, &each.Tech, &each.OwnerID)
			each.Duration = int(math.Round((each.EndDate.Sub(each.StartDate).Hours() / 24 / 30)))
			each.MetaDescription = strings.Split(each.Description, ".")[0]
			each.ParsedStartDate = each.StartDate.Format("02 January 2006")
			each.ParsedEndDate = each.EndDate.Format("02 January 2006")
			each.ParsedFormStartDate = each.StartDate.Format("2006-01-02")
			each.ParsedFormEndDate = each.EndDate.Format("2006-01-02")

			result = append(result, each)
			if err != nil {
				fmt.Println(err.Error())
			}

		}
		return result
	} else {
		// log.Println(project.OwnerID)
		rows, _ := conn.Conn.Query(context.Background(), `select projects.name ,  users.username , projects.id , projects.imagedir ,  projects.startdate , projects.enddate , projects.desc ,projects.stack , projects.ownerid  from projects left join users on projects.ownerid =users.id where projects.ownerid=$1`, project.OwnerID)
		var result []Project
		// log.Println(rows)
		for rows.Next() {
			var each = Project{}

			var err = rows.Scan(&each.Name, &each.OwnerName, &each.ProjectID, &each.FileDir, &each.StartDate, &each.EndDate, &each.Description, &each.Tech, &each.OwnerID)
			each.Duration = int(math.Round((each.EndDate.Sub(each.StartDate).Hours() / 24 / 30)))
			each.MetaDescription = strings.Split(each.Description, ".")[0]
			each.ParsedStartDate = each.StartDate.Format("02 January 2006")
			each.ParsedEndDate = each.EndDate.Format("02 January 2006")
			each.ParsedFormStartDate = each.StartDate.Format("2006-01-02")
			each.ParsedFormEndDate = each.EndDate.Format("2006-01-02")

			// log.Print(each)
			result = append(result, each)

			if err != nil {
				fmt.Println(err.Error())
			}

		}

		return result

	}

}

func CreateProject(project Project) error {

	// log.Println(project.StartDate)

	conn.DatabaseConnect()
	_, err := conn.Conn.Exec(context.Background(), `INSERT INTO projects(name , startdate , enddate, "desc", stack, ownerid, imagedir) values ($1 , $2 , $3 , $4 , $5, $6 , $7)`, project.Name, project.StartDate.Format("2006-01-02"), project.EndDate.Format("2006-01-02"), project.Description, project.Tech, project.OwnerID, project.FileDir)
	return err

}

func UpdateProject(project Project) error {
	conn.DatabaseConnect()
	_, err := conn.Conn.Exec(context.Background(), `UPDATE projects
	SET "name"=$1 , "startdate"=$2 ,  "enddate"=$3, "desc"=$4, "stack"=$5 , "ownerid"=$6 , "imagedir"=$7
	WHERE id=$8`, project.Name, project.StartDate.Format("2006-01-02"), project.EndDate.Format("2006-01-02"), project.Description, project.Tech, project.OwnerID, project.FileDir, project.ProjectID)
	log.Println(project.Name)
	log.Println(err)
	return err

}

func DeleteProject(idproject int) error {

	conn.DatabaseConnect()
	_, err := conn.Conn.Exec(context.Background(), `DELETE FROM public.projects WHERE id=$1`, idproject)
	log.Println(err)
	return err

}
