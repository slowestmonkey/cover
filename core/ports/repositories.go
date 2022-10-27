package ports

import "cover/core/domain"

type AlbumRepository interface {
	Get(id string) (domain.Album, error)
}
