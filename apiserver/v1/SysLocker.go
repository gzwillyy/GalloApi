package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameSysLocker = "galloSysLockers"

// SysLocker 并发锁
type SysLocker struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Key               string `gorm:"column:key;comment:键值" json:"key"`               // 键值
	Version           uint32  `gorm:"column:version;comment:版本号" json:"version"`      // 版本号
	TimeoutAt         uint32  `gorm:"column:timeoutAt;comment:超时时间" json:"timeoutAt"` // 超时时间
}

// TableName SysLocker's table name
func (*SysLocker) TableName() string {
	return TableNameSysLocker
}

// AfterCreate run after create database record.
func (u *SysLocker) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "locker-")).Error
}

type SysLockerList struct {
	metav1.ListMeta `json:",inline"`
	Items []*SysLocker `json:"items"`
}
var SysLockerTableZeroFields = []string{"name", "key"}

