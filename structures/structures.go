package structures

import (
	"time"
)

type Plan struct {
	Fio         string
	Disciplines string
	Hours       int
}

type Itogo struct {
	Disciplines string
	Fio         string
	Plan        int
	September   int
	October     int
	November    int
	December    int
	Give        int
	Remain      int
}

type Itogo2 struct {
	Disciplines string
	Fio         string
	Plan        int
	January     int
	February    int
	March       int
	April       int
	May         int
	June        int
	Give        int
	Remain      int
}

type Passwords2 struct {
	Login     string `json:"Login"`
	Password  string `json:"Password"`
	IdTeacher int    `json:"IdTeacher"`
	Fio       string
	Admin     bool
}

type Passwords struct {
	IdPassword int    `json:"IdPassword""`
	Login      string `json:"Login"`
	Password   string `json:"Password"`
	IdTeacher  int    `json:"IdTeacher"`
}

type Timetable2 struct {
	IdTimetable     int
	DateOfTimetable time.Time
	IdLoad          int
	Hours           int
	TypeOfLesson    bool
	Day             int
	Month           int
	Year            int
}

type TimetableGroups struct {
	IdGroup     int    `json:"IdGroup"`
	Month       int    `json:"Month"`
	Disciplines string `json:"Disciplines"`
	Fio         string `json:"Fio"`
	AllInHours  int    `json:"AllInHours"`
	RemainHours int    `json:"RemainHours"`
	TotalHours  int    `json:"TotalHours"`
	Year        int    `json:"Year"`
	IdLoad      int
	Type        bool
	GroupOld    string
	DayNHour    []DayNHour
}

type DayNHour struct {
	Day   int
	Hours int
}

type Group2 struct {
	IdGroup      int    `json:"IdGroup"`
	GroupOld     string `json:"GroupOld"`
	GroupNew     string
	Form         bool
	Amount       int
	IdSpeciality int
}

type Speciality2 struct {
	IdSpeciality   int    `json:"IdSpeciality"`
	Speciality     string `json:"Speciality"`
	SpecialityCode string `json:"SpecialityCode"`
}

type Discipline2 struct {
	IdDiscipline     int    `json:"IdDiscipline"`
	Disciplines      string `json:"Disciplines"`
	HoursOfLecture   int
	HoursOfPractice  int
	Course           int
	Advice           int
	Exam             int
	HoursOfLecture2  int
	HoursOfPractice2 int
	Course2          int
	Advice2          int
	Exam2            int
}

type LoadList struct {
	IdTeacher        int
	Disciplines      string
	GroupOld         string
	GroupNew         string
	Form             bool
	Amount           int
	InAll            int
	InWeek           int
	HoursOfLecture   int
	HoursOfPractice  int
	Course           int
	Advice           int
	Exam             int
	InAll2           int
	InWeek2          int
	HoursOfLecture2  int
	HoursOfPractice2 int
	Course2          int
	Advice2          int
	Exam2            int
	Hours            int
	Budget           int
	Commerce         int
	Total            int
	Year             int
}

type Teacher struct {
	IdTeacher      int    `json:"IdTeacher"`
	Surname        string `json:"Surname"`
	Name           string `json:"Name"`
	Patronymic     string `json:"Patronymic"`
	DateOfBirth    time.Time
	PassportSeries string
	IdPassport     string
	Email          string
	Telephone      string
	IdPost         int
	DateOfHiring   time.Time
}

type Post struct {
	IdPost int    `json:"IdPost"`
	Posts  string `json:"Posts"`
}

type Load struct {
	IdLoad       int   `json:"IdLoad"`
	IdDiscipline int   `json:"IdDiscipline"`
	IdGroup      int   `json:"IdGroup"`
	IdTeacher    int   `json:"IdTeacher"`
	Year         int   `json:"Year"`
	Yearss       []int `json:"Yearss"`
}
