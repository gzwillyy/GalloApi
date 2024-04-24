package v1

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/gzwillyy/components/pkg/auth"
	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameAdmins = "galloAdmins"

// Admin represents a user restful resource. It is also used as gorm model.
type Admin struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status int `json:"status" gorm:"column:status" validate:"omitempty"`

	// Required: true
	Nickname string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`

	// Required: true
	Password string `json:"password,omitempty" gorm:"column:password" validate:"required"`

	// Required: true
	Email string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`

	Phone string `json:"phone" gorm:"column:phone" validate:"omitempty"`

	IsAdmin int `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`

	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`

	//TotalPolicy uint32 `json:"totalPolicy" gorm:"-" validate:"omitempty"`
}

// AdminList is the whole list of all admins which have been stored in stroage.
type AdminList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:",inline"`

	Items []*Admin `json:"items"`
}

// TableName maps to mysql table name.
func (u *Admin) TableName() string {
	return TableNameAdmins
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *Admin) Compare(pwd string) error {
	if err := auth.Compare(u.Password, pwd); err != nil {
		return fmt.Errorf("failed to compile password: %w", err)
	}

	return nil
}

// AfterCreate run after create database record.
func (u *Admin) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "admin-")).Error
}

var AdminTableZeroFields = []string{"name", "phone"}

