package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/ManuelM07/sports-complexes-api/graph/generated"
	"github.com/ManuelM07/sports-complexes-api/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	s := `INSERT INTO public.user(
			name, years, birthday, weight, height)
			VALUES ($1, $2, $3, $4, $5) returning *;`

	resp, err := insertUser(s, input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// CreateComplex is the resolver for the createComplex field.
func (r *mutationResolver) CreateComplex(ctx context.Context, input model.ComplexInput) (*model.Complex, error) {
	s := `INSERT INTO public.complex(
			name)
			VALUES ($1) returning *;`

	resp, err := insertComplex(s, input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// CreateSchedule is the resolver for the createSchedule field.
func (r *mutationResolver) CreateSchedule(ctx context.Context, input model.ScheduleInput) (*model.Schedule, error) {
	panic(fmt.Errorf("not implemented: CreateSchedule - createSchedule"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	s := `select * from public.user where public.user.id = $1;`

	//connectDB()
	resp, err := getUser(s, id)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Complex is the resolver for the complex field.
func (r *queryResolver) Complex(ctx context.Context, id string) (*model.Complex, error) {
	panic(fmt.Errorf("not implemented: Complex - complex"))
}

// Schedule is the resolver for the schedule field.
func (r *queryResolver) Schedule(ctx context.Context, id string) (*model.Schedule, error) {
	panic(fmt.Errorf("not implemented: Schedule - schedule"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
