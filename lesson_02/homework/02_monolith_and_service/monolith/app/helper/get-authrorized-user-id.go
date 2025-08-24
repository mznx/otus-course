package helper

import "net/http"

func GetAuthorizedUserId(r *http.Request) string {
	return r.Header.Get("X-Authorized-User-Id")
}
