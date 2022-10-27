package main

import (
	"cover/adapters/handlers"
	"cover/adapters/repositories"
	albumsrv "cover/core/services/album"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	spotifyApiConfig := repositories.SpotifyApiConfig{
		ApiUrl:        os.Getenv("SPOTIFY_API_URL"),
		AccountApiUrl: os.Getenv("SPOTIFY_ACCOUNT_API_URL"),
		ClientId:      os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret:  os.Getenv("SPOTIFY_CLIENT_SECRET"),
	}

	albumRepository := repositories.NewSpotifyApi(&spotifyApiConfig)
	albumService := albumsrv.New(albumRepository)
	albumRouter := handlers.NewAlbumHttpHandler(albumService)

	r.Mount("/albums", albumRouter)

	http.ListenAndServe(":3333", r)
}
