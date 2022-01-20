package mongodb

import (
	"github.com/deb-ict/cloudbm-community/pkg/customer"
	"github.com/deb-ict/cloudbm-community/pkg/employee"
	"github.com/deb-ict/cloudbm-community/pkg/product"
	"github.com/deb-ict/cloudbm-community/pkg/project"
	"github.com/deb-ict/cloudbm-community/pkg/storage"
	"github.com/deb-ict/cloudbm-community/pkg/ticket"
	"github.com/deb-ict/cloudbm-community/pkg/timesheet"
	"github.com/deb-ict/cloudbm-community/pkg/user"
)

type database struct {
}

func NewDatabase() storage.RepositoryFactory {
	return &database{}
}

func (db *database) GetUserRepository() user.Repository {
	return &userRepository{}
}

func (db *database) GetEmployeeRepository() employee.Repository {
	return &employeeRepository{}
}

func (db *database) GetCustomerRepository() customer.Repository {
	return &customerRepository{}
}

func (db *database) GetProjectRepository() project.Repository {
	return &projectRepository{}
}

func (db *database) GetProductRepository() product.Repository {
	return &productRepository{}
}

func (db *database) GetTicketRepository() ticket.Repository {
	return &ticketRepository{}
}

func (db *database) GetTimesheetRepository() timesheet.Repository {
	return &timesheetRepository{}
}
