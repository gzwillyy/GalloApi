package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSRecordHourlyStat = "galloNSRecordHourlyStats"

// NSRecordHourlyStat NS记录统计
type NSRecordHourlyStat struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ClusterID         uint32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`        // 集群ID
	NodeID            uint32  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`              // 节点ID
	DomainID          uint32  `gorm:"column:domainId;comment:域名ID" json:"domainId"`          // 域名ID
	RecordID          uint32  `gorm:"column:recordId;comment:记录ID" json:"recordId"`          // 记录ID
	Day               string `gorm:"column:day;comment:YYYYMMDD" json:"day"`                // YYYYMMDD
	Hour              string `gorm:"column:hour;comment:YYYYMMDDHH" json:"hour"`            // YYYYMMDDHH
	CountRequests     uint32  `gorm:"column:countRequests;comment:请求数" json:"countRequests"` // 请求数
	Bytes             uint32  `gorm:"column:bytes;comment:流量" json:"bytes"`                  // 流量
}

// TableName NSRecordHourlyStat's table name
func (*NSRecordHourlyStat) TableName() string {
	return TableNameNSRecordHourlyStat
}

// AfterCreate run after create database record.
func (u *NSRecordHourlyStat) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "stat-")).Error
}

type NSRecordHourlyStatList struct {
	metav1.ListMeta `json:",inline"`
	Items []*NSRecordHourlyStat `json:"items"`
}