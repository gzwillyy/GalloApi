package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSDomainGroup = "galloNSDomainGroups"

// NSDomainGroup 域名分组
type NSDomainGroup struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`       // 用户ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"` // 是否启用
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`           // 排序
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName NSDomainGroup's table name
func (*NSDomainGroup) TableName() string {
	return TableNameNSDomainGroup
}

// AfterCreate run after create database record.
func (u *NSDomainGroup) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "group-")).Error
}

type NSDomainGroupList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NSDomainGroup `json:"items"`
}

var NSDomainGroupTableZeroFields = []string{"name", "state"}
