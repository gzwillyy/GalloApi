package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameReportNode = "galloReportNodes"

// ReportNode 连通性报告终端
type ReportNode struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	UniqueID          string `gorm:"column:uniqueId;comment:唯一ID" json:"uniqueId"`   // 唯一ID
	Secret            string `gorm:"column:secret;comment:密钥" json:"secret"`         // 密钥
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"` // 是否启用
	Location          string `gorm:"column:location;comment:所在区域" json:"location"`   // 所在区域
	Isp               string `gorm:"column:isp;comment:网络服务商" json:"isp"`            // 网络服务商
	AllowIPs          string `gorm:"column:allowIPs;comment:允许的IP" json:"allowIPs"`  // 允许的IP
	IsActive          uint32 `gorm:"column:isActive;comment:是否活跃" json:"isActive"`   // 是否活跃
	Status            string `gorm:"column:status;comment:状态" json:"status"`         // 状态
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
	GroupIds          string `gorm:"column:groupIds;comment:分组ID" json:"groupIds"`   // 分组ID
}

// TableName ReportNode's table name
func (*ReportNode) TableName() string {
	return TableNameReportNode
}

// AfterCreate run after create database record.
func (u *ReportNode) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "node-")).Error
}

type ReportNodeList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*ReportNode `json:"items"`
}
