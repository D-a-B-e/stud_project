package db

import (
	"database/sql"
	"fmt"
	"log"
	"structures"
	"time"
)

var Db *sql.DB

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullInt(s int32) sql.NullInt32 {
	if s == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: s,
		Valid: true,
	}
}

func OutputItogo2(year int, group int) ([]structures.Itogo2, error) {
	rows, err := Db.Query("select disciplines, COALESCE(fio, 'Вакансия'), plan, january, february, march, april, may, june, give, remain from itogo2 where year = $1 and id_group= $2", year, group)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var itogo2 []structures.Itogo2

	for rows.Next() {
		i2 := structures.Itogo2{}
		err := rows.Scan(&i2.Disciplines, &i2.Fio, &i2.Plan, &i2.January, &i2.February, &i2.March, &i2.April, &i2.May, &i2.June, &i2.Give, &i2.Remain)
		if err != nil {
			fmt.Println(err)
			continue
		}
		itogo2 = append(itogo2, i2)
	}
	return itogo2, err
}

func OutputItogo(year int, group int) ([]structures.Itogo, error) {
	rows, err := Db.Query("select disciplines, COALESCE(fio, 'Вакансия'), plan, september, october, november, december, give, remain  from itogo where year = $1 and id_group= $2", year, group)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var itogo []structures.Itogo

	for rows.Next() {
		i := structures.Itogo{}
		err := rows.Scan(&i.Disciplines, &i.Fio, &i.Plan, &i.September, &i.October, &i.November, &i.December, &i.Give, &i.Remain)
		if err != nil {
			fmt.Println(err)
			continue
		}
		itogo = append(itogo, i)
	}
	return itogo, err
}

func OutputPlan(year int, group int, semestr int) ([]structures.Plan, error) {
	rows, err := Db.Query("select COALESCE(fio, 'Вакансия'), disciplines, hours from real_plan where year = $1 and id_group= $2 and semestr = $3 ", year, group, semestr)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var plan []structures.Plan

	for rows.Next() {
		pl := structures.Plan{}
		err := rows.Scan(&pl.Fio, &pl.Disciplines, &pl.Hours)
		if err != nil {
			fmt.Println(err)
			continue
		}
		plan = append(plan, pl)
	}
	return plan, err
}

func OutputPasswords2(Login string, Password string) ([]structures.Passwords2, error) {
	rows, err := Db.Query("select * from real_password where login=$1 and password=$2", Login, Password)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var password []structures.Passwords2

	for rows.Next() {
		pw2 := structures.Passwords2{}
		err := rows.Scan(&pw2.Login, &pw2.Password, &pw2.IdTeacher, &pw2.Fio, &pw2.Admin)
		if err != nil {
			fmt.Println(err)
			continue
		}
		password = append(password, pw2)
	}
	return password, err
}

func OutputTimetableGroups2(Month int, Year int, IdTeacher int) ([]structures.TimetableGroups, error) {
	rows, err := Db.Query("select id_group, month, disciplines, COALESCE(fio, 'Вакансия'), all_in_hours, remain_hours, total_hours, year, "+
		"id_load, type, group_old from real_academic_year where month=$1 and year=$2 and id_teacher=$3", Month, Year,
		IdTeacher) //  нужно написать параметры которые передаются из страницы
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var timetable []structures.TimetableGroups

	for rows.Next() {
		ttg := structures.TimetableGroups{}
		err := rows.Scan(&ttg.IdGroup, &ttg.Month, &ttg.Disciplines, &ttg.Fio, &ttg.AllInHours, &ttg.RemainHours, &ttg.TotalHours, &ttg.Year, &ttg.IdLoad, &ttg.Type, &ttg.GroupOld)
		if err != nil {
			fmt.Println(err)
			continue
		}
		timetable = append(timetable, ttg)
	}
	return timetable, err
}

