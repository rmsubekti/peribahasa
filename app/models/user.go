package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	gorm.Model
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Roles    []Role `gorm:"many2many:user_roles;"`
	Token    string `json:"-"`
}

// Validate on account creation
func (u *User) Validate() error {

	if u.UserName == "" {
		return errors.New("UserName is required")
	}
	// validate email
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid E-Mail address")
	}

	// validate password
	if len(u.Password) < 6 {
		return errors.New("Password should be 6 or more character")
	}

	return nil

}

// Create new User
func (u *User) Create() error {

	if err := u.Validate(); err != nil {
		return err
	}
	// check email if it is used
	temp := &User{}
	err := GetDB().Table("users").Where("email=?", u.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("Connection error")
	}
	// email is in use
	if temp.Email != "" {
		return errors.New("Email already in use by another user")
	}

	// check username if it is used
	e := GetDB().Table("users").Where("user_name=?", u.UserName).First(temp).Error
	if e != nil && e != gorm.ErrRecordNotFound {
		return errors.New("Connection error")
	}
	// email is in use
	if temp.UserName != "" {
		return errors.New("Username already in use by another user")
	}

	// set default user role
	var r Roles
	if err := r.getDefault(); err != nil {
		return err
	}
	u.Roles = r
	// encrypt password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	//get token

	if u.Token, err = generateToken(u.UserName, u.Roles); err != nil {
		return err
	}
	//write account in database
	if err := GetDB().Create(&u).Error; err != nil {
		return err
	}

	return nil
}

// Login account
func (u *User) Login(userNameOrEmail, password string) error {
	err := GetDB().Where("email = ?", userNameOrEmail).Or("user_name = ?", userNameOrEmail).Preload("Roles").Find(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("invalid login credential, Email or username not registered")
		}
		return errors.New("Connection error")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("invalid login credential, password not match")
	}
	return nil
}

// GetUser by id
func (u *User) GetUser(id uint) error {
	err := GetDB().Table("users").Where("id = ?", id).First(&u).Error
	if err != nil {
		return errors.New("No User with that ID")
	}
	return nil
}
