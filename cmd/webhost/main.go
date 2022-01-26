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
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

type Menu struct {
	Name   string
	HRef   string
	Active bool
	Items  []MenuItem
}

type MenuItem struct {
	Name   string
	HRef   string
	Active bool
}

type Page struct {
	Title       string
	ShowSideBar bool
	NavMenu     []Menu
	Model       interface{}
	Username    string
}

//var sessionStore = sessions.NewCookieStore([]byte("secret-cloudbm-key"))
var sessionStore = sessions.NewFilesystemStore("", []byte("secret-cloudbm-key"))

func createMenuItem(name string, href string, activeMenu string, items []MenuItem) Menu {
	active := false
	if name == activeMenu {
		active = true
	}
	return Menu{
		Name:   name,
		HRef:   href,
		Active: active,
		Items:  items,
	}
}
func createSubMenuItem(name string, href string, activeMenuItem string) MenuItem {
	active := false
	if name == activeMenuItem {
		active = true
	}
	return MenuItem{
		Name:   name,
		HRef:   href,
		Active: active,
	}
}

func loadNavMenu(activeMenu string, activeMenuItem string) []Menu {
	menu := []Menu{
		createMenuItem("Home", "/", activeMenu, []MenuItem{}),
		createMenuItem("CRM", "", activeMenu, []MenuItem{
			createSubMenuItem("Companies", "/portal/crm/companies", activeMenuItem),
			createSubMenuItem("Contacts", "/portal/crm/contacts", activeMenuItem),
			createSubMenuItem("Deals", "/portal/crm/deals", activeMenuItem),
		}),
		createMenuItem("Sales", "/", activeMenu, []MenuItem{
			createSubMenuItem("Invoices", "/portal/sales/invoices", activeMenuItem),
		}),
		createMenuItem("HR", "/", activeMenu, []MenuItem{
			createSubMenuItem("Employees", "/portal/hr/employees", activeMenuItem),
		}),
		createMenuItem("ServiceDesk", "/", activeMenu, []MenuItem{
			createSubMenuItem("Tickets", "/servicedesk/agent", activeMenuItem),
		}),
	}

	return menu
}

func NewPage(title string, navMenu []Menu, r *http.Request) Page {
	p := Page{
		Title:   title,
		NavMenu: navMenu,
	}

	session, _ := sessionStore.Get(r, "session-name")
	session_value := session.Values["username"]
	if session_value != nil {
		p.Username = session_value.(string)
	}
	log.Printf("Session value: %v\n", session_value)

	return p
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
	navMenu := loadNavMenu("", "")
	template.Execute(w, NewPage("Page not found", navMenu, r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/home.html",
	))
	navMenu := loadNavMenu("Home", "")
	template.Execute(w, NewPage("Home", navMenu, r))
}

type LoginFormData struct {
	ValidationMessages map[string]string
	Username           string
	Password           string
}

func (data *LoginFormData) Validate() bool {
	data.ValidationMessages = make(map[string]string)
	if len(data.Username) == 0 {
		data.ValidationMessages["Username"] = "Username cannot be empty"
	}
	if len(data.Password) == 0 {
		data.ValidationMessages["Password"] = "Password cannot be empty"
	}
	return len(data.ValidationMessages) == 0
}

func accountLogin_View(w http.ResponseWriter, r *http.Request, d LoginFormData) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/account/login.html",
	))
	navMenu := loadNavMenu("", "")
	template.Execute(w, Page{
		Title:       "Login",
		NavMenu:     navMenu,
		ShowSideBar: false,
		Model:       d,
	})
}
func accountLogin_GetHandler(w http.ResponseWriter, r *http.Request) {
	accountLogin_View(w, r, LoginFormData{})
}
func accountLogin_PostHandler(w http.ResponseWriter, r *http.Request) {
	returnUrl := r.URL.Query().Get("returnUrl")
	if len(returnUrl) == 0 {
		returnUrl = "/"
	}

	formData := LoginFormData{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}
	if formData.Validate() {
		if formData.Password == "secure" {
			//TODO: Set the session
			session, _ := sessionStore.Get(r, "session-name")
			session.Values["username"] = formData.Username
			session.Options = &sessions.Options{
				MaxAge:   86400, //24h (24 * 60 * 60)
				Secure:   true,
				SameSite: http.SameSiteNoneMode,
				Path:     "/",
				//Domain:   "localhost",
			}
			session.Save(r, w)

			http.Redirect(w, r, returnUrl, http.StatusMovedPermanently)
			return
		}
	}

	accountLogin_View(w, r, formData)
}

