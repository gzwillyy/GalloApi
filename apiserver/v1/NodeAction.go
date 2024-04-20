package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeAction = "galloNodeActions"

// NodeAction 节点智能调度设置
type NodeAction struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NodeID            uint32 `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`       // 节点ID
	Role              string `gorm:"column:role;comment:角色" json:"role"`             // 角色
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"` // 是否启用
	Conds             string `gorm:"column:conds;comment:条件" json:"conds"`           // 条件
	Action            string `gorm:"column:action;comment:动作" json:"action"`         // 动作
	Duration          string `gorm:"column:duration;comment:持续时间" json:"duration"`   // 持续时间
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`           // 排序
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName NodeAction's table name
func (*NodeAction) TableName() string {
	return TableNameNodeAction
}

// AfterCreate run after create database record.
func (u *NodeAction) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "action-")).Error
}

type NodeActionList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeAction `json:"items"`
}
