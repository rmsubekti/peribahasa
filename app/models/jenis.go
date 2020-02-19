package models

import (
	"errors"
)

// Jenis has many quotes
type Jenis struct {
	ID         uint
	Nama       string       `json:"type"`
	Peribahasa []Peribahasa `gorm:"foreignkey:JenisID"`
}

// TableName Jenis should not pruralized
func (Jenis) TableName() string {
	return "jenis"
}

//ListJenis Peribahasa
type ListJenis []Jenis

func (j *Jenis) validate() error {
	if j.Nama == "" {
		return errors.New("Nama Jenis Tidak Boleh Kosong")
	}
	return nil
}

// Create Jenis Name
func (j *Jenis) Create() error {
	if err := j.validate(); err != nil {
		return err
	}

	err := GetDB().Create(&j).Error
	if err != nil {
		return err
	}

	return nil
}

//Get Jenis data from database
func (j *Jenis) Get(id int) error {
	err := GetDB().Table("jenis").Where("id=?", id).First(&j).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Jenis data in the database
func (j *Jenis) Update(id int) error {
	if err := j.validate(); err != nil {
		return err
	}

	err := GetDB().Model(&Jenis{}).Where("id = ?", id).Update(&j).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete Jenis data in the database
func (j *Jenis) Delete(id int) error {
	err := GetDB().Where("id = ?", id).Delete(&j).Error
	if err != nil {
		return err
	}
	return nil
}

// List Jenis Peribahasa
func (j *ListJenis) List(start int, max int) error {
	err := GetDB().Find(&j).Error
	if err != nil {
		return err
	}
	return nil
}
