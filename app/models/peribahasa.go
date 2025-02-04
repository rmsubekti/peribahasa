package models

import (
	"errors"
)

// Peribahasa table belong to single Peribahasa and Category
type Peribahasa struct {
	ID       uint   `gorm:"primary_key"  json:"id"`
	TeksAsli string `gorm:"not null;unique" json:"asli"`
	Arti     string `json:"arti"`
	IDAsal   uint   `sql:"type:int REFERENCES asal(id)" json:"-"`
	Asal     Asal   `gorm:"auto_preload;association_foreignkey:IDAsal"`
	IDJenis  uint   `sql:"type:int REFERENCES jenis(id)" json:"-"`
	Jenis    Jenis  `gorm:"auto_preload;association_foreignkey:IDJenis"`
}

// ListPeribahasa list
type ListPeribahasa []Peribahasa

// TableName Peribahasa should not pruralized
func (Peribahasa) TableName() string {
	return "peribahasa"
}

func (p *Peribahasa) validate() error {
	if p.TeksAsli == "" {
		return errors.New("Kolom Teks Asli perlu di isi")
	}
	if p.Arti == "" {
		return errors.New("Kolom Arti harus di isi")
	}
	return nil
}

//Create New Peribahasa
func (p *Peribahasa) Create() error {
	if err := p.validate(); err != nil {
		return err
	}

	if err := GetDB().Create(&p).Error; err != nil {
		return err
	}
	return nil
}

//Get p Peribahasa
func (p *Peribahasa) Get(id int) error {
	if id <= 0 {
		err := GetDB().Order("random()").First(&p).Error
		if err != nil {
			return err
		}
		return nil
	}

	err := GetDB().Table("peribahasa").Where("id=?", id).First(&p).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Peribahasa data in the database
func (p *Peribahasa) Update(id int) error {
	if err := p.validate(); err != nil {
		return err
	}
	err := GetDB().Model(&Peribahasa{}).Where("id = ?", id).Update(&p).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete Peribahasa data in the database
func (p *Peribahasa) Delete(id int) error {
	err := GetDB().Where("id = ?", id).Delete(&p).Error
	if err != nil {
		return err
	}
	return nil
}

// List Peribahasa
func (p *ListPeribahasa) List(start int, max int) error {

	if start > 0 || max > 0 {
		err := GetDB().Preload("asal").Preload("jenis").Offset(start).Limit(max).Find(&p).Error
		if err != nil {
			return err
		}
		return nil
	}
	err := GetDB().Preload("asal").Preload("jenis").Find(&p).Error
	if err != nil {
		return err
	}
	return nil
}
