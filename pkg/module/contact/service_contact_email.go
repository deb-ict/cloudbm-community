package contact

import (
	"context"
)

func (svc *service) GetContactEmails(ctx context.Context, contactId string, pageIndex int, pageSize int) (*EmailList, error) {
	var err error

	//TODO: Validate contactId

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactEmailStore().GetContactEmails(ctx, contactId, pageIndex, pageSize)
}

func (svc *service) GetContactEmailById(ctx context.Context, contactId string, id string) (*Email, error) {
	var err error

	//TODO: Validate contactId
	//TODO: Validate id

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactEmailStore().GetContactEmailById(ctx, contactId, id)
}

func (svc *service) CreateContactEmail(ctx context.Context, contactId string, email Email) (*Email, error) {
	var err error

	//TODO: Validate contactId
	//TODO: Validate id

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	emailType, err := svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, email.TypeId)
	if emailType == nil && err == nil {
		return nil, ErrEmailTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetContactEmailStore().GetContactEmailByType(ctx, contactId, email.TypeId)
	if duplicate != nil {
		return nil, ErrContactEmailDuplicate
	}
	if err != nil {
		return nil, err
	}

	if email.IsPrimary {
		err = svc.ResetPrimaryContactEmail(ctx, contactId)
		if err != nil {
			return nil, err
		}
	}

	newId, err := svc.database.GetContactEmailStore().CreateContactEmail(ctx, contactId, email)
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactEmailStore().GetContactEmailById(ctx, contactId, newId)
}

func (svc *service) UpdateContactEmail(ctx context.Context, contactId string, id string, email Email) (*Email, error) {
	var err error

	//TODO: Validate contactId
	//TODO: Validate id

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	existing, err := svc.database.GetContactEmailStore().GetContactEmailById(ctx, contactId, id)
	if existing == nil {
		return nil, ErrContactEmailNotFound
	}
	if err != nil {
		return nil, err
	}

	emailType, err := svc.database.GetEmailTypeStore().GetEmailTypeById(ctx, email.TypeId)
	if emailType == nil && err == nil {
		return nil, ErrEmailTypeNotFound
	}
	if err != nil {
		return nil, err
	}

	duplicate, err := svc.database.GetContactEmailStore().GetContactEmailByType(ctx, contactId, email.TypeId)
	if duplicate != nil {
		return nil, ErrContactEmailDuplicate
	}
	if err != nil {
		return nil, err
	}

	if email.IsPrimary {
		err = svc.ResetPrimaryContactEmail(ctx, contactId)
		if err != nil {
			return nil, err
		}
	}

	err = svc.database.GetContactEmailStore().UpdateContactEmail(ctx, contactId, id, email)
	if err != nil {
		return nil, err
	}

	return svc.database.GetContactEmailStore().GetContactEmailById(ctx, contactId, id)
}

func (svc *service) DeleteContactEmail(ctx context.Context, contactId string, id string) error {
	var err error

	//TODO: Validate contactId
	//TODO: Validate id

	contact, err := svc.database.GetContactStore().GetContactById(ctx, contactId)
	if contact == nil && err == nil {
		return ErrContactNotFound
	}
	if err != nil {
		return err
	}

	existing, err := svc.database.GetContactEmailStore().GetContactEmailById(ctx, contactId, id)
	if existing == nil {
		return ErrContactEmailNotFound
	}
	if err != nil {
		return err
	}

	return svc.database.GetContactEmailStore().DeleteContactEmail(ctx, contactId, id)
}

func (svc *service) ResetPrimaryContactEmail(ctx context.Context, contactId string) error {
	var err error

	primary, err := svc.database.GetContactEmailStore().GetContactPrimaryEmail(ctx, contactId)
	if err != nil {
		return err
	}

	if primary != nil {
		primary.IsPrimary = false
		err = svc.database.GetContactEmailStore().UpdateContactEmail(ctx, contactId, primary.Id, *primary)
		if err != nil {
			return err
		}
	}

	return nil
}
