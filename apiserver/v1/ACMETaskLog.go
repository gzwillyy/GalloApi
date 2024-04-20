package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameACMETaskLog = "galloACMETaskLogs"

// ACMETaskLog ACME任务运行日志
type ACMETaskLog struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	TaskID            uint32 `gorm:"column:taskId;comment:任务ID" json:"taskId"`       // 任务ID
	IsOk              bool   `gorm:"column:isOk;default:1;comment:是否成功" json:"isOk"` // 是否成功
	Error             string `gorm:"column:error;comment:错误信息" json:"error"`         // 错误信息
}

// TableName ACMETaskLog's table name
func (*ACMETaskLog) TableName() string {
	return TableNameACMETaskLog
}

// AfterCreate run after create database record.
func (u *ACMETaskLog) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "log-")).Error
}

// ACMETaskLogList 返回列表
type ACMETaskLogList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*ACMETaskLog `json:"items"`
}
