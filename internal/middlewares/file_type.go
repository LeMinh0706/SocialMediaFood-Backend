package middlewares

var AllowedType = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
}

// Allow mime
var AllowedMimeTypes = map[string]bool{
	"image/png":  true,
	"image/jpeg": true,
}
