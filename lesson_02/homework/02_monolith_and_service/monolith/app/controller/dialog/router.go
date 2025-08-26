package dialog

import (
	"io"
	"monolith/controller/middleware"
	"monolith/infrastructure/config"
	"monolith/service"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func DialogRouter(router chi.Router, services *service.Service, config *config.Config) {
	router.Route("/dialog/{user_id}/send", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			proxyRequest(config, w, r)
		})
	})

	router.Route("/dialog/{user_id}/list", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			proxyRequest(config, w, r)
		})
	})
}

func proxyRequest(config *config.Config, w http.ResponseWriter, r *http.Request) {
	parsedURL, err := url.Parse(config.Services.Dialog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	proxyReq, err := http.NewRequest(r.Method, parsedURL.String()+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	proxyReq.Header = r.Header.Clone()
	proxyReq.Header.Set("Host", parsedURL.Host)

	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