func accountLogout_PostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "session-name")
	for k := range session.Values {
		delete(session.Values, k)
	}
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// ADMIN - CRM - COMPANY
func adminCrmCompanyOverview_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/portal/crm/company/overview.html",
	))
	navMenu := loadNavMenu("CRM", "Companies")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}
func adminCrmCompanyDetail_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/portal/crm/company/detail.html",
	))
	navMenu := loadNavMenu("CRM", "Companies")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}
func adminCrmCompanyCreate_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/portal/crm/company/create.html",
	))
	navMenu := loadNavMenu("CRM", "Companies")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}
func adminCrmCompanyUpdate_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/portal/crm/company/edit.html",
	))
	navMenu := loadNavMenu("CRM", "Companies")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}

// ADMIN - CRM - CONTACT

func adminCrmContactOverview_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/portal/crm/contact/overview.html",
	))
	navMenu := loadNavMenu("CRM", "Contacts")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}
func adminCrmContactDetail_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/portal/crm/contact/detail.html",
	))
	navMenu := loadNavMenu("CRM", "Contacts")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}
func adminCrmContactCreate_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/portal/crm/contact/create.html",
	))
	navMenu := loadNavMenu("CRM", "Contacts")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}
func adminCrmContactUpdate_GetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/portal/crm/contact/overview.html",
	))
	navMenu := loadNavMenu("CRM", "Contacts")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
		Model:       id,
	})
}

func adminSalesInvoicesOverview_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/home.html",
	))
	navMenu := loadNavMenu("Sales", "Invoices")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}
func adminHrEmployeeOverview_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/home.html",
	))
	navMenu := loadNavMenu("HR", "Employees")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
	})
}
func agentServiceDeskDashboard_GetHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles(
		"./web/layout/public.html",
		"./web/pages/home.html",
	))
	navMenu := loadNavMenu("ServiceDesk", "Tickets")
	template.Execute(w, Page{
		Title:       "Home",
		NavMenu:     navMenu,
		ShowSideBar: true,
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

	router.HandleFunc("/account/login", accountLogin_GetHandler).Methods("GET")
	router.HandleFunc("/account/login", accountLogin_PostHandler).Methods("POST")
	router.HandleFunc("/account/logoff", accountLogout_PostHandler).Methods("GET")

	// Public - ServiceDesk
	router.HandleFunc("/servicedesk", homeHandler).Methods("GET")
	router.HandleFunc("/servicedesk/tickets", homeHandler).Methods("GET")
	router.HandleFunc("/servicedesk/tickets/{id}", homeHandler).Methods("GET")
	router.HandleFunc("/servicedesk/tickets/submit", homeHandler).Methods("GET")

	// Agent - ServiceDesk
	router.HandleFunc("/portal/servicedesk", agentServiceDeskDashboard_GetHandler).Methods("GET")
	router.HandleFunc("/portal/servicedesk/tickets", homeHandler).Methods("GET")
	router.HandleFunc("/portal/servicedesk/tickets/{id}", homeHandler).Methods("GET")

	// Admin - HrEmployees
	router.HandleFunc("/portal/hr/employees", adminHrEmployeeOverview_GetHandler).Methods("GET")
	router.HandleFunc("/portal/hr/employees/{id}", homeHandler).Methods("GET")
	router.HandleFunc("/portal/hr/employees/create", homeHandler).Methods("GET")
	router.HandleFunc("/portal/hr/employees/edit/{id}", homeHandler).Methods("GET")

	// Admin - CrmCompanies
	router.HandleFunc("/portal/crm/companies", adminCrmCompanyOverview_GetHandler).Methods("GET")
	router.HandleFunc("/portal/crm/companies/{id}", adminCrmCompanyDetail_GetHandler).Methods("GET")
	router.HandleFunc("/portal/crm/companies/create", adminCrmCompanyCreate_GetHandler).Methods("GET")
	router.HandleFunc("/portal/crm/companies/edit/{id}", adminCrmCompanyUpdate_GetHandler).Methods("GET")

	// Admin - CrmContacts
	router.HandleFunc("/portal/crm/contacts", adminCrmContactOverview_GetHandler).Methods("GET")
	router.HandleFunc("/portal/crm/contacts/{id}", adminCrmContactDetail_GetHandler).Methods("GET")
	router.HandleFunc("/portal/crm/contacts/create", adminCrmContactCreate_GetHandler).Methods("GET")
	router.HandleFunc("/portal/crm/contacts/edit/{id}", adminCrmContactUpdate_GetHandler).Methods("GET")

	// Admin - SalesInvoices
	router.HandleFunc("/portal/sales/invoices", adminSalesInvoicesOverview_GetHandler).Methods("GET")

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
