package handlers

import (
	"cover/core/ports"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewAlbumHttpHandler(albumService ports.AlbumService) chi.Router {
	router := chi.NewRouter()

	// TODO: move it out and refactor
	router.Get("/{albumId}", func(w http.ResponseWriter, r *http.Request) {
		album, err := albumService.Get(chi.URLParam(r, "albumId"))

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		albumJson, err := json.Marshal(album)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(albumJson)
	})

	return router
}