func OutputTimetableGroups(Month int, Year int, GroupId int) ([]structures.TimetableGroups, error) {
	rows, err := Db.Query("select id_group, month, disciplines, COALESCE(fio, 'Вакансия'), all_in_hours, remain_hours, total_hours, year, "+
		"id_load, type from real_academic_year where month=$1 and year=$2 and id_group=$3", Month, Year, GroupId)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var timetable []structures.TimetableGroups

	for rows.Next() {
		ttg := structures.TimetableGroups{}
		err := rows.Scan(&ttg.IdGroup, &ttg.Month, &ttg.Disciplines, &ttg.Fio, &ttg.AllInHours, &ttg.RemainHours, &ttg.TotalHours, &ttg.Year, &ttg.IdLoad, &ttg.Type)
		if err != nil {
			fmt.Println(err)
			continue
		}
		timetable = append(timetable, ttg)
	}
	return timetable, err
}

func InputTimetable2(Type bool, IdLoad int, DateOfTimetable time.Time, Hours int) error {
	_, err := Db.Exec("insert into timetable2 (type_of_lesson, id_load, date_of_timetable,  hours) values ($1, $2, "+
		"$3,"+
		" $4)", Type, IdLoad, DateOfTimetable, Hours)
	return err
}

func DeleteTimetable2(IdLoad int, DateOfTimetable time.Time, Type bool) error {
	_, err := Db.Exec("delete from timetable2 where id_load=$1 and date_of_timetable=$2 and type_of_lesson=$3",
		IdLoad, DateOfTimetable, Type)
	return err
}

func UpdateTimetable2(IdLoad int, DateOfTimetable time.Time, Type bool, Hours int) error {
	_, err := Db.Exec("update timetable2 set hours=$4 where id_load=$1 and date_of_timetable=$2 and type_of_lesson=$3",
		IdLoad, DateOfTimetable, Type, Hours)
	return err
}

func InputGroups2(g structures.Group2) error {
	_, err := Db.Exec("insert into grupi2 (group_old, group_new, form, amount, id_speciality) values ($1, $2, $3, $4,"+
		" $5) ", g.GroupOld, g.GroupNew, g.Form, g.Amount, g.IdSpeciality)
	return err
}

func DeleteGroups2(g structures.Group2) error {
	_, err := Db.Exec("delete from grupi2 where id_group = $1", g.IdGroup)
	return err
}

func UpdateGroups2(g structures.Group2) error {
	_, err := Db.Exec("update grupi2 set group_old=$2, group_new=$3, form=$4, amount=$5, "+
		"id_speciality=$6 where id_group = $1", g.IdGroup, g.GroupOld, g.GroupNew, g.Form, g.Amount, g.IdSpeciality)
	return err
}

func OutputGroups2() ([]structures.Group2, error) {
	rows, err := Db.Query("select * from public.grupi2 order by id_group")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var groups []structures.Group2

	for rows.Next() {
		g := structures.Group2{}
		err := rows.Scan(&g.IdGroup, &g.GroupOld, &g.GroupNew, &g.Form, &g.Amount, &g.IdSpeciality)
		if err != nil {
			fmt.Println(err)
			continue
		}
		groups = append(groups, g)
	}

	return groups, err
}

func InputDisciplines2(d structures.Discipline2) error {
	_, err := Db.Exec("insert into disciplines2 (disciplines, hours_of_lecture, hours_of_practice, course, advice, "+
		"exam, hours_of_lecture2, hours_of_practice2, course2, advice2, exam2) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ", d.Disciplines, d.HoursOfLecture, d.HoursOfPractice, d.Course, d.Advice, d.Exam, d.HoursOfLecture2, d.HoursOfPractice2, d.Course2, d.Advice2, d.Exam2)
	return err
}

func DeleteDisciplines2(d structures.Discipline2) error {
	_, err := Db.Exec("delete from disciplines2 where id_discipline = $1", d.IdDiscipline)
	return err
}

