package session

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/session/model"
)

type Database interface {
	Sessions() SessionRepository
}

type SessionRepository interface {
	GetSessions(ctx context.Context, offset int64, limit int64, filter *model.SessionFilter, sort *core.Sort) ([]*model.Session, int64, error)
	GetSessionById(ctx context.Context, id string) (*model.Session, error)
	CreateSession(ctx context.Context, model *model.Session) (string, error)
	UpdateSession(ctx context.Context, model *model.Session) error
	DeleteSession(ctx context.Context, model *model.Session) error
}
