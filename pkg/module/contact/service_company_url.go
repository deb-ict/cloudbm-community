package contact

import (
	"context"
)

func (svc *service) GetCompanyUrls(ctx context.Context, companyId string, pageIndex int, pageSize int) (*UrlList, error) {
	var err error

	//TODO: Validate companyId

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyUrlStore().GetCompanyUrls(ctx, companyId, pageIndex, pageSize)
}

func (svc *service) GetCompanyUrlById(ctx context.Context, companyId string, id string) (*Url, error) {
	var err error

	//TODO: Validate companyId
	//TODO: Validate id

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyUrlStore().GetCompanyUrlById(ctx, companyId, id)
}

func (svc *service) CreateCompanyUrl(ctx context.Context, companyId string, url Url) (*Url, error) {
	var err error

	//TODO: Validate companyId
	//TODO: Validate id

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	urlType, err := svc.database.GetUrlTypeStore().GetUrlTypeById(ctx, url.TypeId)
	if urlType == nil && err == nil {
		return nil, ErrPhoneTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetCompanyUrlStore().GetCompanyUrlByType(ctx, companyId, url.TypeId)
	if duplicate != nil {
		return nil, ErrCompanyUrlDuplicate
	}
	if err != nil && err != ErrCompanyUrlNotFound {
		return nil, err
	}

	if url.IsPrimary {
		err = svc.ResetPrimaryCompanyUrl(ctx, companyId)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetCompanyUrlStore().CreateCompanyUrl(ctx, companyId, url)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyUrlStore().GetCompanyUrlById(ctx, companyId, newId)
}

func (svc *service) UpdateCompanyUrl(ctx context.Context, companyId string, id string, url Url) (*Url, error) {
	var err error

	//TODO: Validate companyId
	//TODO: Validate id

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return nil, ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}

	existing, err := svc.database.GetCompanyUrlStore().GetCompanyUrlById(ctx, companyId, id)
	if existing == nil && err == nil {
		return nil, ErrCompanyUrlNotFound
	}
	if err != nil {
		return nil, err
	}

	urlType, err := svc.database.GetUrlTypeStore().GetUrlTypeById(ctx, url.TypeId)
	if urlType == nil && err == nil {
		return nil, ErrPhoneTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetCompanyUrlStore().GetCompanyUrlByType(ctx, companyId, url.TypeId)
	if duplicate != nil {
		return nil, ErrCompanyUrlDuplicate
	}
	if err != nil && err != ErrCompanyUrlNotFound {
		return nil, err
	}

	if url.IsPrimary {
		err = svc.ResetPrimaryCompanyUrl(ctx, companyId)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetCompanyUrlStore().UpdateCompanyUrl(ctx, companyId, id, url)
	if err != nil {
		return nil, err
	}

	return svc.database.GetCompanyUrlStore().GetCompanyUrlById(ctx, companyId, id)
}

func (svc *service) DeleteCompanyUrl(ctx context.Context, companyId string, id string) error {
	var err error

	//TODO: Validate companyId
	//TODO: Validate id

	company, err := svc.database.GetCompanyStore().GetCompanyById(ctx, companyId)
	if company == nil && err == nil {
		return ErrCompanyNotFound
	}
	if err != nil {
		return err
	}

	existing, err := svc.database.GetCompanyUrlStore().GetCompanyUrlById(ctx, companyId, id)
	if existing == nil && err == nil {
		return ErrCompanyUrlNotFound
	}
	if err != nil {
		return err
	}

	return svc.database.GetCompanyUrlStore().DeleteCompanyUrl(ctx, companyId, id)
}

func (svc *service) ResetPrimaryCompanyUrl(ctx context.Context, companyId string) error {
	var err error

	primary, err := svc.database.GetCompanyUrlStore().GetCompanyPrimaryUrl(ctx, companyId)
	if err != nil && err != ErrCompanyUrlNotFound {
		return err
	}

	if primary != nil {
		primary.IsPrimary = false
		err = svc.database.GetCompanyUrlStore().UpdateCompanyUrl(ctx, companyId, primary.Id, *primary)
		if err != nil {
			return err
		}
	}

	return nil
}
