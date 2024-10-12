package util

import (
	"path/filepath"
	"strings"
)

var AllowedType = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
}

func FileUploadCheck(image string) bool {
	ext := strings.ToLower(filepath.Ext(image))

	return AllowedType[ext]
}
