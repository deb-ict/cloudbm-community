package storage

import (
	"github.com/deb-ict/cloudbm-community/pkg/customer"
	"github.com/deb-ict/cloudbm-community/pkg/employee"
	"github.com/deb-ict/cloudbm-community/pkg/product"
	"github.com/deb-ict/cloudbm-community/pkg/project"
	"github.com/deb-ict/cloudbm-community/pkg/ticket"
	"github.com/deb-ict/cloudbm-community/pkg/timesheet"
	"github.com/deb-ict/cloudbm-community/pkg/user"
)

type RepositoryFactory interface {
	GetUserRepository() user.Repository
	GetEmployeeRepository() employee.Repository
	GetCustomerRepository() customer.Repository
	GetProjectRepository() project.Repository
	GetProductRepository() product.Repository
	GetTicketRepository() ticket.Repository
	GetTimesheetRepository() timesheet.Repository
}
