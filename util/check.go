package util

import (
	"path/filepath"
	"regexp"
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

func EmailCheck(email string) bool {
	const eMailcheck = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(eMailcheck)
	return regex.MatchString(email)
}
