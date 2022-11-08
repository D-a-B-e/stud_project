package main

import (
	"database/sql"
	"db"
	_ "github.com/lib/pq"
	"net/http"
	"server"
)

func main() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=1 dbname=praktika1 sslmode=disable"
	db.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Db.Close()
	http.Handle("/", http.FileServer(http.Dir("htdocs")))
	http.HandleFunc("/plan2", server.Plan2)
	http.HandleFunc("/plan", server.Plan)
	http.HandleFunc("/itog", server.Itogo)
	http.HandleFunc("/itog2", server.Itogo2)
	http.HandleFunc("/passwords2", server.Passwords2)
	http.HandleFunc("/timetablegroups2", server.TimetableGroups2)
	http.HandleFunc("/specialities2", server.Specialities2)
	http.HandleFunc("/disciplines", server.Disciplines2)
	http.HandleFunc("/groups", server.Groups2)
	http.HandleFunc("/timetablegroups", server.TimetablesGroups)
	http.HandleFunc("/teachers", server.Teachers)
	http.HandleFunc("/loads", server.Loads)
	http.HandleFunc("/loadlist", server.LoadList)
	http.HandleFunc("/posts", server.Posts)
	http.HandleFunc("/DiciplineJSON", server.DiciplineJSON)
	http.HandleFunc("/PostsJSON", server.PostsJSON)
	http.HandleFunc("/GroupJSON", server.GroupJSON)
	http.HandleFunc("/SkedJSON", server.GroupJSON)
	http.HandleFunc("/ScheduleJSON", server.TeacherJSON)
	http.HandleFunc("/TeacherJSON", server.TeacherJSON)
	http.HandleFunc("/SpecialitiesJSON2", server.SpecialitiesJSON2)
	http.ListenAndServe("0.0.0.0:8081", nil)
}
