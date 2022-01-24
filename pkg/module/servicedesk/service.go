package servicedesk

import (
	"context"
)

type TicketStore interface {
	SearchTickets(ctx context.Context, TicketSearchParams, pageIndex int, pageSize int) (*TicketPage, error)
	GetTicketById(ctx context.Context, id string) (*Ticket, error)
}

type SomeStore interface {
}

type DbStore struct {
}
type GrpcStore struct {
}

type service struct {
}

func (svc *service) SearchTickets() {

}
