package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameVersion = "galloVersions"

// Version 数据库结构版本
type Version struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Version           string `gorm:"column:version" json:"version"`
}

// TableName Version's table name
func (*Version) TableName() string {
	return TableNameVersion
}

// AfterCreate run after create database record.
func (u *Version) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "version-")).Error
}

type VersionList struct {
	metav1.ListMeta `json:",inline"`
	Items []*Version `json:"items"`
}
var VersionTableZeroFields = []string{"name", "version"}

