package session

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/session/model"
)

type Service interface {
	FeatureProvider() core.FeatureProvider

	GetSessions(ctx context.Context, offset int64, limit int64, filter *model.SessionFilter, sort *core.Sort) ([]*model.Session, int64, error)
	GetSessionById(ctx context.Context, id string) (*model.Session, error)
	CreateSession(ctx context.Context, model *model.Session) (*model.Session, error)
	UpdateSession(ctx context.Context, id string, model *model.Session) (*model.Session, error)
	DeleteSession(ctx context.Context, id string) error

	LoadSession(ctx context.Context, id string) (*model.Session, error)
	SaveSession(ctx context.Context, session *model.Session) (*model.Session, error)

	CleanupExpiredSessions(ctx context.Context) error
}
