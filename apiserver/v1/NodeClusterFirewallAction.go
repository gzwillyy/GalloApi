package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeClusterFirewallAction = "galloNodeClusterFirewallActions"

// NodeClusterFirewallAction 防火墙动作
type NodeClusterFirewallAction struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`    // 管理员ID
	ClusterID         uint32 `gorm:"column:clusterId;comment:集群ID" json:"clusterId"` // 集群ID
	EventLevel        string `gorm:"column:eventLevel;comment:级别" json:"eventLevel"` // 级别
	Type              string `gorm:"column:type;comment:动作类型" json:"type"`           // 动作类型
	Params            string `gorm:"column:params;comment:参数" json:"params"`         // 参数
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName NodeClusterFirewallAction's table name
func (*NodeClusterFirewallAction) TableName() string {
	return TableNameNodeClusterFirewallAction
}

// AfterCreate run after create database record.
func (u *NodeClusterFirewallAction) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "firewall-")).Error
}

type NodeClusterFirewallActionList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeClusterFirewallAction `json:"items"`
}
