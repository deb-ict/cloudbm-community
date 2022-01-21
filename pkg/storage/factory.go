package storage

import (
	"github.com/deb-ict/cloudbm-community/pkg/module/customer"
	"github.com/deb-ict/cloudbm-community/pkg/module/employee"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/project"
	"github.com/deb-ict/cloudbm-community/pkg/module/ticket"
	"github.com/deb-ict/cloudbm-community/pkg/module/timesheet"
	"github.com/deb-ict/cloudbm-community/pkg/module/user"
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
