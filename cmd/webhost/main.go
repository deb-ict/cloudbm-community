package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/http/webhost"
	"github.com/deb-ict/cloudbm-community/pkg/module/customer"
	"github.com/deb-ict/cloudbm-community/pkg/module/employee"
	"github.com/deb-ict/cloudbm-community/pkg/module/product"
	"github.com/deb-ict/cloudbm-community/pkg/module/project"
	"github.com/deb-ict/cloudbm-community/pkg/module/ticket"
	"github.com/deb-ict/cloudbm-community/pkg/module/timesheet"
	"github.com/deb-ict/cloudbm-community/pkg/module/user"
	"github.com/deb-ict/cloudbm-community/pkg/storage/mongodb"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Menu struct {
	Name  string
	Href  string
	Items []MenuItem
}

type MenuItem struct {
}

type Page struct {
	Title       string
	ShowSideBar bool
	Model       interface{}
}

func GetConfigPath(configPath string) string {
	if len(configPath) == 0 {
		configPath = os.Getenv("CONFIG_PATH")
	}
	if len(configPath) == 0 {
		configPath = "/etc/cloudbm/webhost.yaml"
	}
	return configPath
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/404.html",
	))
	template.Execute(w, Page{
		Title:       "Page not found",
		ShowSideBar: false,
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/home.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
	})
}

// ADMIN - CRM - COMPANY
func adminCrmCompanyOverview_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/admin/crm/company/overview.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
	})
}
func adminCrmCompanyDetail_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/admin/crm/company/detail.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
	})
}
func adminCrmCompanyCreate_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/admin/crm/company/create.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
	})
}
func adminCrmCompanyUpdate_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/admin/crm/company/edit.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
	})
}

// ADMIN - CRM - CONTACT

func adminCrmContactOverview_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/admin/crm/contact/overview.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
	})
}
func adminCrmContactDetail_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/admin/crm/contact/detail.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
	})
}
func adminCrmContactCreate_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/admin/crm/contact/create.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
	})
}
func adminCrmContactUpdate_GetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/admin/crm/contact/overview.html",
	))
	template.Execute(w, Page{
		Title:       "Home",
		ShowSideBar: true,
		Model:       id,
	})
}

func main() {
	var err error

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

	// Setup the router
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/home", homeHandler).Methods("GET")
	router.HandleFunc("/index", homeHandler).Methods("GET")

	// Public - ServiceDesk
	router.HandleFunc("/servicedesk", homeHandler).Methods("GET")
	router.HandleFunc("/servicedesk/tickets", homeHandler).Methods("GET")
	router.HandleFunc("/servicedesk/tickets/{id}", homeHandler).Methods("GET")
	router.HandleFunc("/servicedesk/tickets/submit", homeHandler).Methods("GET")

	// Agent - ServiceDesk
	router.HandleFunc("/servicedesk/agent/dashboard", homeHandler).Methods("GET")
	router.HandleFunc("/servicedesk/agent/tickets", homeHandler).Methods("GET")
	router.HandleFunc("/servicedesk/agent/tickets/{id}", homeHandler).Methods("GET")

	// Admin - HrEmployees
	router.HandleFunc("/admin/hr/employees", homeHandler).Methods("GET")
	router.HandleFunc("/admin/hr/employees/{id}", homeHandler).Methods("GET")
	router.HandleFunc("/admin/hr/employees/create", homeHandler).Methods("GET")
	router.HandleFunc("/admin/hr/employees/edit/{id}", homeHandler).Methods("GET")

	// Admin - CrmCompanies
	router.HandleFunc("/admin/crm/companies", adminCrmCompanyOverview_GetHandler).Methods("GET")
	router.HandleFunc("/admin/crm/companies/{id}", adminCrmCompanyDetail_GetHandler).Methods("GET")
	router.HandleFunc("/admin/crm/companies/create", adminCrmCompanyCreate_GetHandler).Methods("GET")
	router.HandleFunc("/admin/crm/companies/edit/{id}", adminCrmCompanyUpdate_GetHandler).Methods("GET")

	// Admin - CrmContacts
	router.HandleFunc("/admin/crm/contacts", adminCrmContactOverview_GetHandler).Methods("GET")
	router.HandleFunc("/admin/crm/contacts/{id}", adminCrmContactDetail_GetHandler).Methods("GET")
	router.HandleFunc("/admin/crm/contacts/create", adminCrmContactCreate_GetHandler).Methods("GET")
	router.HandleFunc("/admin/crm/contacts/edit/{id}", adminCrmContactUpdate_GetHandler).Methods("GET")

	// Admin - ServiceDesk (metadata)
	router.HandleFunc("/admin/servicedesk/types", homeHandler).Methods("GET")
	router.HandleFunc("/admin/servicedesk/statuses", homeHandler).Methods("GET")
	router.HandleFunc("/admin/servicedesk/priorities", homeHandler).Methods("GET")

	// Serve the static files
	fs := http.FileServer(http.Dir("./web/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Initialize the webhost
	host := webhost.NewWebHost(router)
	err = host.GetConfig().Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Setup the api
	db := mongodb.NewDatabase()
	err = db.GetConfig().Load(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}

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
