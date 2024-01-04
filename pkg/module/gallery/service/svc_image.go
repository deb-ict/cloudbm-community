package service

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery/model"
)

func (svc *service) GetImages(ctx context.Context, offset int64, limit int64, filter *model.ImageFilter, sort *core.Sort) ([]*model.Image, int64, error) {
	filter.Language = localization.NormalizeLanguage(filter.Language)
	data, count, err := svc.database.Images().GetImages(ctx, offset, limit, filter, sort)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (svc *service) GetImageById(ctx context.Context, id string) (*model.Image, error) {
	data, err := svc.database.Images().GetImageById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, gallery.ErrImageNotFound
	}

	return data, nil
}

func (svc *service) GetImageByName(ctx context.Context, language string, name string) (*model.Image, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedName := svc.stringNormalizer.NormalizeString(name)

	data, err := svc.database.Images().GetImageByName(ctx, normalizedLanguage, normalizedName)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, gallery.ErrImageNotFound
	}

	return data, nil
}

func (svc *service) GetImageBySlug(ctx context.Context, language string, slug string) (*model.Image, error) {
	normalizedLanguage := localization.NormalizeLanguage(language)
	normalizedSlug := svc.stringNormalizer.NormalizeString(slug)

	data, err := svc.database.Images().GetImageBySlug(ctx, normalizedLanguage, normalizedSlug)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, gallery.ErrImageNotFound
	}

	return data, nil
}

func (svc *service) CreateImage(ctx context.Context, model *model.Image) (*model.Image, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = ""

	err := svc.checkDuplicateImage(ctx, model)
	if err != nil {
		return nil, err
	}

	newId, err := svc.database.Images().CreateImage(ctx, model)
	if err != nil {
		return nil, err
	}

	return svc.GetImageById(ctx, newId)
}

func (svc *service) UpdateImage(ctx context.Context, id string, model *model.Image) (*model.Image, error) {
	model.Normalize(svc.stringNormalizer)
	model.Id = id

	err := svc.checkDuplicateImage(ctx, model)
	if err != nil {
		return nil, err
	}

	data, err := svc.database.Images().GetImageById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, gallery.ErrImageNotFound
	}
	data.UpdateModel(model)

	err = svc.database.Images().UpdateImage(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetImageById(ctx, id)
}

func (svc *service) DeleteImage(ctx context.Context, id string) error {
	data, err := svc.database.Images().GetImageById(ctx, id)
	if err != nil {
		return err
	}
	if data == nil {
		return gallery.ErrImageNotFound
	}

	err = svc.database.Images().DeleteImage(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) checkDuplicateImage(ctx context.Context, model *model.Image) error {
	for _, translation := range model.Translations {
		if err := svc.checkDuplicateImageName(ctx, model, translation); err != nil {
			return err
		}
		if err := svc.checkDuplicateImageSlug(ctx, model, translation); err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) checkDuplicateImageName(ctx context.Context, model *model.Image, translation *model.ImageTranslation) error {
	duplicate, err := svc.database.Images().GetImageByName(ctx, translation.Language, translation.NormalizedName)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return gallery.ErrImageDuplicateName
	}
	return nil
}

func (svc *service) checkDuplicateImageSlug(ctx context.Context, model *model.Image, translation *model.ImageTranslation) error {
	duplicate, err := svc.database.Images().GetImageBySlug(ctx, translation.Language, translation.Slug)
	if err != nil {
		return err
	}
	if duplicate != nil && duplicate.Id != model.Id {
		return gallery.ErrImageDuplicateSlug
	}
	return nil
}
