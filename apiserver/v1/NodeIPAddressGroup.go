package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeIPAddressGroup = "galloNodeIPAddressGroups"

// NodeIPAddressGroup IP地址分组
type NodeIPAddressGroup struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Value             string `gorm:"column:value;comment:IP值" json:"value"` // IP值
}

// TableName NodeIPAddressGroup's table name
func (*NodeIPAddressGroup) TableName() string {
	return TableNameNodeIPAddressGroup
}

// AfterCreate run after create database record.
func (u *NodeIPAddressGroup) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "group-")).Error
}

type NodeIPAddressGroupList struct {
	metav1.ListMeta `json:",inline"`
	Items []*NodeIPAddressGroup `json:"items"`
}