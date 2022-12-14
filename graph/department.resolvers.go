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
func (r *departmentResolver) Company(ctx context.Context, obj *model.Department) (*model.Company, error) {
	return r.CompanyLoader.Load(obj.CompanyID)
}

// Employees is the resolver for the employees field.
func (r *departmentResolver) Employees(ctx context.Context, obj *model.Department) (*model.EmployeePagination, error) {
	panic(fmt.Errorf("not implemented: Employees - employees"))
}

// Department returns DepartmentResolver implementation.
func (r *Resolver) Department() DepartmentResolver { return &departmentResolver{r} }

type departmentResolver struct{ *Resolver }
