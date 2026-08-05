package gophers

import (
	"mime"
	"slices"
)

type MimeContentType struct {
}

func (me MimeContentType) GetFileExtension(contentType string) (string, error) {
	var extensions, e = mime.ExtensionsByType(contentType)
	if e != nil {
		return "", e
	}
	if slices.Contains(extensions, ".jpg") {
		return ".jpg", e
	}
	return extensions[0], e
}
