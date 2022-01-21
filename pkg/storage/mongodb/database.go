package mongodb

import (
	"context"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/module/customer"
	"github.com/deb-ict/cloudbm-community/pkg/module/employee"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/project"
	"github.com/deb-ict/cloudbm-community/pkg/module/ticket"
	"github.com/deb-ict/cloudbm-community/pkg/module/timesheet"
	"github.com/deb-ict/cloudbm-community/pkg/module/user"
	"github.com/deb-ict/cloudbm-community/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
	storage.RepositoryFactory
	GetConfig() Config
	Connect() error
	Disconnect() error
}

type database struct {
	config   Config
	client   *mongo.Client
	database *mongo.Database
}

func NewDatabase() Database {
	return &database{
		config: NewConfig(),
	}
}

func (db *database) GetConfig() Config {
	return db.config
}

func (db *database) Connect() error {
	var err error
	client, err := mongo.NewClient(options.Client().ApplyURI(db.config.GetUri()))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	db.client = client
	db.database = client.Database(db.config.GetDatabase())

	return nil
}

func (db *database) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := db.client.Disconnect(ctx)
	if err != nil {
		return err
	}

	db.database = nil
	db.client = nil

	return nil
}

func (db *database) GetUserRepository() user.Repository {
	return &userRepository{
		db: *db,
	}
}

func (db *database) GetEmployeeRepository() employee.Repository {
	return &employeeRepository{
		db: *db,
	}
}

func (db *database) GetCustomerRepository() customer.Repository {
	return &customerRepository{
		db: *db,
	}
}

func (db *database) GetProjectRepository() project.Repository {
	return &projectRepository{
		db: *db,
	}
}

func (db *database) GetProductRepository() product.Repository {
	return &productRepository{
		db: *db,
	}
}

func (db *database) GetTicketRepository() ticket.Repository {
	return &ticketRepository{
		db: *db,
	}
}

func (db *database) GetTimesheetRepository() timesheet.Repository {
	return &timesheetRepository{
		db: *db,
	}
}

func (db *database) getNormalizedPaging(pageIndex int, pageSize int) (int, int) {
	return db.getNormalizedPageIndex(pageIndex), db.getNormalizedPageSize(pageSize)
}

func (db *database) getNormalizedPageIndex(pageIndex int) int {
	if pageIndex < 1 {
		return 1
	}
	return pageIndex
}

func (db *database) getNormalizedPageSize(pageSize int) int {
	if pageSize > db.GetConfig().GetMaxPageSize() {
		return db.GetConfig().GetMaxPageSize()
	}
	if pageSize < 1 {
		return 1
	}
	return pageSize
}
