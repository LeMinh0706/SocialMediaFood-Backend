package util

import (
	"path/filepath"
	"strings"
)

var AllowType = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
}

func FileExtCheck(image string) bool {
	ext := strings.ToLower(filepath.Ext(image))

	return AllowType[ext]
}
