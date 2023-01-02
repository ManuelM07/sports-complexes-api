// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Complex struct {
	ID   *int   `json:"id"`
	Name string `json:"name"`
}

type ComplexInput struct {
	Name string `json:"name"`
}

type Schedule struct {
	ID    *int    `json:"id"`
	Start *string `json:"start"`
	End   *string `json:"end"`
}

type ScheduleComplex struct {
	ID          *int  `json:"id"`
	ScheduleID  int   `json:"schedule_id"`
	ComplexID   int   `json:"complex_id"`
	Available   *bool `json:"available"`
	LimitPeople *int  `json:"limit_people"`
	CountPeople *int  `json:"count_people"`
}

type ScheduleComplexInput struct {
	ScheduleID  int   `json:"schedule_id"`
	ComplexID   int   `json:"complex_id"`
	Available   *bool `json:"available"`
	LimitPeople *int  `json:"limit_people"`
	CountPeople *int  `json:"count_people"`
}

type ScheduleInput struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type User struct {
	ID       *int     `json:"id"`
	Name     string   `json:"name"`
	Years    int      `json:"years"`
	Birthday *string  `json:"birthday"`
	Weight   *float64 `json:"weight"`
	Height   *int     `json:"height"`
}

type UserComplex struct {
	ID        *int `json:"id"`
	UserID    int  `json:"user_id"`
	ComplexID int  `json:"complex_id"`
}

type UserComplexInput struct {
	UserID    int `json:"user_id"`
	ComplexID int `json:"complex_id"`
}

type UserInput struct {
	Name     string   `json:"name"`
	Years    *int     `json:"years"`
	Birthday *string  `json:"birthday"`
	Weight   *float64 `json:"weight"`
	Height   *int     `json:"height"`
}
