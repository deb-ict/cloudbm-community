package timesheet

type Service interface {
}

type Repository interface {
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
