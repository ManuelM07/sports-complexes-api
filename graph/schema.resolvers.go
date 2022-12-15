package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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
	s := `INSERT INTO public.schedule(
		start, "end")
		VALUES ($1, $2) returning id, CAST(start AS string), CAST("end" AS string);`

	resp, err := insertSchedule(s, input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// CreateScheduleComplex is the resolver for the createScheduleComplex field.
func (r *mutationResolver) CreateScheduleComplex(ctx context.Context, input model.ScheduleComplexInput) (*model.ScheduleComplex, error) {
	s := `INSERT INTO public.schedule_complex(
		schedule_id, complex_id, available, limit_people, count_people)
		VALUES ($1, $2, $3, $4, $5) returning *;`

	resp, err := insertScheduleComplex(s, input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// CreateUserComplex is the resolver for the createUserComplex field.
func (r *mutationResolver) CreateUserComplex(ctx context.Context, input model.UserComplexInput) (*model.UserComplex, error) {
	s := `INSERT INTO public.user_complex(
		user_id, complex_id)
		VALUES ($1, $2) returning *;`

	resp, err := insertUserComplex(s, input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	s := `select * from public.user where public.user.id = $1;`

	resp, err := getUser(s, id)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Complex is the resolver for the complex field.
func (r *queryResolver) Complex(ctx context.Context, id string) (*model.Complex, error) {
	s := `select * from public.complex where public.complex.id = $1;`

	resp, err := getComplex(s, id)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Schedule is the resolver for the schedule field.
func (r *queryResolver) Schedule(ctx context.Context, id string) (*model.Schedule, error) {
	s := `select * from public.schedule where public.schedule.id = $1;`

	resp, err := getSchedule(s, id)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
