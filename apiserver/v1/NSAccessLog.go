package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSAccessLog = "galloNSAccessLogs"

// NSAccessLog 域名服务访问日志
type NSAccessLog struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NodeID            uint32  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`       // 节点ID
	DomainID          uint32  `gorm:"column:domainId;comment:域名ID" json:"domainId"`   // 域名ID
	RecordID          uint32  `gorm:"column:recordId;comment:记录ID" json:"recordId"`   // 记录ID
	Content           string `gorm:"column:content;comment:访问数据" json:"content"`     // 访问数据
	RequestID         string `gorm:"column:requestId;comment:请求ID" json:"requestId"` // 请求ID
	RemoteAddr        string `gorm:"column:remoteAddr;comment:IP" json:"remoteAddr"` // IP
}

// TableName NSAccessLog's table name
func (*NSAccessLog) TableName() string {
	return TableNameNSAccessLog
}

// AfterCreate run after create database record.
func (u *NSAccessLog) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "log-")).Error
}

type NSAccessLogList struct {
	metav1.ListMeta `json:",inline"`
	Items []*NSAccessLog `json:"items"`
}