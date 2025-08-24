package middleware

import (
	"dialog-service/service"
	"dialog-service/service/auth"
	"log"
	"net/http"
	"strings"
)

func CheckToken(services *service.Service) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			token := strings.TrimSpace(strings.Replace(authorization, "Bearer", "", 1))

			res, err := services.Auth.UserCheckToken.Handle(r.Context(), &auth.CheckTokenData{Token: token})

			if err != nil {
				log.Println("Authorization error:", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if res == nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			r.Header.Add("X-Authorized-User-Id", res.UserID)

			h.ServeHTTP(w, r)
		})
	}
}