func UpdateDisciplines2(d structures.Discipline2) error {
	_, err := Db.Exec("update disciplines2 set disciplines=$2, hours_of_lecture=$3, hours_of_practice=$4, course=$5, "+
		"advice=$6, exam=$7, hours_of_lecture2=$8, hours_of_practice2=$9, course2=$10, advice2=$11, exam2=$12 where id_discipline = $1 ", d.IdDiscipline, d.Disciplines, d.HoursOfLecture, d.HoursOfPractice, d.Course, d.Advice, d.Exam, d.HoursOfLecture2, d.HoursOfPractice2, d.Course2, d.Advice2, d.Exam2)
	return err
}

func OutputDisciplines2() ([]structures.Discipline2, error) {
	rows, err := Db.Query("select * from public.disciplines2 order by id_discipline")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var disciplines []structures.Discipline2

	for rows.Next() {
		d := structures.Discipline2{}
		err := rows.Scan(&d.IdDiscipline, &d.Disciplines, &d.HoursOfLecture, &d.HoursOfPractice, &d.Course, &d.Advice, &d.Exam, &d.HoursOfLecture2, &d.HoursOfPractice2, &d.Course2, &d.Advice2, &d.Exam2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		disciplines = append(disciplines, d)
	}

	return disciplines, err
}

func InputSpecialities2(s structures.Speciality2) error {
	_, err := Db.Exec("insert into specialities2 (speciality, speciality_code) values ($1, $2) ", s.Speciality, s.SpecialityCode)
	return err
}

func DeleteSpecialities2(s structures.Speciality2) error {
	_, err := Db.Exec("delete from specialities2 where id_speciality = $1", s.IdSpeciality)
	return err
}

func UpdateSpecialities2(s structures.Speciality2) error {
	_, err := Db.Exec("update specialities2 set speciality = $2, speciality_code=$3 where id_speciality = $1", s.IdSpeciality, s.Speciality, s.SpecialityCode)
	return err
}

func OutputSpecialities2() ([]structures.Speciality2, error) {
	rows, err := Db.Query("select id_speciality, speciality, COALESCE(speciality_code,'')  from specialities2 order by id_speciality")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var specialities []structures.Speciality2

	for rows.Next() {
		s := structures.Speciality2{}
		err := rows.Scan(&s.IdSpeciality, &s.Speciality, &s.SpecialityCode)
		if err != nil {
			fmt.Println(err)
			continue
		}
		specialities = append(specialities, s)
	}
	return specialities, err
}

func OutputLoadList(t int, t1 int) ([]structures.LoadList, error) {
	rows, err := Db.Query("select * from public.real_load where id_teacher = $1 and year = $2", t, t1)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var loadlist []structures.LoadList

	for rows.Next() {
		ll := structures.LoadList{}
		err := rows.Scan(&ll.IdTeacher, &ll.Disciplines, &ll.GroupOld, &ll.GroupNew, &ll.Form, &ll.Amount, &ll.InAll, &ll.InWeek, &ll.HoursOfLecture, &ll.HoursOfPractice, &ll.Course, &ll.Advice, &ll.Exam, &ll.InAll2, &ll.InWeek2, &ll.HoursOfLecture2, &ll.HoursOfPractice2, &ll.Course2, &ll.Advice2, &ll.Exam2, &ll.Hours, &ll.Budget, &ll.Commerce, &ll.Total, &ll.Year)
		if err != nil {
			fmt.Println(err)
			continue
		}
		loadlist = append(loadlist, ll)
	}
	return loadlist, err
}

func InputPosts(p structures.Post) error {
	_, err := Db.Exec("insert into posts (posts) values ($1) ", p.Posts)
	return err
}

func DeletePosts(p structures.Post) error {
	_, err := Db.Exec("delete from posts where id_post = $1", p.IdPost)
	return err
}

func UpdatePosts(p structures.Post) error {
	_, err := Db.Exec("update posts set posts = $2 where id_post = $1", p.IdPost, p.Posts)
	return err
}

func OutputPosts() ([]structures.Post, error) {
	rows, err := Db.Query("select * from posts order by id_post")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var posts []structures.Post

	for rows.Next() {
		p := structures.Post{}
		err := rows.Scan(&p.IdPost, &p.Posts)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	return posts, err
}

func InputLoads(l structures.Load) error {
	if l.Year == 0 {
		l.Year = time.Now().Year()
	}
	_, err := Db.Exec("insert into load2 (id_discipline, id_group, id_teacher, year) values ($1, $2, $3, $4) ",
		l.IdDiscipline, l.IdGroup, l.IdTeacher, l.Year)
	return err
}

func DeleteLoads(l structures.Load) error {
	_, err := Db.Exec("delete from load2 where id_load = $1", l.IdLoad)
	return err
}

func UpdateLoads(l structures.Load) error {
	_, err := Db.Exec("update load2 set id_discipline = $2, id_group = $3, id_teacher = $4, "+
		"year=$5 where id_load = $1", l.IdLoad, l.IdDiscipline, l.IdGroup, l.IdTeacher, l.Year)
	return err
}

func OutputLoads() ([]structures.Load, error) {
	rows, err := Db.Query("select id_load, id_discipline, id_group, id_teacher, year from load2 order by id_load")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var loads []structures.Load

	for rows.Next() {
		l := structures.Load{}
		err := rows.Scan(&l.IdLoad, &l.IdDiscipline, &l.IdGroup, &l.IdTeacher, &l.Year)
		if err != nil {
			fmt.Println(err)
			continue
		}
		loads = append(loads, l)
	}
	return loads, err
}

func InputTeachers(t structures.Teacher) error {
	if t.DateOfHiring == time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) {
		t.DateOfHiring = time.Now()
	}
	if t.DateOfBirth == time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC) {
		t.DateOfBirth = t.DateOfHiring.AddDate(-18, 0, 0)
	}
	_, err := Db.Exec("insert into teachers (surname, name, patronymic, date_of_birth, passport_series, id_passport,"+
		" email, telephone, id_post, date_of_hiring) values ($1, $2, $3, $4, $5, $6, "+
		"$7, $8, $9, $10) ", t.Surname, t.Name, t.Patronymic, t.DateOfBirth, NewNullString(t.PassportSeries), NewNullString(t.IdPassport),
		NewNullString(t.Email), NewNullString(t.Telephone), t.IdPost, t.DateOfHiring)
	return err
}

