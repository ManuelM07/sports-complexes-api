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
	resp, err := insertUser(input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// CreateComplex is the resolver for the createComplex field.
func (r *mutationResolver) CreateComplex(ctx context.Context, input model.ComplexInput) (*model.Complex, error) {
	resp, err := insertComplex(input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// CreateSchedule is the resolver for the createSchedule field.
func (r *mutationResolver) CreateSchedule(ctx context.Context, input model.ScheduleInput) (*model.Schedule, error) {
	resp, err := insertSchedule(input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// CreateScheduleComplex is the resolver for the createScheduleComplex field.
func (r *mutationResolver) CreateScheduleComplex(ctx context.Context, input model.ScheduleComplexInput) (*model.ScheduleComplex, error) {
	resp, err := insertScheduleComplex(input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// CreateUserComplex is the resolver for the createUserComplex field.
func (r *mutationResolver) CreateUserComplex(ctx context.Context, input model.UserComplexInput) (*model.UserComplex, error) {
	resp, err := insertUserComplex(input)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// UpdateComplex is the resolver for the updateComplex field.
func (r *mutationResolver) UpdateComplex(ctx context.Context, input model.ComplexInput) (*model.Complex, error) {
	panic(fmt.Errorf("not implemented: UpdateComplex - updateComplex"))
}

// DeleteSchedule is the resolver for the deleteSchedule field.
func (r *mutationResolver) DeleteSchedule(ctx context.Context, input model.ScheduleInput) (*model.Schedule, error) {
	panic(fmt.Errorf("not implemented: DeleteSchedule - deleteSchedule"))
}

// UpdateScheduleComplex is the resolver for the updateScheduleComplex field.
func (r *mutationResolver) UpdateScheduleComplex(ctx context.Context, input model.ScheduleComplexInput) (*model.ScheduleComplex, error) {
	panic(fmt.Errorf("not implemented: UpdateScheduleComplex - updateScheduleComplex"))
}

// UpdateUserComplex is the resolver for the updateUserComplex field.
func (r *mutationResolver) UpdateUserComplex(ctx context.Context, input model.UserComplexInput) (*model.UserComplex, error) {
	panic(fmt.Errorf("not implemented: UpdateUserComplex - updateUserComplex"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	resp, err := getUser(id)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Complex is the resolver for the complex field.
func (r *queryResolver) Complex(ctx context.Context, id string) (*model.Complex, error) {
	resp, err := getComplex(id)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Schedule is the resolver for the schedule field.
func (r *queryResolver) Schedule(ctx context.Context, id string) (*model.Schedule, error) {
	resp, err := getSchedule(id)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	resp, err := getUsers()
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Complexs is the resolver for the complexs field.
func (r *queryResolver) Complexs(ctx context.Context) ([]*model.Complex, error) {
	resp, err := getComplexs()
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// Schedules is the resolver for the schedules field.
func (r *queryResolver) Schedules(ctx context.Context) ([]*model.Schedule, error) {
	resp, err := getSchedules()
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// ScheduleComplex is the resolver for the scheduleComplex field.
func (r *queryResolver) ScheduleComplex(ctx context.Context, complexID string, available *bool) ([]*model.ScheduleComplex, error) {
	resp, err := getScheduleComplex(complexID, available)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// UserComplexToUser is the resolver for the userComplexToUser field.
func (r *queryResolver) UserComplexToUser(ctx context.Context, userID string) ([]*model.UserComplex, error) {
	resp, err := getUserComplexToUser(userID)
	if err != nil {

		log.Fatal(err)
	}
	return resp, nil
}

// UserComplexToComplex is the resolver for the userComplexToComplex field.
func (r *queryResolver) UserComplexToComplex(ctx context.Context, complexID string) ([]*model.UserComplex, error) {
	resp, err := getUserComplexToComplex(complexID)
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
