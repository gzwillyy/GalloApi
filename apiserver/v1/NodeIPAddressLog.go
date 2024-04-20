package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeIPAddressLog = "galloNodeIPAddressLogs"

// NodeIPAddressLog IP状态变更日志
type NodeIPAddressLog struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AddressID         uint32 `gorm:"column:addressId;comment:地址ID" json:"addressId"`   // 地址ID
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`      // 管理员ID
	Description       string `gorm:"column:description;comment:描述" json:"description"` // 描述
	IsUp              uint32 `gorm:"column:isUp;comment:是否在线" json:"isUp"`             // 是否在线
	IsOn              bool   `gorm:"column:isOn;comment:是否启用" json:"isOn"`             // 是否启用
	CanAccess         uint32 `gorm:"column:canAccess;comment:是否可访问" json:"canAccess"`  // 是否可访问
	Day               string `gorm:"column:day;comment:YYYYMMDD，用来清理" json:"day"`      // YYYYMMDD，用来清理
	BackupIP          string `gorm:"column:backupIP;comment:备用IP" json:"backupIP"`     // 备用IP
}

// TableName NodeIPAddressLog's table name
func (*NodeIPAddressLog) TableName() string {
	return TableNameNodeIPAddressLog
}

// AfterCreate run after create database record.
func (u *NodeIPAddressLog) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "log-")).Error
}

type NodeIPAddressLogList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeIPAddressLog `json:"items"`
}
