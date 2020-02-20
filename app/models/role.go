package models

import (
	"errors"
)

// RoleType custom type
type RoleType string

// RoleUser enum
const (
	RoleUser   RoleType = "user"
	RoleAdmin  RoleType = "admin"
	RoleEditor RoleType = "editor"
)

// RoleTypes list
type RoleTypes []RoleType

// Role Account
type Role struct {
	ID       uint
	RoleName RoleType
}

// Roles list
type Roles []Role

// IsAllowed : Check if user has access to the resource
func (r Roles) IsAllowed(allowedRoles RoleTypes) error {
	if len(r) == 0 {
		return errors.New("There is no role to check")
	}

	for _, v := range allowedRoles {
		if r.isExist(v) {
			return nil
		}
	}
	return errors.New("Keep outside")
}

// isExist in the logged user roles
func (r Roles) isExist(expected RoleType) bool {
	for _, v := range r {
		if v.RoleName == expected {
			return true
		}
	}
	return false
}

// getDefault roles
func (r *Roles) getDefault() error {
	role := &Role{}
	err := GetDB().Table("roles").Where("role_name=?", string(RoleUser)).First(role).Error
	if err != nil {
		return err
	}
	r = &Roles{*role}
	return nil
}
