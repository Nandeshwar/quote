package quote

import "strings"

func validJPG(imageName string) bool {
	if strings.Contains(strings.ToLower(imageName), "jpg") || strings.Contains(strings.ToLower(imageName), "jpeg") {
		return true
	}
	return false
}
