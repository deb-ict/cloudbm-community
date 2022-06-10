package contact

import (
	"context"
)

func (svc *service) GetUrlTypes(ctx context.Context, pageIndex int, pageSize int) (*UrlTypeList, error) {
	return svc.database.GetUrlTypeStore().GetUrlTypes(ctx, pageIndex, pageSize)
}

func (svc *service) GetUrlTypeById(ctx context.Context, id string) (*UrlType, error) {
	//TODO: Validate id

	return svc.database.GetUrlTypeStore().GetUrlTypeById(ctx, id)
}

func (svc *service) CreateUrlType(ctx context.Context, urlType UrlType) (*UrlType, error) {
	var err error

	//TODO: Validate model

	duplicate, err := svc.database.GetUrlTypeStore().GetUrlTypeByName(ctx, urlType.Name)
	if duplicate != nil {
		return nil, ErrUrlTypeDuplicate
	}
	if err != nil {
		return nil, err
	}

	if urlType.IsDefault {
		err = NewService().ResetDefaultUrlType(ctx)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetUrlTypeStore().CreateUrlType(ctx, urlType)
	if err != nil {
		return nil, err
	}

	return svc.database.GetUrlTypeStore().GetUrlTypeById(ctx, newId)
}

func (svc *service) UpdateUrlType(ctx context.Context, id string, urlType UrlType) (*UrlType, error) {
	var err error

	//TODO: Validate id
	//TODO: Validate model

	existing, err := svc.database.GetUrlTypeStore().GetUrlTypeById(ctx, id)
	if existing == nil && err != nil {
		return nil, ErrUrlTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	if existing.IsSystem {
		return nil, ErrUrlTypeReadOnly
	}

	if urlType.IsDefault {
		err = NewService().ResetDefaultUrlType(ctx)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetUrlTypeStore().UpdateUrlType(ctx, id, urlType)
	if err != nil {
		return nil, err
	}

	return svc.database.GetUrlTypeStore().GetUrlTypeById(ctx, id)
}

func (svc *service) DeleteUrlType(ctx context.Context, id string) error {
	var err error

	//TODO: Validate id

	existing, err := svc.database.GetUrlTypeStore().GetUrlTypeById(ctx, id)
	if existing == nil && err != nil {
		return ErrUrlTypeNotFound
	}
	if err != nil {
		return err
	}

	if existing.IsSystem {
		return ErrUrlTypeReadOnly
	}

	return svc.database.GetUrlTypeStore().DeleteUrlType(ctx, id)
}

func (svc *service) ResetDefaultUrlType(ctx context.Context) error {
	var err error

	defaultType, err := svc.database.GetUrlTypeStore().GetDefaultUrlType(ctx)
	if err != nil {
		return err
	}

	if defaultType != nil {
		defaultType.IsDefault = false
		err = svc.database.GetUrlTypeStore().UpdateUrlType(ctx, defaultType.Id, *defaultType)
		if err != nil {
			return err
		}
	}

	return nil
}
