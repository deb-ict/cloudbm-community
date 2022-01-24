package servicedesk

import "time"

type Ticket struct {
	Id          string      `json:"id"`
	Type        Type        `json:"type"`       // incident, request, ...
	Status      Status      `json:"status"`     // new, inprogress, closed, ...
	Priority    Priority    `json:"priority"`   // low, medium, high
	Resolution  *Resolution `json:"resolution"` // unresolved, duplicate, ...
	Channel     Channel     `json:"channel"`    // web, email, ...
	Components  []Component `json:"components"`
	Description string      `json:"description"`
	Submitter   User        `json:"submitter"`
	Assignee    User        `json:"assignee"`
	Group       *UserGroup  `json:"group"`
	Followers   []User      `json:"followers"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type TicketPage struct {
	PageIndex int      `json:"page_index"`
	PageSize  int      `json:"page_size"`
	Count     int      `json:"count"`
	Data      []Ticket `json:"data"`
}

type TicketSearchParams struct {
}

type Comment struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}

type Attachment struct {
	Id string `json:"id"`
}

type Type struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Status struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Priority struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Resolution struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Channel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Component struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserGroup struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
