// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Complex struct {
	ID   *int   `json:"id"`
	Name string `json:"name"`
}

type ComplexInput struct {
	Name string `json:"name"`
}

type Schedule struct {
	ID    *int      `json:"id"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type ScheduleComplex struct {
	ID          string `json:"id"`
	ScheduleID  string `json:"schedule_id"`
	ComplexID   string `json:"complex_id"`
	Available   *bool  `json:"available"`
	LimitPeople *int   `json:"limit_people"`
	CountPeople *int   `json:"count_people"`
}

type ScheduleInput struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
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
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ComplexID string `json:"complex_id"`
}

type UserInput struct {
	Name     string   `json:"name"`
	Years    *int     `json:"years"`
	Birthday *string  `json:"birthday"`
	Weight   *float64 `json:"weight"`
	Height   *int     `json:"height"`
}
