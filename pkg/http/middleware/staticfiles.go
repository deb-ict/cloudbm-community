package middleware

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/module/gallery"
	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_STATIC_ASSET_URI string = "/static/"
	GALLERY_BASE_URI         string = "/assets/gallery/"
)

type StaticFileMiddlewareConfig struct {
	StaticAsset StaticAssetConfig  `yaml:"Static"`
	Gallery     GalleryAssetConfig `yaml:"Gallery"`
}

type StaticAssetConfig struct {
	Uri    string `yaml:"UriPrefix"`
	Folder string `yaml:"Folder"`
}

type GalleryAssetConfig struct {
	Uri string `yaml:"UriPrefix"`
}

type StaticFileMiddleware struct {
	staticAssetBaseUri string
	staticAssetFolder  string
	galleryBaseUri     string
	galleryService     gallery.Service
}

func NewStaticFileMiddleware(galleryService gallery.Service, config *StaticFileMiddlewareConfig) *StaticFileMiddleware {
	if config == nil {
		config = &StaticFileMiddlewareConfig{}
	}
	config.StaticAsset.LoadEnvironment()
	config.StaticAsset.EnsureDefaults()
	config.Gallery.LoadEnvironment()
	config.Gallery.EnsureDefaults()

	return &StaticFileMiddleware{
		staticAssetBaseUri: config.StaticAsset.Uri,
		staticAssetFolder:  config.StaticAsset.Folder,
		galleryBaseUri:     config.Gallery.Uri,
		galleryService:     galleryService,
	}
}

func (m *StaticFileMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Uri path for css, js, ...
		if strings.HasPrefix(r.URL.Path, m.staticAssetBaseUri) {
			path := strings.Replace(r.URL.Path, m.staticAssetBaseUri, m.staticAssetFolder, 1)
			http.ServeFile(w, r, path)
			return
		}

		// Gallery service
		if strings.HasPrefix(r.URL.Path, m.galleryBaseUri) {
			// Parse the image id
			imageId := strings.Replace(r.URL.Path, m.galleryBaseUri, "", 1)
			if imageId == "" {
				return
			}

			// Get the image data
			buffer, contentType, fileName, err := m.galleryService.GetImageData(r.Context(), imageId)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			defer buffer.Close()

			// Return the image
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("Content-Disposition", "inline; filename="+fileName)
			io.Copy(w, buffer)
			return
		}

		// Do something here
		next.ServeHTTP(w, r)
	})
}

func (cfg *StaticFileMiddlewareConfig) LoadEnvironment() {
	cfg.StaticAsset.LoadEnvironment()
	cfg.Gallery.LoadEnvironment()
}

func (cfg *StaticFileMiddlewareConfig) EnsureDefaults() {
	cfg.StaticAsset.EnsureDefaults()
	cfg.Gallery.EnsureDefaults()
}

func (cfg *StaticAssetConfig) LoadEnvironment() {
	uri, ok := os.LookupEnv("STATIC_URI")
	if ok {
		logrus.Info("Override static uri from environment")
		cfg.Uri = uri
	}
	folder, ok := os.LookupEnv("STATIC_FOLDER")
	if ok {
		logrus.Info("Override static folder from environment")
		cfg.Folder = folder
	}
}

func (cfg *StaticAssetConfig) EnsureDefaults() {
	if cfg.Uri == "" {
		cfg.Uri = DEFAULT_STATIC_ASSET_URI
	}
	if !strings.HasPrefix(cfg.Uri, "/") {
		cfg.Uri = "/" + cfg.Uri
	}
	if !strings.HasSuffix(cfg.Uri, "/") {
		cfg.Uri = cfg.Uri + "/"
	}
	if cfg.Folder == "" {
		cfg.Folder = "/var/www/html/static/"
	}
	if !strings.HasSuffix(cfg.Folder, "/") {
		cfg.Folder = cfg.Folder + "/"
	}
}

func (cfg *GalleryAssetConfig) LoadEnvironment() {
	uri, ok := os.LookupEnv("GALLERY_URI")
	if ok {
		logrus.Info("Override gallery uri from environment")
		cfg.Uri = uri
	}
}

func (cfg *GalleryAssetConfig) EnsureDefaults() {
	if cfg.Uri == "" {
		cfg.Uri = GALLERY_BASE_URI
	}
	if !strings.HasPrefix(cfg.Uri, "/") {
		cfg.Uri = "/" + cfg.Uri
	}
	if !strings.HasSuffix(cfg.Uri, "/") {
		cfg.Uri = cfg.Uri + "/"
	}
}
