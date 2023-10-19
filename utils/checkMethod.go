package utils

import "net/http"

func CheckMethod(methods []string, r *http.Request) bool {
	for _, s := range methods {
		if r.Method == s {
			return true
		}
	}
	return false
}
