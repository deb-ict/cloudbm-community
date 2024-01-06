package service

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"time"

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
	model.StorageFolder = ""
	model.FileName = ""
	model.FileSize = 0
	model.MimeType = ""
	model.Width = 0
	model.Height = 0

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

func (svc *service) GetImageData(ctx context.Context, id string) (io.ReadCloser, string, error) {
	data, err := svc.database.Images().GetImageById(ctx, id)
	if err != nil {
		return nil, "", err
	}
	if data == nil {
		return nil, "", gallery.ErrImageNotFound
	}

	filePath := filepath.Join(data.StorageFolder, data.FileName)
	_, err = os.Stat(filePath)
	if err != nil {
		return nil, "", err
	}

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return nil, "", err
	}

	return file, data.MimeType, nil
}

func (svc *service) SetImageFile(ctx context.Context, id string, file io.Reader, mimeType string, originalFileName string) (*model.Image, error) {
	// Get the file extension based on the mime type
	fileExt := ""
	switch mimeType {
	case "image/jpeg":
		fileExt = ".jpg"
	case "image/png":
		fileExt = ".pgn"
	default:
		return nil, gallery.ErrImageFormatNotSupported
	}

	// Get the storage folder
	now := time.Now().UTC()
	localFolder := filepath.Join(svc.StorageFolder(), "images", fmt.Sprintf("%04d/%02d/", now.Year(), int(now.Month())))
	core.EnsureFolder(localFolder)

	// Open the local file
	localFileName := fmt.Sprintf("%s%s", id, fileExt)
	localFilePath := filepath.Join(localFolder, localFileName)
	localFile, err := os.OpenFile(localFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer localFile.Close()

	// Copy the file data
	_, err = io.Copy(localFile, file)
	if err != nil {
		return nil, err
	}

	// Set the image file info
	image, err := svc.SetImageFileInfo(ctx, id, localFolder, localFileName, originalFileName, mimeType)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (svc *service) SetImageFileInfo(ctx context.Context, id string, localFolder string, localFileName string, originalFileName string, mimeType string) (*model.Image, error) {
	data, err := svc.database.Images().GetImageById(ctx, id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, gallery.ErrImageNotFound
	}

	filePath := filepath.Join(localFolder, localFileName)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config image.Config
	switch mimeType {
	case "image/jpeg":
		config, err = jpeg.DecodeConfig(file)
	case "image/png":
		config, err = png.DecodeConfig(file)
	default:
		err = gallery.ErrImageFormatNotSupported
	}
	if err != nil {
		return nil, err
	}

	data.OriginalFileName = originalFileName
	data.StorageFolder = localFolder
	data.FileName = localFileName
	data.FileSize = fileInfo.Size()
	data.MimeType = mimeType
	data.Width = config.Width
	data.Height = config.Height

	err = svc.database.Images().UpdateImage(ctx, data)
	if err != nil {
		return nil, err
	}

	return svc.GetImageById(ctx, id)
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
