package gallery

import (
	"context"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery/model"
)

type Service interface {
	StringNormalizer() core.StringNormalizer
	FeatureProvider() core.FeatureProvider
	LanguageProvider() localization.LanguageProvider

	GetCategories(ctx context.Context, offset int64, limit int64, filter *model.CategoryFilter, sort *core.Sort) ([]*model.Category, int64, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetCategoryByName(ctx context.Context, language string, name string) (*model.Category, error)
	GetCategoryBySlug(ctx context.Context, language string, slug string) (*model.Category, error)
	CreateCategory(ctx context.Context, model *model.Category) (*model.Category, error)
	UpdateCategory(ctx context.Context, id string, model *model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, id string) error

	GetMediaFiles(ctx context.Context, offset int64, limit int64, filter *model.MediaFilter, sort *core.Sort) ([]*model.Media, int64, error)
	GetMediaFileById(ctx context.Context, id string) (*model.Media, error)
	GetMediaFileByName(ctx context.Context, language string, name string) (*model.Media, error)
	GetMediaFileBySlug(ctx context.Context, language string, slug string) (*model.Media, error)
	CreateMediaFile(ctx context.Context, model *model.Media) (*model.Media, error)
	UpdateMediaFile(ctx context.Context, id string, model *model.Media) (*model.Media, error)
	DeleteMediaFile(ctx context.Context, id string) error
}
