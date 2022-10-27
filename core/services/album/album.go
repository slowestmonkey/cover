package albumsrv

import (
	"cover/core/domain"
	"cover/core/ports"
	"errors"
)

type service struct {
	albumRepository ports.AlbumRepository
}

func New(albumRepository ports.AlbumRepository) *service {
	return &service{
		albumRepository: albumRepository,
	}
}

func (s *service) Get(id string) (domain.Album, error) {
	album, err := s.albumRepository.Get(id)

	if err != nil {
		return domain.Album{}, errors.New("cannot get album")
	}

	return album, nil
}
