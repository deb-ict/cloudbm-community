package service

import (
	"context"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/module/session"
	"github.com/deb-ict/cloudbm-community/pkg/module/session/model"
)

func (svc *service) GetSessions(ctx context.Context, offset int64, limit int64, filter *model.SessionFilter, sort *core.Sort) ([]*model.Session, int64, error) {
	data, count, err := svc.database.Sessions().GetSessions(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetSessionById(ctx context.Context, id string) (*model.Session, error) {
	data, err := svc.database.Sessions().GetSessionById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, session.ErrSessionNotFound
	}
	if data.HasExpired() {
		return data, session.ErrSessionExpired
	}

	return data, nil
}

func (svc *service) CreateSession(ctx context.Context, model *model.Session) (*model.Session, error) {
	model.Id = ""
	model.CreatedAt = time.Now().UTC()
	model.SetExpiration(model.Lifetime)

	newId, err := svc.database.Sessions().CreateSession(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetSessionById(ctx, newId)
}

func (svc *service) UpdateSession(ctx context.Context, id string, model *model.Session) (*model.Session, error) {
	model.Id = id

	data, err := svc.database.Sessions().GetSessionById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, session.ErrSessionNotFound
	}
	data.UpdateModel(model)

	err = svc.database.Sessions().UpdateSession(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetSessionById(ctx, id)
}

func (svc *service) DeleteSession(ctx context.Context, id string) error {
	data, err := svc.database.Sessions().GetSessionById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return session.ErrSessionNotFound
	}

	err = svc.database.Sessions().DeleteSession(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) LoadSession(ctx context.Context, id string) (*model.Session, error) {
	return svc.getOrCreateSession(ctx, id)
}

func (svc *service) SaveSession(ctx context.Context, model *model.Session) (*model.Session, error) {
	data, err := svc.getOrCreateSession(ctx, model.Id)
	if err != nil {
		return nil, err
	}
	data.UpdateModel(model)
	data.SetExpiration(data.Lifetime)

	data, err = svc.UpdateSession(ctx, model.Id, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (svc *service) getOrCreateSession(ctx context.Context, id string) (*model.Session, error) {
	var data *model.Session = nil
	var err error = nil

	if id != "" {
		data, err = svc.GetSessionById(ctx, id)

		if err != nil {
			if err == session.ErrSessionExpired {
				err = svc.DeleteSession(ctx, id)
				data = nil
			}

			if err != session.ErrSessionNotFound {
				return nil, err
			}
		}
	}

	if data == nil {
		data = &model.Session{
			Lifetime:             15 * time.Minute,
			UseSlidingExpiration: true,
			Data:                 make(map[string]string),
		}
		data, err = svc.CreateSession(ctx, data)
		if err != nil {
			return nil, err
		}
	} else {
		if data.UseSlidingExpiration {
			data.SetExpiration(data.Lifetime)
			err = svc.database.Sessions().UpdateSession(ctx, data)
			if err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}

func (svc *service) CleanupExpiredSessions(ctx context.Context) error {
	filter := &model.SessionFilter{
		ExpiresBefore: time.Now().UTC(),
	}

	for {
		sessions, _, err := svc.GetSessions(ctx, 0, 100, filter, nil)
		if err != nil {
			return err
		}

		if len(sessions) == 0 {
			break
		}

		for _, session := range sessions {
			err = svc.DeleteSession(ctx, session.Id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
