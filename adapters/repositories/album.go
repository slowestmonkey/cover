package repositories

import (
	"cover/core/domain"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SpotifyApiConfig struct {
	ApiUrl        string
	AccountApiUrl string
	ClientId      string
	ClientSecret  string
}

type Spotify struct {
	ApiUrl string
	Token  string
}

func NewSpotifyApi(config *SpotifyApiConfig) *Spotify {
	authEncoded := base64.StdEncoding.EncodeToString([]byte(config.ClientId + ":" + config.ClientSecret))
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	form := url.Values{}
	form.Add("grant_type", "client_credentials")

	request, err := http.NewRequest(http.MethodPost, config.AccountApiUrl+"/token", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatalln("cannot create a request")
	}

	request.Header.Set("Authorization", "Basic "+authEncoded)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.PostForm = form

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("cannot authenticate in Spotify")
	}

	authResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("cannot authenticate")
	}

	type auth struct {
		AccessToken string `json:"access_token"`
		// token_type   string
		// expires_in   int
	}

	authData := auth{}

	// FIXME: can I assign err like that?
	err = json.Unmarshal(authResponse, &authData)
	if err != nil {
		log.Fatalln("cannot parse auth response")
	}

	return &Spotify{ApiUrl: config.ApiUrl, Token: authData.AccessToken}
}

func (s *Spotify) Get(id string) (domain.Album, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest(http.MethodGet, s.ApiUrl+"/albums/"+id, nil)
	if err != nil {
		log.Fatalln("cannot create request")
	}

	request.Header.Set("Authorization", "Bearer "+s.Token)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("cannot receive album")
	}

	albumResponse, err := io.ReadAll(response.Body)
	album := domain.Album{}

	json.Unmarshal(albumResponse, &album)
	if err != nil {
		log.Fatalln("cannot parse album response")
	}

	return album, nil
}
