package contact

import (
	"context"
)

func (svc *service) GetEmailTypes(ctx context.Context, pageIndex int, pageSize int) (*EmailTypeList, error) {
	return svc.database.GetEmailTypeStore().GetEmailTypes(ctx, pageIndex, pageSize)
}

func (svc *service) GetEmailTypeById(ctx context.Context, id string) (*EmailType, error) {
	//TODO: Validate id

	return svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, id)
}

func (svc *service) CreateEmailType(ctx context.Context, emailType EmailType) (*EmailType, error) {
	var err error

	//TODO: Validate model

	duplicate, err := svc.database.GetEmailTypeStore().GetEmailTypeByName(ctx, emailType.Name)
	if duplicate != nil {
		return nil, ErrEmailTypeDuplicate
	}
	if err != nil {
		return nil, err
	}

	if emailType.IsDefault {
		err = NewService().ResetDefaultEmailType(ctx)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetEmailTypeStore().CreateEmailType(ctx, emailType)
	if err != nil {
		return nil, err
	}

	return svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, newId)
}

func (svc *service) UpdateEmailType(ctx context.Context, id string, emailType EmailType) (*EmailType, error) {
	var err error

	//TODO: Validate id
	//TODO: Validate model

	existing, err := svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, id)
	if existing == nil && err != nil {
		return nil, ErrEmailTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	if existing.IsSystem {
		return nil, ErrEmailTypeReadOnly
	}

	if emailType.IsDefault {
		err = NewService().ResetDefaultEmailType(ctx)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetEmailTypeStore().UpdateEmailType(ctx, id, emailType)
	if err != nil {
		return nil, err
	}

	return svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, id)
}

func (svc *service) DeleteEmailType(ctx context.Context, id string) error {
	var err error

	//TODO: Validate id

	existing, err := svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, id)
	if existing == nil && err != nil {
		return ErrEmailTypeNotFound
	}
	if err != nil {
		return err
	}

	if existing.IsSystem {
		return ErrEmailTypeReadOnly
	}

	return svc.database.GetEmailTypeStore().DeleteEmailType(ctx, id)
}

func (svc *service) ResetDefaultEmailType(ctx context.Context) error {
	var err error

	defaultType, err := svc.database.GetEmailTypeStore().GetDefaultEmailType(ctx)
	if err != nil {
		return err
	}

	if defaultType != nil {
		defaultType.IsDefault = false
		err = svc.database.GetEmailTypeStore().UpdateEmailType(ctx, defaultType.Id, *defaultType)
		if err != nil {
			return err
		}
	}

	return nil
}
