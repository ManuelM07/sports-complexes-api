package graph

// This file containt all querys and mutations and execute the connection to database in postgresql

import (
	"context"
	"log"
	"os"

	"github.com/ManuelM07/sports-complexes-api/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

// --------------- Querys ---------------
// getUser is the query to show a user
func getUser(id string) (*model.User, error) {
	stmt := `SELECT * FROM public.user WHERE public.user.id = $1;`

	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	var user *model.User

	rows, err := dbpool.Query(context.Background(), stmt, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user = new(model.User)
		if err := rows.Scan(&user.ID, &user.Name, &user.Years, &user.Birthday, &user.Weight, &user.Height); err != nil {
			log.Fatal(err)
		}
	}

	return user, nil
}

// getUsers is the query to show a list of users
func getUsers() ([]*model.User, error) {
	stmt := `SELECT * FROM public.user;`

	users := make([]*model.User, 0)
	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	var user *model.User

	rows, err := dbpool.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user = new(model.User)
		if err := rows.Scan(&user.ID, &user.Name, &user.Years, &user.Birthday, &user.Weight, &user.Height); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

// getComplex is the query to show a complex
func getComplex(id string) (*model.Complex, error) {
	stmt := `SELECT * FROM public.complex WHERE public.complex.id = $1;`

	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	complex := new(model.Complex)
	err = dbpool.QueryRow(context.Background(), stmt, id).Scan(&complex.ID, &complex.Name)
	if err != nil {
		return nil, err
	}

	return complex, nil
}

// getComplexs is the query to show a list of complexes
func getComplexs() ([]*model.Complex, error) {
	stmt := `SELECT * FROM public.complex;`

	complexs := make([]*model.Complex, 0)
	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	var complex *model.Complex

	rows, err := dbpool.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		complex = new(model.Complex)
		if err := rows.Scan(&complex.ID, &complex.Name); err != nil {
			log.Fatal(err)
		}
		complexs = append(complexs, complex)
	}

	return complexs, nil
}

// getSchedule is the query to show a schedule
func getSchedule(id string) (*model.Schedule, error) {
	stmt := `SELECT id, CAST(start AS string), CAST("end" AS string) FROM public.schedule where public.schedule.id = $1;`

	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	schedule := new(model.Schedule)
	err = dbpool.QueryRow(context.Background(), stmt, id).Scan(&schedule.ID, &schedule.Start, &schedule.End)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// getSchedules is the query to show a list of schedules
func getSchedules() ([]*model.Schedule, error) {
	stmt := `SELECT id, CAST(start AS string), CAST("end" AS string) FROM public.schedule;`

	schedules := make([]*model.Schedule, 0)
	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	var schedule *model.Schedule

	rows, err := dbpool.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		schedule = new(model.Schedule)
		if err := rows.Scan(&schedule.ID, &schedule.Start, &schedule.End); err != nil {
			log.Fatal(err)
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

// getScheduleComplex is the query to show a list of schedule for a complex
func getScheduleComplex(id string, available *bool) ([]*model.ScheduleComplex, error) {
	stmt := `SELECT p.id, p.schedule_id, p.complex_id, p.available, p.limit_people, p.count_people, 
	s.id, CAST(s.start AS string), CAST(s."end" AS string) 
	FROM public.schedule_complex p
	JOIN schedule s
		ON s.id = p.schedule_id AND p.complex_id = $1;`

	scheduleComplexs := make([]*model.ScheduleComplex, 0)
	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	var scheduleComplex *model.ScheduleComplex
	var schedule *model.Schedule

	rows, err := dbpool.Query(context.Background(), stmt, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		scheduleComplex = new(model.ScheduleComplex)
		schedule = new(model.Schedule)
		if err := rows.Scan(&scheduleComplex.ID, &scheduleComplex.ScheduleID, &scheduleComplex.ComplexID,
			&scheduleComplex.Available, &scheduleComplex.LimitPeople, &scheduleComplex.CountPeople,
			&schedule.ID, &schedule.Start, &schedule.End); err != nil {
			log.Fatal(err)
		}
		scheduleComplex.Schedule = schedule
		scheduleComplexs = append(scheduleComplexs, scheduleComplex)
	}

	return scheduleComplexs, nil
}

// getUserComplexToUser is the query to show a list of complexes for a user
func getUserComplexToUser(id string) ([]*model.UserComplex, error) {
	stmt := `SELECT p.id, p.user_id, p.complex_id, c.id, c.name
	FROM public.user_complex p
	JOIN public.complex c
		ON p.complex_id = c.id AND p.user_id = $1;
	`

	userComplexs := make([]*model.UserComplex, 0)
	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	var userComplex *model.UserComplex
	var complex *model.Complex

	rows, err := dbpool.Query(context.Background(), stmt, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		userComplex = new(model.UserComplex)
		complex = new(model.Complex)
		if err := rows.Scan(&userComplex.ID, &userComplex.UserID, &userComplex.ComplexID,
			&complex.ID, &complex.Name); err != nil {
			log.Fatal(err)
		}
		userComplex.Complexes = complex
		userComplexs = append(userComplexs, userComplex)
	}

	return userComplexs, nil
}

// getUserComplexToUser is the query to show a list of users for a complex
func getUserComplexToComplex(id string) ([]*model.UserComplex, error) {
	stmt := `SELECT p.id, p.user_id, p.complex_id, u.id, u.name, u.years, u.birthday, u.weight, u.height
	FROM public.user_complex p
	JOIN public.user u
		ON p.user_id = u.id AND p.complex_id = $1;
	`

	userComplexs := make([]*model.UserComplex, 0)
	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	var userComplex *model.UserComplex
	var user *model.User

	rows, err := dbpool.Query(context.Background(), stmt, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		userComplex = new(model.UserComplex)
		user = new(model.User)
		if err := rows.Scan(&userComplex.ID, &userComplex.UserID, &userComplex.ComplexID,
			&user.ID, &user.Name, &user.Years, &user.Birthday, &user.Weight, &user.Height); err != nil {
			log.Fatal(err)
		}
		userComplex.Users = user
		userComplexs = append(userComplexs, userComplex)
	}

	return userComplexs, nil
}

// --------------- Mutations ---------------
// insertUser is the mutation to create a new user
func insertUser(input model.UserInput) (*model.User, error) {
	stmt := `INSERT INTO public.user(
		name, years, birthday, weight, height)
		VALUES ($1, $2, $3, $4, $5) RETURNING *;`

	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	user := new(model.User)
	err = dbpool.QueryRow(context.Background(), stmt, input.Name, *input.Years, *input.Birthday, *input.Weight, *input.Height).Scan(&user.ID, &user.Name, &user.Years, &user.Birthday, &user.Weight, &user.Height)
	if err != nil {
		return nil, err
	}

	return user, nil

}

// insertComplex is the mutation to create a new complex
func insertComplex(input model.ComplexInput) (*model.Complex, error) {
	stmt := `INSERT INTO public.complex(name)
		VALUES ($1) RETURNING *;`

	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	complex := new(model.Complex)
	err = dbpool.QueryRow(context.Background(), stmt, input.Name).Scan(&complex.ID, &complex.Name)
	if err != nil {
		return nil, err
	}

	return complex, nil
}

// insertSchedule is the mutation to create a new schedule
func insertSchedule(input model.ScheduleInput) (*model.Schedule, error) {
	stmt := `INSERT INTO public.schedule(
		start, "end")
		VALUES ($1, $2) returning id, CAST(start AS string), CAST("end" AS string);`

	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	schedule := new(model.Schedule)
	err = dbpool.QueryRow(context.Background(), stmt, input.Start, input.End).Scan(&schedule.ID, &schedule.Start, &schedule.End)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// insertScheduleComplex is the mutation to create a new scheduleComplex
func insertScheduleComplex(input model.ScheduleComplexInput) (*model.ScheduleComplex, error) {
	stmt := `INSERT INTO public.schedule_complex(
		schedule_id, complex_id, available, limit_people, count_people)
		VALUES ($1, $2, $3, $4, $5) RETURNING *;`

	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	scheduleComplex := new(model.ScheduleComplex)
	err = dbpool.QueryRow(context.Background(), stmt, input.ScheduleID, input.ComplexID, input.Available, input.LimitPeople,
		input.CountPeople).Scan(&scheduleComplex.ID, &scheduleComplex.ScheduleID, &scheduleComplex.ComplexID,
		&scheduleComplex.Available, &scheduleComplex.LimitPeople, &scheduleComplex.CountPeople)
	if err != nil {
		return nil, err
	}

	return scheduleComplex, nil
}

// insertUserComplex is the mutation to create a new userComplex
func insertUserComplex(input model.UserComplexInput) (*model.UserComplex, error) {
	stmt := `INSERT INTO public.user_complex(
		user_id, complex_id)
		VALUES ($1, $2) RETURNING *;`

	dbpool, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	userComplex := new(model.UserComplex)
	err = dbpool.QueryRow(context.Background(), stmt, input.UserID, input.ComplexID).Scan(&userComplex.ID, &userComplex.UserID, &userComplex.ComplexID)
	if err != nil {
		return nil, err
	}

	return userComplex, nil
}

// connection to db
func connectDB() (*pgxpool.Pool, error) {
	// Set connection pool configuration, with maximum connection pool size.
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	// Create a connection pool to database.
	dbpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return dbpool, nil
}
