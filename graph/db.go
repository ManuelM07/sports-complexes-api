package graph

import (
	"context"
	"log"
	"os"

	"github.com/ManuelM07/sports-complexes-api/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

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
