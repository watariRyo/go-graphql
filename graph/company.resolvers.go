package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/watariRyo/go-graphql/graph/model"
)

// Departments is the resolver for the departments field.
func (r *companyResolver) Departments(ctx context.Context, obj *model.Company) (*model.DepartmentPagination, error) {
	nodes := []*model.Department{
		{
			ID:             "depart1",
			DepartmentName: "部署1",
			Email:          "test1@mail.com",
			CompanyID:      obj.ID,
		},
		{
			ID:             "depart2",
			DepartmentName: "部署2",
			Email:          "test2@mail.com",
			CompanyID:      obj.ID,
		},
	}
	return &model.DepartmentPagination{
		PageInfo: &model.PaginationInfo{
			Page:             1,
			PaginationLength: 1,
			HasNextPage:      false,
			Count:            2,
			TotalCount:       2,
		},
		Nodes: nodes,
	}, nil
}

// Employees is the resolver for the employees field.
func (r *companyResolver) Employees(ctx context.Context, obj *model.Company) (*model.EmployeePagination, error) {
	panic(fmt.Errorf("not implemented: Employees - employees"))
}

// Company returns CompanyResolver implementation.
func (r *Resolver) Company() CompanyResolver { return &companyResolver{r} }

type companyResolver struct{ *Resolver }
