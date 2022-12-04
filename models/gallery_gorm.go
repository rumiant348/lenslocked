package models

import "github.com/jinzhu/gorm"

type Gallery struct {
	gorm.Model
	UserID uint     `gorm:"not_null;index"`
	Title  string   `gorm:"not_null"`
	Images []string `gorm:"-"`
}

type galleryGorm struct {
	db *gorm.DB
}

var _ GalleryDB = &galleryGorm{}

func (gg *galleryGorm) ByID(id uint) (*Gallery, error) {
	var gallery Gallery
	db := gg.db.Where("id = ?", id)
	err := first(db, &gallery)
	if err != nil {
		return nil, err
	}
	return &gallery, nil
}

func (gg *galleryGorm) ByUserID(id uint) ([]Gallery, error) {
	var galleries []Gallery
	db := gg.db.Where("user_id = ?", id)
	err := db.Find(&galleries).Error
	if err != nil {
		return nil, err
	}
	return galleries, nil
}

func (gg *galleryGorm) Create(gallery *Gallery) error {
	return gg.db.Create(gallery).Error
}

func (gg *galleryGorm) Update(gallery *Gallery) error {
	return gg.db.Save(gallery).Error
}

func (gg *galleryGorm) Delete(id uint) error {
	gallery := Gallery{Model: gorm.Model{ID: id}}
	return gg.db.Delete(gallery).Error
}

// todo: repeat with tests
func (g *Gallery) ImagesSplitN(n int) [][]string {
	ret := make([][]string, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]string, 0)
	}
	for i, img := range g.Images {
		bucket := i % n
		ret[bucket] = append(ret[bucket], img)
	}
	return ret
}