func DeleteTeachers(t structures.Teacher) error {
	_, err := Db.Exec("delete from teachers where id_teacher = $1", t.IdTeacher)
	return err
}

func UpdateTeachers(t structures.Teacher) error {
	_, err := Db.Exec("update teachers set surname=$2, name=$3, patronymic=$4, date_of_birth=$5, "+
		" passport_series=$6, id_passport=$7, email=$8, telephone=$9, id_post=$10, date_of_hiring=$11 "+
		"where id_teacher = $1", t.IdTeacher, t.Surname, t.Name, t.Patronymic, t.DateOfBirth,
		NewNullString(t.PassportSeries), NewNullString(t.IdPassport), NewNullString(t.Email), NewNullString(t.Telephone), t.IdPost, t.DateOfHiring)
	return err
}

func OutputTeachers() ([]structures.Teacher, error) { // вывод таблицы учителей
	rows, err := Db.Query(` select id_teacher, surname, COALESCE("name", ''),COALESCE(patronymic, ''), date_of_birth, COALESCE(passport_series, ''), 
	COALESCE(id_passport, ''),  COALESCE(email,''), 
	 COALESCE(telephone, ''), id_post, date_of_hiring from teachers order by id_teacher`)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var teachers []structures.Teacher

	for rows.Next() {
		t := structures.Teacher{}
		err := rows.Scan(&t.IdTeacher, &t.Surname, &t.Name, &t.Patronymic, &t.DateOfBirth,
			&t.PassportSeries, &t.IdPassport, &t.Email, &t.Telephone, &t.IdPost, &t.DateOfHiring)
		if err != nil {
			fmt.Println(err)
			continue
		}
		teachers = append(teachers, t)
	}

	return teachers, err
}
