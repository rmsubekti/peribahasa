package models

import (
	"errors"
)

// Asal has many quotes
type Asal struct {
	ID         uint
	Bahasa     string       `json:"bahasa"`
	Peribahasa []Peribahasa `gorm:"foreignkey:AsalID"`
}

//ListAsal type
type ListAsal []Asal

// TableName Asal should not pruralized
func (Asal) TableName() string {
	return "asal"
}

func (a *Asal) validate() error {
	if a.Bahasa == "" {
		return errors.New("Bahasa asal harus diisi")
	}
	return nil
}

// Create Asal
func (a *Asal) Create() error {
	if err := a.validate(); err != nil {
		return err
	}

	err := GetDB().Create(&a).Error
	if err != nil {
		return err
	}

	return nil
}

//Get Asal data from database
func (a *Asal) Get(id int) error {
	err := GetDB().Table("asal").Where("id=?", id).First(&a).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Asal data in the database
func (a *Asal) Update(id int) error {
	if err := a.validate(); err != nil {
		return err
	}

	err := GetDB().Model(&Asal{}).Where("id = ?", id).Update(&a).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete Asal data in the database
func (a *Asal) Delete(id int) error {
	err := GetDB().Where("id = ?", id).Delete(&a).Error
	if err != nil {
		return err
	}
	return nil
}

// List asal
func (a *ListAsal) List(start int, max int) error {
	if start > 0 || max > 0 {
		err := GetDB().Offset(start).Limit(max).Find(&a).Error
		if err != nil {
			return err
		}
		return nil
	}
	err := GetDB().Find(&a).Error
	if err != nil {
		return err
	}
	return nil
}
