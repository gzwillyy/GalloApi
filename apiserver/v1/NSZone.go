package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSZone = "galloNSZones"

// NSZone 域名子域
type NSZone struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	DomainID          uint32 `gorm:"column:domainId;comment:域名ID" json:"domainId"`   // 域名ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"` // 是否启用
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`           // 排序
	Version           uint32 `gorm:"column:version;comment:版本" json:"version"`       // 版本
	Tsig              string `gorm:"column:tsig;comment:TSIG配置" json:"tsig"`         // TSIG配置
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName NSZone's table name
func (*NSZone) TableName() string {
	return TableNameNSZone
}

// AfterCreate run after create database record.
func (u *NSZone) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "zone-")).Error
}

type NSZoneList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NSZone `json:"items"`
}
