package v1

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/deb-ict/cloudbm-community/pkg/http/rest"
	"github.com/deb-ict/cloudbm-community/pkg/module/gallery/model"
	"github.com/gorilla/mux"
)

type ImageV1 struct {
	Id           string                `json:"id"`
	Translations []*ImageTranslationV1 `json:"translations"`
	FileName     string                `json:"fileName"`
	FileSize     int64                 `json:"fileSize"`
	MimeType     string                `json:"fileType"`
	Width        int                   `json:"width"`
	Height       int                   `json:"height"`
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

	filter := api.parseImageFilter(r)
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
		response.Items = append(response.Items, ImageToListItemViewModel(item, filter.Language, api.service.LanguageProvider().DefaultLanguage(ctx)))
	}

	rest.WriteResult(w, response)
}

func (api *apiV1) GetImageByIdHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]
	result, err := api.service.GetImageById(ctx, id)
	if api.handleError(w, err) {
		return
	}

	response := ImageToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) CreateImageHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model *CreateImageV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.CreateImage(ctx, ImageFromCreateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ImageToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) UpdateImageHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var model *UpdateImageV1
	err := json.NewDecoder(r.Body).Decode(&model)
	if api.handleError(w, err) {
		return
	}

	result, err := api.service.UpdateImage(ctx, id, ImageFromUpdateViewModel(model))
	if api.handleError(w, err) {
		return
	}

	response := ImageToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DeleteImageHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	err := api.service.DeleteImage(ctx, id)
	if api.handleError(w, err) {
		return
	}

	rest.WriteStatus(w, http.StatusNoContent)
}

func (api *apiV1) UploadImageFileHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

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
	result, err := api.service.SetImageFile(ctx, id, file, mimeType, header.Filename)
	if api.handleError(w, err) {
		return
	}

	response := ImageToViewModel(result)
	rest.WriteResult(w, response)
}

func (api *apiV1) DownloadImageFileHandlerV1(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	file, mimeType, err := api.service.GetImageData(ctx, id)
	if api.handleError(w, err) {
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", mimeType)
	_, _ = io.Copy(w, file)
}

func (api *apiV1) parseImageFilter(r *http.Request) *model.ImageFilter {
	filter := &model.ImageFilter{}

	filter.Language = r.URL.Query().Get("language")
	if filter.Language == "" {
		filter.Language = api.service.LanguageProvider().UserLanguage(r.Context())
	}
	filter.Name = r.URL.Query().Get("name")

	return filter
}

func ImageToViewModel(model *model.Image) *ImageV1 {
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
		viewModel.Translations = append(viewModel.Translations, ImageTranslationToViewModel(translation))
	}
	return viewModel
}

func ImageToListItemViewModel(model *model.Image, language string, defaultLanguage string) *ImageListItemV1 {
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

func ImageFromCreateViewModel(viewModel *CreateImageV1) *model.Image {
	model := &model.Image{
		OriginalFileName: viewModel.FileName,
		Translations:     make([]*model.ImageTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ImageTranslationFromViewModel(translation))
	}
	return model
}

func ImageFromUpdateViewModel(viewModel *UpdateImageV1) *model.Image {
	model := &model.Image{
		OriginalFileName: viewModel.FileName,
		Translations:     make([]*model.ImageTranslation, 0),
	}
	for _, translation := range viewModel.Translations {
		model.Translations = append(model.Translations, ImageTranslationFromViewModel(translation))
	}
	return model
}

func ImageTranslationToViewModel(model *model.ImageTranslation) *ImageTranslationV1 {
	return &ImageTranslationV1{
		Language:    model.Language,
		Name:        model.Name,
		Slug:        model.Slug,
		Summary:     model.Summary,
		Description: model.Description,
	}
}

func ImageTranslationFromViewModel(viewModel *ImageTranslationV1) *model.ImageTranslation {
	return &model.ImageTranslation{
		Language:    viewModel.Language,
		Name:        viewModel.Name,
		Slug:        viewModel.Slug,
		Summary:     viewModel.Summary,
		Description: viewModel.Description,
	}
}
