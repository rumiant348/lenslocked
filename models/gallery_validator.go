package models

type galleryValidator struct {
	GalleryDB
}

type galleryValFn func(*Gallery) error

func runGalleryValFns(gallery *Gallery, fns ...galleryValFn) error {
	for _, fn := range fns {
		if err := fn(gallery); err != nil {
			return err
		}
	}
	return nil
}

const (
	ErrUserIDRequired modelError = "models: user ID is required"
	ErrTitleRequired  modelError = "models: title is required"
)

func (gv *galleryValidator) userIDRequired(g *Gallery) error {
	if g.UserID <= 0 {
		return ErrUserIDRequired
	}
	return nil
}

func (gv *galleryValidator) titleRequired(g *Gallery) error {
	if g.Title == "" {
		return ErrTitleRequired
	}
	return nil
}

func (gv *galleryValidator) nonZeroId(gallery *Gallery) error {
	if gallery.ID <= 0 {
		return ErrIDInvalid
	}
	return nil
}

func (gv *galleryValidator) Create(gallery *Gallery) error {
	err := runGalleryValFns(gallery,
		gv.userIDRequired,
		gv.titleRequired,
	)
	if err != nil {
		return err
	}
	return gv.GalleryDB.Create(gallery)
}

func (gv *galleryValidator) Update(gallery *Gallery) error {
	err := runGalleryValFns(gallery,
		gv.titleRequired,
		gv.titleRequired,
	)
	if err != nil {
		return err
	}
	return gv.GalleryDB.Update(gallery)
}

func (gv *galleryValidator) Delete(id uint) error {
	var gallery Gallery
	gallery.ID = id
	if err := runGalleryValFns(&gallery, gv.nonZeroId); err != nil {
		return err
	}
	return gv.GalleryDB.Delete(gallery.ID)
}
