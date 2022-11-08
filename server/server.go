package server

import (
	"db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"structures"
	"text/template"
	"time"
)

func Plan2(w http.ResponseWriter, r *http.Request) {
	var year, _ = strconv.Atoi(r.FormValue("year"))
	var group, _ = strconv.Atoi(r.FormValue("group"))
	var semestr, _ = strconv.Atoi(r.FormValue("semestr"))
	fmt.Println("Получен запрос: " + r.Method + " Specialities2")
	switch r.Method {
	case "GET":
		data, err := db.OutputPlan(year, group, semestr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		pl, err := template.ParseFiles("./templates/plan2.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = pl.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Plan(w http.ResponseWriter, r *http.Request) {
	var year, _ = strconv.Atoi(r.FormValue("year"))
	var group, _ = strconv.Atoi(r.FormValue("group"))
	var semestr, _ = strconv.Atoi(r.FormValue("semestr"))
	fmt.Println("Получен запрос: " + r.Method + " Specialities2")
	switch r.Method {
	case "GET":
		data, err := db.OutputPlan(year, group, semestr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		pl, err := template.ParseFiles("./templates/plan.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = pl.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Itogo(w http.ResponseWriter, r *http.Request) {
	var year, _ = strconv.Atoi(r.FormValue("year"))
	var group, _ = strconv.Atoi(r.FormValue("group"))
	fmt.Println("Получен запрос: " + r.Method + " Specialities2")
	switch r.Method {
	case "GET":
		data, err := db.OutputItogo(year, group)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		i, err := template.ParseFiles("./templates/itog.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = i.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Itogo2(w http.ResponseWriter, r *http.Request) {
	var year, _ = strconv.Atoi(r.FormValue("year"))
	var group, _ = strconv.Atoi(r.FormValue("group"))
	fmt.Println("Получен запрос: " + r.Method + " Specialities2")
	switch r.Method {
	case "GET":
		data, err := db.OutputItogo2(year, group)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		i2, err := template.ParseFiles("./templates/itog2.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = i2.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func TimetableGroups2(w http.ResponseWriter, r *http.Request) {
	var Month, _ = strconv.Atoi(r.FormValue("Month"))
	var Year, _ = strconv.Atoi(r.FormValue("Year"))
	var TeacherID, _ = strconv.Atoi(r.FormValue("TeacherId"))
	fmt.Println("Получен запрос: " + r.Method + " TimetablesGroups")
	switch r.Method {
	case "GET":
		type Result struct {
			TimetableGroups []structures.TimetableGroups
		}
		var Res Result
		var err error
		Res.TimetableGroups, err = db.OutputTimetableGroups2(Month, Year, TeacherID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for jdx := 0; jdx < len(Res.TimetableGroups); jdx++ {
			rows, err := db.Db.Query("SELECT date_part('day',date_of_timetable), hours  FROM timetable2 where id_load = $1 and date_part('month',date_of_timetable)=$2 and date_part('year',date_of_timetable)=$3 and type_of_lesson = $4", Res.TimetableGroups[jdx].IdLoad, Month, Year, Res.TimetableGroups[jdx].Type)
			if err != nil {
				log.Fatal(err.Error())
				panic(err)
			}
			for rows.Next() {
				ttg := structures.DayNHour{}
				err := rows.Scan(&ttg.Day, &ttg.Hours)
				if err != nil {
					fmt.Println(err)
					continue
				}
				Res.TimetableGroups[jdx].DayNHour = append(Res.TimetableGroups[jdx].DayNHour, ttg)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		tt, err := json.Marshal(Res.TimetableGroups)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(tt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Passwords2(w http.ResponseWriter, r *http.Request) {
	var Login = r.FormValue("Login")
	var Password = r.FormValue("Password")
	switch r.Method {
	case "POST":
		data, err := db.OutputPasswords2(Login, Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		pw2, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(pw2)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func TimetablesGroups(w http.ResponseWriter, r *http.Request) {
	var Month, _ = strconv.Atoi(r.FormValue("Month"))
	var Year, _ = strconv.Atoi(r.FormValue("Year"))
	var GroupId, _ = strconv.Atoi(r.FormValue("GroupId"))
	var TimeOfTimetable, _ = time.Parse("2006-01-02", r.FormValue("DateOfTimetable"))
	var Hours, _ = strconv.Atoi(r.FormValue("Hours"))
	var IdLoad, _ = strconv.Atoi(r.FormValue("IdLoad"))
	var Type, _ = strconv.ParseBool(r.FormValue("Type"))
	fmt.Println("Получен запрос: " + r.Method + " TimetablesGroups")
	switch r.Method {
	case "GET":
		type Result struct {
			TimetableGroups []structures.TimetableGroups
		}
		var Res Result
		var err error
		Res.TimetableGroups, err = db.OutputTimetableGroups(Month, Year, GroupId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for jdx := 0; jdx < len(Res.TimetableGroups); jdx++ {
			rows, err := db.Db.Query("SELECT date_part('day',date_of_timetable), hours  FROM timetable2 where id_load = $1 and date_part('month',date_of_timetable)=$2 and date_part('year',date_of_timetable)=$3 and type_of_lesson = $4", Res.TimetableGroups[jdx].IdLoad, Month, Year, Res.TimetableGroups[jdx].Type)
			if err != nil {
				log.Fatal(err.Error())
				panic(err)
			}
			for rows.Next() {
				ttg := structures.DayNHour{}
				err := rows.Scan(&ttg.Day, &ttg.Hours)
				if err != nil {
					fmt.Println(err)
					continue
				}
				Res.TimetableGroups[jdx].DayNHour = append(Res.TimetableGroups[jdx].DayNHour, ttg)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		tt, err := json.Marshal(Res.TimetableGroups)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(tt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "POST":
		_, err := ioutil.ReadAll(r.Body)
		err = db.InputTimetable2(Type, IdLoad, TimeOfTimetable, Hours)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		_, err := ioutil.ReadAll(r.Body)
		if Hours == 0 {
			err = db.DeleteTimetable2(IdLoad, TimeOfTimetable, Type)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			err = db.UpdateTimetable2(IdLoad, TimeOfTimetable, Type, Hours)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

func SpecialitiesJSON2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " SpecialitiesJSON2")
	switch r.Method {
	case "GET":
		data, err := db.OutputSpecialities2()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err1, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(err1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Specialities2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " Specialities2")
	switch r.Method {
	case "GET":
		data, err := db.OutputSpecialities2()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		s, err := template.ParseFiles("./templates/specialities2.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = s.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var s structures.Speciality2
		err = json.Unmarshal(body, &s)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if s.Speciality != "" {
			regex := `^[а-яА-яa-zA-z]*$`
			matched, _ := regexp.Match(regex, []byte(s.Speciality))
			if matched == true {
				err = db.InputSpecialities2(s)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var s structures.Speciality2
		err = json.Unmarshal(body, &s)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.DeleteSpecialities2(s)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var s structures.Speciality2
		err = json.Unmarshal(body, &s)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if s.Speciality != "" {
			regex := `^[а-яА-яa-zA-z]*$`
			matched, _ := regexp.Match(regex, []byte(s.Speciality))
			if matched == true {
				err = db.UpdateSpecialities2(s)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Groups2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " Groups2")
	switch r.Method {
	case "GET":

		type Result struct {
			Groups     []structures.Group2
			Speciality []structures.Speciality2
		}
		var Res Result
		var err error
		Res.Groups, err = db.OutputGroups2()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		g, err := template.ParseFiles("./templates/group.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		Res.Speciality, err = db.OutputSpecialities2()
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = g.Execute(w, Res)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var gro structures.Group2
		err = json.Unmarshal(body, &gro)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if gro.GroupNew == "" {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			err = db.InputGroups2(gro)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var gro structures.Group2
		err = json.Unmarshal(body, &gro)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.DeleteGroups2(gro)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var gro structures.Group2
		err = json.Unmarshal(body, &gro)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if gro.GroupNew == "" {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			err = db.UpdateGroups2(gro)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

func Disciplines2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " Disciplines2")
	switch r.Method {
	case "GET":
		data, err := db.OutputDisciplines2()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		d, err := template.ParseFiles("./templates/discipline.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = d.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var dis structures.Discipline2
		err = json.Unmarshal(body, &dis)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if dis.Disciplines != "" {
			regex := `^[а-яА-яa-zA-z]*$`
			matched, _ := regexp.Match(regex, []byte(dis.Disciplines))
			if matched == true {
				err = db.InputDisciplines2(dis)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var dis structures.Discipline2
		err = json.Unmarshal(body, &dis)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.DeleteDisciplines2(dis)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var dis structures.Discipline2
		err = json.Unmarshal(body, &dis)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if dis.Disciplines != "" {
			regex := `^[а-яА-яa-zA-z]*$`
			matched, _ := regexp.Match(regex, []byte(dis.Disciplines))
			if matched == true {
				err = db.UpdateDisciplines2(dis)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Teachers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " Teachers")
	switch r.Method {
	case "GET":
		type Result struct {
			Teachers []structures.Teacher
			Postt    []structures.Post
		}
		var Res Result
		var err error
		Res.Teachers, err = db.OutputTeachers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		t, err := template.ParseFiles("./templates/teacher.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		Res.Postt, err = db.OutputPosts()
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, Res)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var tec structures.Teacher
		err = json.Unmarshal(body, &tec)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err2 := false
		regex := `^[а-яА-яa-zA-z]+$`
		matched, _ := regexp.Match(regex, []byte(tec.Name))
		matched2, _ := regexp.Match(regex, []byte(tec.Surname))
		if matched == false || matched2 == false {
			w.WriteHeader(http.StatusInternalServerError)
			err2 = true
		}
		regex2 := `^\+?[78][-\(]?\d{3}\)?-?\d{3}-?\d{2}-?\d{2}$`
		matched3, _ := regexp.Match(regex2, []byte(tec.Telephone))
		if matched3 != true {
			if tec.Telephone != "" {
				w.WriteHeader(http.StatusInternalServerError)
				err2 = true
			}
		}
		regex3 := `^\w+@\w+\.\w+$`
		matched4, _ := regexp.Match(regex3, []byte(tec.Email))
		if matched4 != true {
			if tec.Email != "" {
				w.WriteHeader(http.StatusInternalServerError)
				err2 = true
			}
		}
		regex4 := `^\d{4}$`
		regex5 := `^\d{6}$`
		matched5, _ := regexp.Match(regex4, []byte(tec.PassportSeries))
		matched6, _ := regexp.Match(regex5, []byte(tec.IdPassport))
		if matched5 != true || matched6 != true {
			if tec.PassportSeries != "" || tec.IdPassport != "" {
				w.WriteHeader(http.StatusInternalServerError)
				err2 = true
			}
		}
		if tec.DateOfBirth != time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) || tec.DateOfHiring != time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) {
			if tec.DateOfBirth != time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) && tec.DateOfHiring == time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) {
				t1 := time.Now().AddDate(-18, 0, 0)
				t2 := tec.DateOfBirth.After(t1)
				if t2 == true {
					w.WriteHeader(http.StatusInternalServerError)
					err2 = true
				}
			} else {
				t1 := tec.DateOfHiring.AddDate(-18, 0, 0)
				t2 := tec.DateOfBirth.After(t1)
				if t2 == true {
					w.WriteHeader(http.StatusInternalServerError)
					err2 = true
				}
			}
		}
		if err2 == false {
			err = db.InputTeachers(tec)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var tec structures.Teacher
		err = json.Unmarshal(body, &tec)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.DeleteTeachers(tec)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var tec structures.Teacher
		err = json.Unmarshal(body, &tec)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err2 := false
		regex := `^[а-яА-яa-zA-z]+$`
		matched, _ := regexp.Match(regex, []byte(tec.Name))
		matched2, _ := regexp.Match(regex, []byte(tec.Surname))
		if matched == false || matched2 == false {
			w.WriteHeader(http.StatusInternalServerError)
			err2 = true
		}
		regex2 := `^\+?[78][-\(]?\d{3}\)?-?\d{3}-?\d{2}-?\d{2}$`
		matched3, _ := regexp.Match(regex2, []byte(tec.Telephone))
		if matched3 != true {
			if tec.Telephone != "" {
				w.WriteHeader(http.StatusInternalServerError)
				err2 = true
			}
		}
		regex3 := `^\w+@\w+\.\w+$`
		matched4, _ := regexp.Match(regex3, []byte(tec.Email))
		if matched4 != true {
			if tec.Email != "" {
				w.WriteHeader(http.StatusInternalServerError)
				err2 = true
			}
		}
		regex4 := `^\d{4}$`
		regex5 := `^\d{6}$`
		matched5, _ := regexp.Match(regex4, []byte(tec.PassportSeries))
		matched6, _ := regexp.Match(regex5, []byte(tec.IdPassport))
		if matched5 != true || matched6 != true {
			if tec.PassportSeries != "" || tec.IdPassport != "" {
				w.WriteHeader(http.StatusInternalServerError)
				err2 = true
			}
		}
		if tec.DateOfBirth != time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) || tec.DateOfHiring != time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) {
			if tec.DateOfBirth != time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) && tec.DateOfHiring == time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) {
				t1 := time.Now().AddDate(-18, 0, 0)
				t2 := tec.DateOfBirth.After(t1)
				if t2 == true {
					w.WriteHeader(http.StatusInternalServerError)
					err2 = true
				}
			} else {
				t1 := tec.DateOfHiring.AddDate(-18, 0, 0)
				t2 := tec.DateOfBirth.After(t1)
				if t2 == true {
					w.WriteHeader(http.StatusInternalServerError)
					err2 = true
				}
			}
		}
		if err2 == false {
			err = db.UpdateTeachers(tec)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Loads(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " Loads")
	switch r.Method {
	case "GET":
		type Result struct {
			Loads      []structures.Load
			Discipline []structures.Discipline2
			Group      []structures.Group2
			Teacher    []structures.Teacher
		}
		var Res Result
		var err error
		Res.Loads, err = db.OutputLoads()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for jdx := 0; jdx < len(Res.Loads); jdx++ {
			for i := 0; i <= 10; i++ {
				if i == 0 {
					Res.Loads[jdx].Yearss = append(Res.Loads[jdx].Yearss, Res.Loads[jdx].Year)
				} else {
					Res.Loads[jdx].Yearss = append(Res.Loads[jdx].Yearss, Res.Loads[jdx].Year-5+i)
				}
			}
		}
		w.Header().Set("Content-Type", "text/html")
		l, err := template.ParseFiles("./templates/load.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		Res.Discipline, err = db.OutputDisciplines2()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		Res.Group, err = db.OutputGroups2()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		Res.Teacher, err = db.OutputTeachers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = l.Execute(w, Res)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var loa structures.Load
		err = json.Unmarshal(body, &loa)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.InputLoads(loa)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var loa structures.Load
		err = json.Unmarshal(body, &loa)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.DeleteLoads(loa)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var loa structures.Load
		err = json.Unmarshal(body, &loa)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.UpdateLoads(loa)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func Posts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " Posts")
	switch r.Method {
	case "GET":
		data, err := db.OutputPosts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		p, err := template.ParseFiles("./templates/post.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = p.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var p structures.Post
		err = json.Unmarshal(body, &p)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if p.Posts != "" {
			regex := `^[а-яА-яa-zA-z]*$`
			matched, _ := regexp.Match(regex, []byte(p.Posts))
			if matched == true {
				err = db.InputPosts(p)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var p structures.Post
		err = json.Unmarshal(body, &p)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.DeletePosts(p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var p structures.Post
		err = json.Unmarshal(body, &p)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if p.Posts != "" {
			regex := `^[а-яА-яa-zA-z]*$`
			matched, _ := regexp.Match(regex, []byte(p.Posts))
			if matched == true {
				err = db.UpdatePosts(p)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func PostsJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " PostsJSON")
	switch r.Method {
	case "GET":
		data, err := db.OutputPosts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err1, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(err1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func DiciplineJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " DiciplineJSON")
	switch r.Method {
	case "GET":
		data, err := db.OutputDisciplines2()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err1, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(err1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func GroupJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " GroupJSON")
	switch r.Method {
	case "GET":
		data, err := db.OutputGroups2()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err1, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(err1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func TeacherJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method + " TeacherJSON")
	switch r.Method {
	case "GET":
		data, err := db.OutputTeachers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err1, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(err1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

func LoadList(w http.ResponseWriter, r *http.Request) {
	var t, _ = strconv.Atoi(r.FormValue("t"))
	var t1, _ = strconv.Atoi(r.FormValue("t1"))
	fmt.Println("Получен запрос: " + r.Method + " LoadList")
	switch r.Method {
	case "GET":
		data1, err := db.OutputLoadList(t, t1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		l2, err := template.ParseFiles("./templates/loadlist.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = l2.Execute(w, data1)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
