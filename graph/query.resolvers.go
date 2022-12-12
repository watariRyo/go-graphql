package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/watariRyo/go-graphql/graph/model"
)

// Company is the resolver for the company field.
func (r *queryResolver) Company(ctx context.Context, id string) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: Company - company"))
}

// Companies is the resolver for the companies field.
func (r *queryResolver) Companies(ctx context.Context, limit int, offset *int) (*model.CompanyPagination, error) {
	panic(fmt.Errorf("not implemented: Companies - companies"))
}

// Department is the resolver for the department field.
func (r *queryResolver) Department(ctx context.Context, id string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented: Department - department"))
}

// Departments is the resolver for the departments field.
func (r *queryResolver) Departments(ctx context.Context, limit int, offset int) (*model.DepartmentPagination, error) {
	panic(fmt.Errorf("not implemented: Departments - departments"))
}

// Employee is the resolver for the employee field.
func (r *queryResolver) Employee(ctx context.Context, id *int) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented: Employee - employee"))
}

// Employees is the resolver for the employees field.
func (r *queryResolver) Employees(ctx context.Context, limit int, offset *int, email *string, gender *model.Gender, isManager *bool, hasDependent *bool) (*model.EmployeePagination, error) {
	panic(fmt.Errorf("not implemented: Employees - employees"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }