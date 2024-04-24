package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameReportNodeGroup = "galloReportNodeGroups"

// ReportNodeGroup 监控终端区域
type ReportNodeGroup struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	State             bool `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
	IsOn              bool `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"` // 是否启用
}

// TableName ReportNodeGroup's table name
func (*ReportNodeGroup) TableName() string {
	return TableNameReportNodeGroup
}

// AfterCreate run after create database record.
func (u *ReportNodeGroup) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "group-")).Error
}

type ReportNodeGroupList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*ReportNodeGroup `json:"items"`
}

var ReportNodeGroupTableZeroFields = []string{"name", "state", "isOn"}

