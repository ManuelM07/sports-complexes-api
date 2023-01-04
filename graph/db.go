package graph

import (
	"context"
	"log"
	"os"

	"github.com/ManuelM07/sports-complexes-api/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

// --------------- Querys ---------------
func getUser(stmt string, id string) (*model.User, error) {
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

func getUsers(stmt string) ([]*model.User, error) {
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

func getComplex(stmt string, id string) (*model.Complex, error) {
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

func getComplexs(stmt string) ([]*model.Complex, error) {
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

func getSchedule(stmt string, id string) (*model.Schedule, error) {
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

func getSchedules(stmt string) ([]*model.Schedule, error) {
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

func getScheduleComplex(stmt string, id string, available *bool) ([]*model.ScheduleComplex, error) {
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

// --------------- Mutations ---------------

func insertUser(stmt string, input model.UserInput) (*model.User, error) {
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

func insertComplex(stmt string, input model.ComplexInput) (*model.Complex, error) {
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

func insertSchedule(stmt string, input model.ScheduleInput) (*model.Schedule, error) {
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

func insertScheduleComplex(stmt string, input model.ScheduleComplexInput) (*model.ScheduleComplex, error) {
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

func insertUserComplex(stmt string, input model.UserComplexInput) (*model.UserComplex, error) {
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
