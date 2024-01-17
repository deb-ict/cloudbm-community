package gallery

import (
	"context"
	"io"

	"github.com/deb-ict/cloudbm-community/pkg/core"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery/model"
)

type Service interface {
	StringNormalizer() core.StringNormalizer
	FeatureProvider() core.FeatureProvider
	LanguageProvider() localization.LanguageProvider
	StorageFolder() string

	GetImages(ctx context.Context, offset int64, limit int64, filter *model.ImageFilter, sort *core.Sort) ([]*model.Image, int64, error)
	GetImageById(ctx context.Context, id string) (*model.Image, error)
	GetImageByName(ctx context.Context, language string, name string) (*model.Image, error)
	GetImageBySlug(ctx context.Context, language string, slug string) (*model.Image, error)
	CreateImage(ctx context.Context, model *model.Image) (*model.Image, error)
	UpdateImage(ctx context.Context, id string, model *model.Image) (*model.Image, error)
	DeleteImage(ctx context.Context, id string) error
	GetImageData(ctx context.Context, id string) (io.ReadCloser, string, string, error)
	SetImageData(ctx context.Context, id string, file io.Reader, mimeType string, originalFileName string) (*model.Image, error)
}
