package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameReportResult = "galloReportResults"

// ReportResult 连通性监控结果
type ReportResult struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Type              string  `gorm:"column:type;comment:对象类型" json:"type"`                       // 对象类型
	TargetID          uint32   `gorm:"column:targetId;comment:对象ID" json:"targetId"`               // 对象ID
	TargetDesc        string  `gorm:"column:targetDesc;comment:对象描述" json:"targetDesc"`           // 对象描述
	ReportNodeID      uint32   `gorm:"column:reportNodeId;comment:监控节点ID" json:"reportNodeId"`     // 监控节点ID
	IsOk              uint32   `gorm:"column:isOk;default:1;comment:是否可连接" json:"isOk"`            // 是否可连接
	Level             string  `gorm:"column:level;comment:级别" json:"level"`                       // 级别
	CostMs            float64 `gorm:"column:costMs;default:0.00;comment:单次连接花费的时间" json:"costMs"` // 单次连接花费的时间
	Error             string  `gorm:"column:error;comment:产生的错误信息" json:"error"`                  // 产生的错误信息
	CountUp           uint32   `gorm:"column:countUp;comment:连续上线次数" json:"countUp"`               // 连续上线次数
	CountDown         uint32   `gorm:"column:countDown;comment:连续下线次数" json:"countDown"`           // 连续下线次数
}

// TableName ReportResult's table name
func (*ReportResult) TableName() string {
	return TableNameReportResult
}

// AfterCreate run after create database record.
func (u *ReportResult) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "result-")).Error
}

type ReportResultList struct {
	metav1.ListMeta `json:",inline"`
	Items []*ReportResult `json:"items"`
}