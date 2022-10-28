package handlers

import (
	"cover/core/ports"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewAlbumHttpRouter(albumService ports.AlbumService) chi.Router {
	router := chi.NewRouter()
	handler := AlbumHttpHandler{albumService}

	router.Get("/{albumId}", handler.GetAlbum)

	return router
}

type AlbumHttpHandler struct {
	service ports.AlbumService
}

func (handler *AlbumHttpHandler) GetAlbum(w http.ResponseWriter, r *http.Request) {
	album, err := handler.service.Get(chi.URLParam(r, "albumId"))

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
}
