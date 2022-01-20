package main

import (
	"flag"
	"os"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/customer"
	"github.com/deb-ict/cloudbm-community/pkg/employee"
	"github.com/deb-ict/cloudbm-community/pkg/product"
	"github.com/deb-ict/cloudbm-community/pkg/project"
	"github.com/deb-ict/cloudbm-community/pkg/storage/mongodb"
	"github.com/deb-ict/cloudbm-community/pkg/ticket"
	"github.com/deb-ict/cloudbm-community/pkg/timesheet"
	"github.com/deb-ict/cloudbm-community/pkg/user"
	"github.com/deb-ict/cloudbm-community/pkg/webhost"
	"github.com/joho/godotenv"
)

func GetConfigPath(configPath string) string {
	if len(configPath) == 0 {
		configPath = os.Getenv("CONFIG_PATH")
	}
	if len(configPath) == 0 {
		configPath = "/etc/cloudbm/webhost.yaml"
	}
	return configPath
}

func main() {
	// Parse arguments
	var configPath string
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&configPath, "config", "", "the path of the configuration file")
	flag.Parse()

	// Load the environment config and get the correct config path
	if _, err := os.Stat("webhost.env"); err == nil {
		godotenv.Load("webhost.env")
	}
	configPath = GetConfigPath(configPath)

	// Initialize the webhost
	host := webhost.NewWebHost()
	host.GetConfig().Load(configPath)

	// Setup the api
	db := mongodb.NewDatabase()

	// Initialize the service
	userService := user.NewService(db.GetUserRepository())
	employeeService := employee.NewService(db.GetEmployeeRepository())
	customerService := customer.NewService(db.GetCustomerRepository())
	productService := product.NewService(db.GetProductRepository())
	projectService := project.NewService(db.GetProjectRepository())
	ticketService := ticket.NewService(db.GetTicketRepository())
	timesheetService := timesheet.NewService(db.GetTimesheetRepository())

	if host.GetConfig().IsApiEnabled() {
		host.AddApiHandler(user.NewApi(userService))
		host.AddApiHandler(employee.NewApi(employeeService))
		host.AddApiHandler(customer.NewApi(customerService))
		host.AddApiHandler(product.NewApi(productService))
		host.AddApiHandler(project.NewApi(projectService))
		host.AddApiHandler(ticket.NewApi(ticketService))
		host.AddApiHandler(timesheet.NewApi(timesheetService))
	}

	// Run the host
	host.Run()

	// Exit
	os.Exit(0)
}
