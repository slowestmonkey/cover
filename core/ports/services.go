package ports

import "cover/core/domain"

type AlbumService interface {
	Get(id string) (domain.Album, error)
}
