package urlbuilder

import (
	"strings"

	"oneimg/backend/models"
)

// BuildFileURL returns the accessible URL for a stored file based on the configured access domain.
func BuildFileURL(path string, setting models.Settings) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return ""
	}

	// Absolute URLs are returned as-is.
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	domain := strings.TrimSpace(setting.FileAccessDomain)
	if domain == "" {
		return path
	}

	domain = strings.TrimRight(domain, "/")
	return domain + path
}

// DecorateImageURLs updates url fields for a single image.
func DecorateImageURLs(image *models.Image, setting models.Settings) {
	if image == nil {
		return
	}

	image.Url = BuildFileURL(image.Url, setting)
	image.Thumbnail = BuildFileURL(image.Thumbnail, setting)
}

// DecorateImagesURLs updates url fields for a slice of images.
func DecorateImagesURLs(images []models.Image, setting models.Settings) {
	if len(images) == 0 {
		return
	}

	for i := range images {
		DecorateImageURLs(&images[i], setting)
	}
}
