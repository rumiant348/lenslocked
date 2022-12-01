package models

type GalleryDB interface {
	Create(gallery *Gallery) error
}
