package v1

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/localization"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery/model"
	"github.com/deb-ict/go-router"
)

type ImageV1 struct {
	Id           string                `json:"id"`
	Translations []*ImageTranslationV1 `json:"translations"`
	FileName     string                `json:"fileName"`
	FileSize     int64                 `json:"fileSize"`
	MimeType     string                `json:"fileType"`
	Width        int32                 `json:"width"`
	Height       int32                 `json:"height"`
}

type ImageTranslationV1 struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type ImageListV1 struct {
	rest.PaginatedList
	Items []*ImageListItemV1 `json:"items"`
}

type ImageListItemV1 struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Summary  string `json:"summary"`
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	MimeType string `json:"fileType"`
}

type CreateImageV1 struct {
	Translations []*ImageTranslationV1 `json:"translations"`
	FileName     string                `json:"fileName"`
}

type UpdateImageV1 struct {
	Translations []*ImageTranslationV1 `json:"translations"`
	FileName     string                `json:"fileName"`
}

func (api *apiV1) GetImagesHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := api.parseImageFilterV1(r)
	paging := rest.GetPaging(r)
	sort := rest.GetSorting(r)

	result, count, err := api.service.GetImages(ctx, (paging.PageIndex-1)*paging.PageSize, paging.PageSize, filter, sort)
	if api.handleError(w, err) {
		return
	}

	response := ImageListV1{
		PaginatedList: rest.PaginatedList{
			PageIndex: paging.PageIndex,
			PageSize:  paging.PageSize,
			ItemCount: count,
		},
		Items: make([]*ImageListItemV1, 0),
	}
	for _, item := range result {
		response.Items = append(response.Items, ImageToListItemViewModelV1(item, filter.Language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetImageByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")
	result, err := api.service.GetImageById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := ImageToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateImageHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateImageV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateImage(ctx, ImageFromCreateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := ImageToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateImageHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	var model *UpdateImageV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateImage(ctx, id, ImageFromUpdateViewModelV1(model))
	if api.handleError(w, err) {
		return
	}

	response := ImageToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteImageHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := api.service.DeleteImage(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) UploadImageFileHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	err := r.ParseMultipartForm(10 << 20)
	if api.handleError(w, err) {
		return
	}

	file, header, err := r.FormFile("file")
	if api.handleError(w, err) {
		return
	}
	defer file.Close()

	mimeType := header.Header.Get("Content-Type")
	result, err := api.service.SetImageData(ctx, id, file, mimeType, header.Filename)
	if api.handleError(w, err) {
		return
	}

	response := ImageToViewModelV1(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DownloadImageFileHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := router.Param(r, "id")

	file, mimeType, originalFileName, err := api.service.GetImageData(ctx, id)
	if api.handleError(w, err) {
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Content-Disposition", "attachment; filename="+originalFileName)
	_, _ = io.Copy(w, file)
}

func (api *apiV1) parseImageFilterV1(r *http.Request) *model.ImageFilter {
	return &model.ImageFilter{
		Language: localization.GetHttpRequestLanguage(r, api.service.LanguageProvider()),
		Name:     r.URL.Query().Get("name"),
	}
}

func ImageToViewModelV1(model *model.Image) *ImageV1 {
	viewModel := &ImageV1{
		Id:           model.Id,
		Translations: make([]*ImageTranslationV1, 0),
		FileName:     model.FileName,
		FileSize:     model.FileSize,
		MimeType:     model.MimeType,
		Width:        model.Width,
		Height:       model.Height,
	}
	for _, translation := range model.Translations {
		viewModel.Translations = append(viewModel.Translations, ImageTranslationToViewModelV1(translation))
	}
	return viewModel
}

func ImageToListItemViewModelV1(model *model.Image, language string, defaultLanguage string) *ImageListItemV1 {
	translation := model.GetTranslation(language, defaultLanguage)
	return &ImageListItemV1{
		Id:       model.Id,
		Name:     translation.Name,
		Slug:     translation.Slug,
		Summary:  translation.Summary,
		FileName: model.OriginalFileName,
		FileSize: model.FileSize,
		MimeType: model.MimeType,
	}
}

func ImageFromCreateViewModelV1(viewModel *CreateImageV1) *model.Image {
	model := &model.Image{
		OriginalFileName: viewModel.FileName,
		Translations:     make([]*model.ImageTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ImageTranslationFromViewModelV1(translation))
	}
	return model
}

func ImageFromUpdateViewModelV1(viewModel *UpdateImageV1) *model.Image {
	model := &model.Image{
		OriginalFileName: viewModel.FileName,
		Translations:     make([]*model.ImageTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ImageTranslationFromViewModelV1(translation))
	}
	return model
}

func ImageTranslationToViewModelV1(model *model.ImageTranslation) *ImageTranslationV1 {
	return &ImageTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func ImageTranslationFromViewModelV1(viewModel *ImageTranslationV1) *model.ImageTranslation {
	return &model.ImageTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Summary:     viewModel.Summary,
		Description: viewModel.Description,
	}
}
